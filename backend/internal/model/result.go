package model

type Address struct {
	IP string `json:"ip"`
}

type ScanResult struct {
	Name      string    `json:"name"`
	Domain    string    `json:"domain"`
	Tag       string    `json:"tag"`
	Addresses []Address `json:"addresses"`
}
