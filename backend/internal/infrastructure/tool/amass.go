package tool

import (
	"fmt"
	"os/exec"
)

// AmassTool provides basic domain enumeration and intelligence gathering
type AmassTool struct {
	dockerImage string
}

// NewAmassTool creates a new instance of AmassTool
func NewAmassTool() (*AmassTool, error) {
	return &AmassTool{
		dockerImage: "caffix/amass",
	}, nil
}

// ScanDomain performs domain enumeration with active scanning and brute forcing
func (a *AmassTool) ScanDomain(domain string) (string, error) {
	cmd := exec.Command("docker", 
		"run",
		"--name", "amass-scan",
		"--rm",
		"-v", "C:/Projects/Cyvore/data:/data",
		a.dockerImage,
		"enum",
		"-active",
		"-brute",
		"-rf", "/data/resolvers.txt",
		"-log", "/data/amass/log.txt",
		"-o", "/data/amass/results.txt",
		"-dir", "/data/amass",
		"-d", domain)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("amass enum error: %v\nOutput: %s", err, string(output))
	}

	return string(output), nil
}
