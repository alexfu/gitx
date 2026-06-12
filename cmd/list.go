package cmd

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"gitx/internal/config"

	"charm.land/lipgloss/v2/list"
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

		var extensions []string
		filepath.WalkDir(datadir, func(path string, d fs.DirEntry, err error) error {
			basename := filepath.Base(path)
			if strings.HasPrefix(basename, "git-") && !d.IsDir() {
				result, ok := strings.CutPrefix(path, datadir)
				if ok {
					extensions = append(extensions, result[1:])
				}
			}
			return nil
		})

		if len(extensions) > 0 {
			fmt.Fprintf(os.Stdout, "%s\n", list.New(extensions))
		} else {
			fmt.Fprintf(os.Stdout, "No extensions\n")
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
