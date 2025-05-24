package oslist

// FedoraProvider provides Fedora OS information
type FedoraProvider struct{}

// Name returns the name of the OS
func (p *FedoraProvider) Name() string {
	return "Fedora"
}

// GetVersions returns the available versions of Fedora
func (p *FedoraProvider) GetVersions() (*OSInfo, error) {
	return &OSInfo{
		Name: "Fedora",
		Versions: []string{
			"39",
			"38",
			"37",
			"36",
		},
	}, nil
}

// GetDownloadURL returns the download URL for a specific version
func (p *FedoraProvider) GetDownloadURL(version string) (string, error) {
	switch version {
	case "39":
		return "https://download.fedoraproject.org/pub/fedora/linux/releases/39/Workstation/x86_64/iso/Fedora-Workstation-Live-x86_64-39-1.5.iso", nil
	case "38":
		return "https://download.fedoraproject.org/pub/fedora/linux/releases/38/Workstation/x86_64/iso/Fedora-Workstation-Live-x86_64-38-1.6.iso", nil
	case "37":
		return "https://download.fedoraproject.org/pub/fedora/linux/releases/37/Workstation/x86_64/iso/Fedora-Workstation-Live-x86_64-37-1.7.iso", nil
	case "36":
		return "https://download.fedoraproject.org/pub/fedora/linux/releases/36/Workstation/x86_64/iso/Fedora-Workstation-Live-x86_64-36-1.5.iso", nil
	default:
		return "", ErrVersionNotFound
	}
}