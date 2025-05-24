package oslist

import "errors"

var ErrVersionNotFound = errors.New("version not found")

// WindowsProvider provides Windows OS information
type WindowsProvider struct{}

// Name returns the name of the OS
func (p *WindowsProvider) Name() string {
	return "Windows"
}

// GetVersions returns the available versions of Windows
func (p *WindowsProvider) GetVersions() (*OSInfo, error) {
	return &OSInfo{
		Name: "Windows",
		Versions: []string{
			"11",
			"10",
		},
	}, nil
}

// GetDownloadURL returns the download URL for a specific version
// For Windows, we don't provide direct download URLs since they require Microsoft's license agreement
func (p *WindowsProvider) GetDownloadURL(version string) (string, error) {
	switch version {
	case "11":
		return "https://www.microsoft.com/software-download/windows11", nil
	case "10":
		return "https://www.microsoft.com/software-download/windows10", nil
	default:
		return "", ErrVersionNotFound
	}
}