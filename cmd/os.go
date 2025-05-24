package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/rixen/rx/internal/oslist"
	"github.com/spf13/cobra"
	"os"
)

var osCmd = &cobra.Command{
	Use:   "os",
	Short: "Manage operating systems",
	Long:  `Commands for listing and managing supported operating systems.`,
}

var osListCmd = &cobra.Command{
	Use:   "list",
	Short: "List supported operating systems",
	Long:  `List all operating systems supported by rx.`,
	Run: func(cmd *cobra.Command, args []string) {
		providers := oslist.GetAllProviders()
		
		fmt.Println("Supported operating systems:")
		fmt.Println()
		
		for _, provider := range providers {
			bold := color.New(color.FgCyan, color.Bold).SprintFunc()
			fmt.Printf("  %s\n", bold(provider.Name()))
		}
		
		fmt.Println("\nTo see available versions, run: rx os versions <os-name>")
	},
}

var osVersionsCmd = &cobra.Command{
	Use:   "versions [os-name]",
	Short: "List available versions for an OS",
	Long:  `List all available versions for the specified operating system.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		osName := args[0]
		
		provider, err := oslist.GetProvider(osName)
		if err != nil {
			PrintError("%v", err)
			os.Exit(1)
		}
		
		PrintInfo("Getting available versions for %s...", provider.Name())
		
		osInfo, err := provider.GetVersions()
		if err != nil {
			PrintError("Error getting versions: %v", err)
			os.Exit(1)
		}
		
		bold := color.New(color.FgGreen, color.Bold).SprintFunc()
		fmt.Printf("\nAvailable versions for %s:\n", bold(provider.Name()))
		fmt.Println()
		
		for i, version := range osInfo.Versions {
			fmt.Printf("  %s\n", version)
			// Add latest tag to the first one
			if i == 0 {
				latest := color.New(color.FgYellow).SprintFunc()
				fmt.Printf("    %s\n", latest("(latest)"))
			}
		}
		
		fmt.Println("\nTo create a VM with this OS, run:")
		fmt.Printf("  rx create --name myvm --os %s\n", osName)
	},
}

func init() {
	osCmd.AddCommand(osListCmd)
	osCmd.AddCommand(osVersionsCmd)
}