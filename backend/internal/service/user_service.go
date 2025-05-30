package service

import (
	"crypto/rand"
	"encoding/base64"
	"time"

	"backend/internal/model"
	"backend/internal/storage"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	storage *storage.UserStorage
}

func NewUserService() (*UserService, error) {
	userStorage, err := storage.NewUserStorage()
	if err != nil {
		return nil, err
	}

	return &UserService{
		storage: userStorage,
	}, nil
}

func (s *UserService) Register(reg model.UserRegistration) error {
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(reg.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Create user
	user := &model.User{
		ID:        generateID(),
		Username:  reg.Username,
		Password:  string(hashedPassword),
		Email:     reg.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Save user
	return s.storage.CreateUser(user)
}

func (s *UserService) Login(login model.UserLogin) (*model.User, error) {
	// Get user
	user, err := s.storage.GetUser(login.Username)
	if err != nil {
		return nil, err
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)); err != nil {
		return nil, model.ErrInvalidPassword
	}

	return user, nil
}

// generateID generates a random ID
func generateID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
} 