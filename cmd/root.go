package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git-zen",
	Short: "A unified Git CLI extension suite",
	Long:  `git-zen is an interactive TUI for git branch management, AI commit summaries, and disaster recovery.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}