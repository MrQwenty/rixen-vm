package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rx",
	Short: "Rixen - A VM manager for developers",
	Long: `Rixen (rx) is a command-line tool for managing virtual machines on macOS.
It supports creating, starting, and managing VMs with a focus on developer experience.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	cobra.OnInitialize(initConfig)

	// Add commands
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(osCmd)
}

func initConfig() {
	// Create ~/.rx directory if it doesn't exist
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		os.Exit(1)
	}

	rxDir := fmt.Sprintf("%s/.rx", homeDir)
	if _, err := os.Stat(rxDir); os.IsNotExist(err) {
		if err := os.Mkdir(rxDir, 0755); err != nil {
			fmt.Println("Error creating ~/.rx directory:", err)
			os.Exit(1)
		}
	}

	// Create subdirectories
	for _, dir := range []string{"iso", "vms"} {
		path := fmt.Sprintf("%s/%s", rxDir, dir)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			if err := os.Mkdir(path, 0755); err != nil {
				fmt.Println("Error creating directory:", path, err)
				os.Exit(1)
			}
		}
	}
}

// PrintSuccess prints a success message in green
func PrintSuccess(format string, a ...interface{}) {
	green := color.New(color.FgGreen).SprintFunc()
	message := fmt.Sprintf(format, a...)
	fmt.Println(green("✓ " + message))
}

// PrintInfo prints an info message in blue
func PrintInfo(format string, a ...interface{}) {
	blue := color.New(color.FgBlue).SprintFunc()
	message := fmt.Sprintf(format, a...)
	fmt.Println(blue("ℹ " + message))
}

// PrintError prints an error message in red
func PrintError(format string, a ...interface{}) {
	red := color.New(color.FgRed).SprintFunc()
	message := fmt.Sprintf(format, a...)
	fmt.Println(red("✗ " + message))
}