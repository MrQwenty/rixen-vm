package oslist

import (
	"fmt"
)

// OSInfo represents information about an OS
type OSInfo struct {
	Name     string   `json:"name"`
	Versions []string `json:"versions"`
}

// OSProvider is the interface that all OS providers must implement
type OSProvider interface {
	// Name returns the name of the OS
	Name() string
	
	// GetVersions returns the available versions of the OS
	GetVersions() (*OSInfo, error)
	
	// GetDownloadURL returns the download URL for a specific version
	GetDownloadURL(version string) (string, error)
}

// GetProvider returns the provider for the specified OS
func GetProvider(osName string) (OSProvider, error) {
	switch osName {
	case "ubuntu":
		return &UbuntuProvider{}, nil
	case "fedora":
		return &FedoraProvider{}, nil
	case "windows":
		return &WindowsProvider{}, nil
	default:
		return nil, fmt.Errorf("unsupported OS: %s", osName)
	}
}

// GetAllProviders returns all available OS providers
func GetAllProviders() []OSProvider {
	return []OSProvider{
		&UbuntuProvider{},
		&FedoraProvider{},
		&WindowsProvider{},
	}
}