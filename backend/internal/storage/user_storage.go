package storage

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"

	"backend/internal/models"
)

type UserStorage struct {
	mu    sync.RWMutex
	users map[string]*models.User
	path  string
}

func NewUserStorage() (*UserStorage, error) {
	// Create data directory if it doesn't exist
	dataDir := "data"
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return nil, err
	}

	storage := &UserStorage{
		users: make(map[string]*models.User),
		path:  filepath.Join(dataDir, "users.json"),
	}

	// Load existing users if the file exists
	if err := storage.load(); err != nil {
		return nil, err
	}

	return storage, nil
}

func (s *UserStorage) load() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check if file exists
	if _, err := os.Stat(s.path); os.IsNotExist(err) {
		// Create empty file if it doesn't exist
		return s.save()
	}

	// Read file
	data, err := os.ReadFile(s.path)
	if err != nil {
		return err
	}

	// Parse JSON
	if len(data) > 0 {
		if err := json.Unmarshal(data, &s.users); err != nil {
			return err
		}
	}

	return nil
}

func (s *UserStorage) save() error {
	data, err := json.MarshalIndent(s.users, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.path, data, 0644)
}

func (s *UserStorage) CreateUser(user *models.User) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check if user already exists
	if _, exists := s.users[user.Username]; exists {
		return models.ErrUserAlreadyExists
	}

	// Add user to map
	s.users[user.Username] = user

	// Save to file
	return s.save()
}

func (s *UserStorage) GetUser(username string) (*models.User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	user, exists := s.users[username]
	if !exists {
		return nil, models.ErrUserNotFound
	}

	return user, nil
}

func (s *UserStorage) UpdateUser(user *models.User) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check if user exists
	if _, exists := s.users[user.Username]; !exists {
		return models.ErrUserNotFound
	}

	// Update user
	s.users[user.Username] = user

	// Save to file
	return s.save()
} 