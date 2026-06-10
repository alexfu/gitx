/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"gitx/internal/config"
	"gitx/internal/repo"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:          "add",
	Short:        "A brief description of your command",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]

		// Fetch extension
		extension, err := repo.DownloadExtention(name)
		if err != nil {
			return err
		}

		// Save extension
		gitxhome, err := config.EnsureGitxHome()
		if err != nil {
			return err
		}

		filename := filepath.Join(gitxhome, fmt.Sprintf("git-%s", name))
		err = os.WriteFile(filename, extension, 0o644)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
