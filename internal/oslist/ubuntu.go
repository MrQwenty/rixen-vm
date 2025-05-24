package oslist

// UbuntuProvider provides Ubuntu OS information
type UbuntuProvider struct{}

// Name returns the name of the OS
func (p *UbuntuProvider) Name() string {
	return "Ubuntu"
}

// GetVersions returns the available versions of Ubuntu
func (p *UbuntuProvider) GetVersions() (*OSInfo, error) {
	return &OSInfo{
		Name: "Ubuntu",
		Versions: []string{
			"22.04",
			"20.04",
			"18.04",
			"16.04",
		},
	}, nil
}

// GetDownloadURL returns the download URL for a specific version
func (p *UbuntuProvider) GetDownloadURL(version string) (string, error) {
	switch version {
	case "22.04":
		return "https://releases.ubuntu.com/22.04/ubuntu-22.04.3-desktop-amd64.iso", nil
	case "20.04":
		return "https://releases.ubuntu.com/20.04/ubuntu-20.04.6-desktop-amd64.iso", nil
	case "18.04":
		return "https://releases.ubuntu.com/18.04/ubuntu-18.04.6-desktop-amd64.iso", nil
	case "16.04":
		return "https://releases.ubuntu.com/16.04/ubuntu-16.04.7-desktop-amd64.iso", nil
	default:
		return "", ErrVersionNotFound
	}
}