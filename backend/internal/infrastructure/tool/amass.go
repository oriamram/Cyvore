package tool

import (
	"backend/internal/model"
	"encoding/json"
	"os/exec"
	"strings"
)

func RunAmass(domain string) ([]model.ScanResult, error) {
	cmd := exec.Command("amass", "enum", "-d", domain, "-json", "-")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(strings.NewReader(string(output)))
	var results []model.ScanResult
	for decoder.More() {
		var entry model.ScanResult
		err := decoder.Decode(&entry)
		if err == nil {
			results = append(results, entry)
		}
	}
	return results, nil
}
