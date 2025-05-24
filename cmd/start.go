package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"runtime"

	"github.com/rixen/rx/internal/config"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start [vm-name]",
	Short: "Start a virtual machine",
	Long:  `Start a virtual machine using QEMU with shared folders and networking.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		vmName := args[0]

		// Trova la home dell’utente
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Printf("Errore: impossibile trovare la home directory: %v\n", err)
			os.Exit(1)
		}

		vmDir := filepath.Join(homeDir, ".rx", "vms", vmName)
		configPath := filepath.Join(vmDir, "config.json")

		// Verifica esistenza VM
		if _, err := os.Stat(configPath); os.IsNotExist(err) {
			fmt.Printf("Errore: la VM '%s' non esiste\n", vmName)
			os.Exit(1)
		}

		// Carica il config.json
		configData, err := os.ReadFile(configPath)
		if err != nil {
			fmt.Printf("Errore nella lettura del file di configurazione: %v\n", err)
			os.Exit(1)
		}

		var config config.VMConfig
		if err := json.Unmarshal(configData, &config); err != nil {
			fmt.Printf("Errore nel parsing del file di configurazione: %v\n", err)
			os.Exit(1)
		}

		// Verifica installazione QEMU
		_, err = exec.LookPath("qemu-system-x86_64")
		if err != nil {
			fmt.Println("Errore: qemu-system-x86_64 non trovato. Installa QEMU (es: brew install qemu)")
			os.Exit(1)
		}

		// Cartella condivisa
		sharedDir := filepath.Join(homeDir, ".rx", "workspaces", config.Name)
		os.MkdirAll(sharedDir, 0755)

		// Costruisci gli argomenti QEMU
		qemuArgs := []string{
			"-name", config.Name,
			"-cpu", func() string {
				if runtime.GOOS == "darwin" {
					return "max"
				}
				return "host"
			}(),
			"-smp", fmt.Sprintf("%d", config.CPUCount),
			"-m", fmt.Sprintf("%d", config.RAMSize),
			"-hda", config.DiskPath,
		}

		// ISO se presente
		if config.ISO != "" {
			if _, err := os.Stat(config.ISO); err == nil {
				qemuArgs = append(qemuArgs, "-cdrom", config.ISO, "-boot", "d")
			}
		}

		// Networking e display
		qemuArgs = append(qemuArgs,
			"-device", "virtio-net,netdev=net0",
			"-netdev", "user,id=net0,hostfwd=tcp::2222-:22",
			"-vga", "virtio",
			"-display", "default",
		)

		// VirtioFS
		qemuArgs = append(qemuArgs,
			"-fsdev", fmt.Sprintf("local,id=fsdev0,path=%s,security_model=none", sharedDir),
			"-device", "virtio-9p-pci,fsdev=fsdev0,mount_tag=hostshare",
		)

		fmt.Printf("🚀 Avvio VM '%s'...\n", config.Name)

		qemuCmd := exec.Command("qemu-system-x86_64", qemuArgs...)
		qemuCmd.Stdin = os.Stdin
		qemuCmd.Stdout = os.Stdout
		qemuCmd.Stderr = os.Stderr

		if err := qemuCmd.Run(); err != nil {
			fmt.Printf("Errore nell'avvio della VM: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
