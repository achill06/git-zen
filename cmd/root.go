package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git-zen",
	Short: "A blazing-fast TUI for git branch management",
	Long:  `git-zen - An open-source CLI extension offering an interactive terminal UI  for seamless branch switching and asynchronous PR status tracking.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
