package model

// DomainInfo represents information about a domain
type DomainInfo struct {
	Domain     string   `json:"domain"`
	Subdomains []string `json:"subdomains"`
	IPs        []string `json:"ips"`
} 