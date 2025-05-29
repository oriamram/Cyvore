package service

import (
	"backend/internal/infrastructure/tool"
	"fmt"
)

// ScanService handles domain scanning operations
type ScanService struct {
	amassTool *tool.AmassTool
}

// NewScanService creates a new instance of ScanService
func NewScanService() (*ScanService, error) {
	amassTool, err := tool.NewAmassTool()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize amass tool: %v", err)
	}

	return &ScanService{
		amassTool: amassTool,
	}, nil
}

// GetDomainInfo retrieves information about a domain
func (s *ScanService) GetDomainInfo(domain string) (string, error) {
	// Run the scan
	output, err := s.amassTool.ScanDomain(domain)
	if err != nil {
		return "", fmt.Errorf("scan failed: %v", err)
	}

	return output, nil
}
