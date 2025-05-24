package config

// VMConfig holds the configuration for a VM
type VMConfig struct {
    Name     string `json:"name"`
    ISO      string `json:"iso"`
    OS       string `json:"os"`
    CPUCount int    `json:"cpu_count"`
    RAMSize  int    `json:"ram_size"`
    DiskSize int    `json:"disk_size"`
    DiskPath string `json:"disk_path"`
}