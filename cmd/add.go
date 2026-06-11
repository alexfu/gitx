package cmd

import (
	"errors"
	"fmt"
	"os"

	"gitx/internal/extensions"

	"github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing/transport"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:          "add",
	Short:        "Add extensions",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]

		// Fetch extension
		path, err := extensions.Download(name)
		if err != nil {
			if errors.Is(err, git.ErrTargetDirNotEmpty) {
				fmt.Fprintln(os.Stderr, "Extension already exists!")
				return nil
			}
			if errors.Is(err, transport.ErrAuthenticationRequired) {
				fmt.Fprintln(os.Stderr, "Extension not found or private.")
				return nil
			}
			return err
		}

		// Install extension
		err = extensions.Install(path)
		if err != nil {
			return err
		}

		fmt.Println("✅ Extension installed!")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
