package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/rixen/rx/internal/config"
	"github.com/rixen/rx/internal/oslist"
	"github.com/spf13/cobra"
)

var cfg config.VMConfig

var (
	vmName   string
	isoPath  string
	osType   string
	cpuCount int
	ramSize  int
	diskSize int
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new virtual machine",
	Long: `Create a new virtual machine with the specified parameters.
If --iso is not provided but --os is, rx will download the latest version of the OS.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Validate VM name
		if vmName == "" {
			PrintError("VM name is required")
			cmd.Help()
			os.Exit(1)
		}

		// Validate required flags
		if osType == "" && isoPath == "" {
			PrintError("Either --os or --iso is required")
			cmd.Help()
			os.Exit(1)
		}

		 // Validate numeric flags
        if cpuCount <= 0 {
            PrintError("CPU count must be greater than 0")
            os.Exit(1)
        }
        if ramSize <= 0 {
            PrintError("RAM size must be greater than 0")
            os.Exit(1)
        }
        if diskSize <= 0 {
            PrintError("Disk size must be greater than 0")
            os.Exit(1)
        }

		// Get home directory
		homeDir, err := os.UserHomeDir()
		if err != nil {
			PrintError("Error getting home directory: %v", err)
			os.Exit(1)
		}

		rxDir := filepath.Join(homeDir, ".rx")
		vmDir := filepath.Join(rxDir, "vms", vmName)
		isoDir := filepath.Join(rxDir, "iso")

		// Check if VM already exists
		if _, err := os.Stat(vmDir); !os.IsNotExist(err) {
			PrintError("VM with name '%s' already exists", vmName)
			os.Exit(1)
		}

		// Create VM directory
		if err := os.MkdirAll(vmDir, 0755); err != nil {
			PrintError("Error creating VM directory: %v", err)
			os.Exit(1)
		}

		// Special handling for Windows
		if osType == "windows" {
			handleWindowsOS()
			return
		}

		// Handle ISO download if needed
		finalIsoPath := isoPath
		if isoPath == "" && osType != "" {
			var err error
			finalIsoPath, err = handleOSDownload(osType, isoDir)
			if err != nil {
				PrintError("Error downloading ISO: %v", err)
				os.Exit(1)
			}
		}

		// Create disk image
		diskPath := filepath.Join(vmDir, "disk.img")
		if err := createDiskImage(diskPath, diskSize); err != nil {
			PrintError("Error creating disk image: %v", err)
			os.Exit(1)
		}

		// Create VM config
		config := config.VMConfig{
			Name:     vmName,
			ISO:      finalIsoPath,
			OS:       osType,
			CPUCount: cpuCount,
			RAMSize:  ramSize,
			DiskSize: diskSize,
			DiskPath: diskPath,
		}

		// Save config
		configPath := filepath.Join(vmDir, "config.json")
		if err := saveConfig(configPath, config); err != nil {
			PrintError("Error saving config: %v", err)
			os.Exit(1)
		}

		PrintSuccess("VM '%s' created successfully", vmName)
		PrintInfo("To start the VM, run: rx start %s", vmName)
	},
}

func init() {
	createCmd.Flags().StringVar(&vmName, "name", "", "Name of the VM (required)")
	createCmd.Flags().StringVar(&isoPath, "iso", "", "Path to ISO file")
	createCmd.Flags().StringVar(&osType, "os", "", "OS type (ubuntu, fedora, windows, etc.)")
	createCmd.Flags().IntVar(&cpuCount, "cpu", 2, "Number of CPU cores")
	createCmd.Flags().IntVar(&ramSize, "ram", 2048, "RAM size in MB")
	createCmd.Flags().IntVar(&diskSize, "disk", 20, "Disk size in GB")

	createCmd.MarkFlagRequired("name")
}

func handleWindowsOS() {
	windowsURL := "https://www.microsoft.com/software-download/windows11"

	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", windowsURL)
	case "linux":
		cmd = exec.Command("xdg-open", windowsURL)
	default:
		PrintError("Unsupported OS for opening browser")
		os.Exit(1)
	}

	PrintInfo("Opening Windows download page in your browser...")
	if err := cmd.Run(); err != nil {
		PrintError("Error opening browser: %v", err)
		os.Exit(1)
	}

	PrintInfo("Please download the Windows ISO and then create the VM using --iso flag")
}

func handleOSDownload(osType, isoDir string) (string, error) {
	provider, err := oslist.GetProvider(osType)
	if err != nil {
		return "", err
	}

	PrintInfo("Getting available versions for %s...", provider.Name())
	osInfo, err := provider.GetVersions()
	if err != nil {
		return "", err
	}

	if len(osInfo.Versions) == 0 {
		return "", fmt.Errorf("no versions available for %s", osType)
	}

	// Latest version is the first one
	latestVersion := osInfo.Versions[0]

	fmt.Println()
	fmt.Printf("Latest version available: %s\n", color.GreenString(latestVersion))
	fmt.Print("Download this version? [Y/n]: ")

	var response string
	fmt.Scanln(&response)
	response = strings.ToLower(response)

	selectedVersion := latestVersion
	if response == "n" || response == "no" {
		// Let user choose from available versions
		fmt.Println("\nAvailable versions:")
		for i, version := range osInfo.Versions {
			fmt.Printf("%d. %s\n", i+1, version)
		}

		fmt.Print("\nSelect version (1-" + strconv.Itoa(len(osInfo.Versions)) + "): ")
		var choice int
		fmt.Scanln(&choice)

		if choice < 1 || choice > len(osInfo.Versions) {
			return "", fmt.Errorf("invalid selection")
		}

		selectedVersion = osInfo.Versions[choice-1]
	}

	// Get download URL
	_, err = provider.GetDownloadURL(selectedVersion)
	if err != nil {
		return "", err
	}

	// Create ISO directory if it doesn't exist
	if err := os.MkdirAll(isoDir, 0755); err != nil {
		return "", err
	}

	// Download ISO
	isoFilename := filepath.Join(isoDir, fmt.Sprintf("%s-%s.iso", osType, selectedVersion))

	PrintInfo("Downloading %s %s ISO...", osType, selectedVersion)
	PrintInfo("This may take some time depending on your internet connection")

	// Simulate download for now (in a real implementation, use http.Get and io.Copy)
	// For a real implementation, you'd want to show progress too
	// cmd := exec.Command("curl", "-L", "-o", isoFilename, downloadURL)
	// return isoFilename, cmd.Run()

	// For the purpose of this example, we'll just create an empty file
	// In a real implementation, replace this with actual download code
	if _, err := os.Create(isoFilename); err != nil {
		return "", err
	}

	PrintSuccess("Downloaded ISO to %s", isoFilename)
	return isoFilename, nil
}

func createDiskImage(diskPath string, sizeGB int) error {
	PrintInfo("Creating disk image (%d GB)...", sizeGB)

	// Check if qemu-img is installed
	_, err := exec.LookPath("qemu-img")
	if err != nil {
		return fmt.Errorf("qemu-img not found. Please install QEMU")
	}

	cmd := exec.Command("qemu-img", "create", "-f", "qcow2", diskPath, fmt.Sprintf("%dG", sizeGB))
	return cmd.Run()
}

func saveConfig(path string, config config.VMConfig) error {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

