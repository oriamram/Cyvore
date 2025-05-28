package service

import (
	"backend/internal/infrastructure/tool"
	"backend/internal/model"
)

func StartAmassScan(domain string) ([]model.ScanResult, error) {
	return tool.RunAmass(domain)
}
