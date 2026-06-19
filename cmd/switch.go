package cmd

import (
	"log"
	"os"
	"os/exec"

	"github.com/achill06/git-zen/internal/tui"
	"github.com/spf13/cobra"
)

var switchCmd = &cobra.Command{
	Use:   "switch",
	Short: "Fuzzy search and switch branches",
	Long:  `Open an interactive dashboard to fuzzy-search your local branches. If authenticated with the GitHub CLI, it will automatically fetch and display live Pull Request statuses in the background.`,
	Run: func(cmd *cobra.Command, args []string) {

		selectedBranch, err := tui.Run()
		if err != nil {
			log.Fatalf("Error running TUI: %v", err)
		}

		if selectedBranch != "" {
			checkoutCmd := exec.Command("git", "checkout", selectedBranch)
			checkoutCmd.Stdout = os.Stdout
			checkoutCmd.Stderr = os.Stderr

			if err := checkoutCmd.Run(); err != nil {
				log.Fatalf("Failed to switch branch: %v", err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(switchCmd)
}
