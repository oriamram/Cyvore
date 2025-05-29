package tool

import (
	"fmt"
	"os/exec"
	"sync"

	"backend/internal/config"
)

// AmassTool provides basic domain enumeration and intelligence gathering
type AmassTool struct {
	dockerImage string
	cmd         *exec.Cmd
	isRunning   bool
	mu          sync.Mutex
}

// NewAmassTool creates a new instance of AmassTool
func NewAmassTool() (*AmassTool, error) {
	return &AmassTool{
		dockerImage: "caffix/amass",
		isRunning:   false,
	}, nil
}

// IsScanning returns whether a scan is currently running
func (a *AmassTool) IsScanning() bool {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.isRunning
}

// ScanDomain performs domain enumeration with active scanning and brute forcing
func (a *AmassTool) ScanDomain(domain string) error {
	a.mu.Lock()
	if a.isRunning {
		a.mu.Unlock()
		return fmt.Errorf("scan already in progress")
	}
	a.isRunning = true
	a.mu.Unlock()

	// Stop any existing scan first
	a.StopScan()

	// Get the data path from config
	cfg := config.Get()
	dataPath := cfg.DataPath

	cmd := exec.Command("docker", 
		"run",
		"--name", "amass-scan",
		"--rm",
		"-v", fmt.Sprintf("%s:/data", dataPath),
		a.dockerImage,
		"enum",
		"-active",
		"-brute",
		"-rf", "/data/resolvers.txt",
		"-log", "/data/amass/log.txt",
		"-o", "/data/amass/results.txt",
		"-dir", "/data/amass",
		"-d", domain)

	// Start the command in the background
	if err := cmd.Start(); err != nil {
		a.mu.Lock()
		a.isRunning = false
		a.mu.Unlock()
		return fmt.Errorf("failed to start amass scan: %v", err)
	}

	a.cmd = cmd

	// Start a goroutine to wait for the command to complete
	go func() {
		cmd.Wait()
		a.mu.Lock()
		a.isRunning = false
		a.mu.Unlock()
	}()

	return nil
}

// StopScan stops the running scan by stopping the Docker container
func (a *AmassTool) StopScan() error {
	cmd := exec.Command("docker", "stop", "amass-scan")
	output, err := cmd.CombinedOutput()
	if err != nil {
		// Ignore error if container doesn't exist
		if string(output) != "Error: No such container: amass-scan\n" {
			return fmt.Errorf("failed to stop amass scan: %v\nOutput: %s", err, string(output))
		}
	}
	
	a.mu.Lock()
	a.isRunning = false
	a.mu.Unlock()
	
	return nil
}
