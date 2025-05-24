package oslist

import "errors"

// Define common errors
var (
	// ErrUnsupportedOS is returned when an unsupported OS is requested
	ErrUnsupportedOS = errors.New("unsupported operating system")
	
	// ErrNoVersionsAvailable is returned when no versions are available for an OS
	ErrNoVersionsAvailable = errors.New("no versions available for this OS")
)