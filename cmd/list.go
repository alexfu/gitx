package cmd

import (
	"errors"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	"gitx/internal/config"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List currently installed git extensions.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		datadir := config.GitxDataDir()
		if datadir == "" {
			return errors.New("failed to resolve gitx data directory")
		}

		filepath.WalkDir(datadir, func(path string, d fs.DirEntry, err error) error {
			basename := filepath.Base(path)
			if strings.HasPrefix(basename, "git-") && !d.IsDir() {
				result, ok := strings.CutPrefix(path, datadir)
				if ok {
					fmt.Printf("%s\n", result[1:])
				}
			}
			return nil
		})

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
