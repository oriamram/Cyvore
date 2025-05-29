package service

import (
	"backend/internal/infrastructure/tool"
	"fmt"
	"sync"
)

var (
	instance *ScanService
	once     sync.Once
)

// ScanService handles domain scanning operations
type ScanService struct {
	amassTool *tool.AmassTool
}

// GetScanService returns the singleton instance of ScanService
func GetScanService() (*ScanService, error) {
	var initErr error
	once.Do(func() {
		amassTool, err := tool.NewAmassTool()
		if err != nil {
			initErr = fmt.Errorf("failed to initialize amass tool: %v", err)
			return
		}
		instance = &ScanService{
			amassTool: amassTool,
		}
	})
	if initErr != nil {
		return nil, initErr
	}
	return instance, nil
}

// GetDomainInfo starts a scan for the given domain
func (s *ScanService) GetDomainInfo(domain string) error {
	// Start the scan
	return s.amassTool.ScanDomain(domain)
}

// StopScan stops the running scan
func (s *ScanService) StopScan() error {
	return s.amassTool.StopScan()
}

// IsScanning returns whether a scan is currently running
func (s *ScanService) IsScanning() bool {
	return s.amassTool.IsScanning()
}
