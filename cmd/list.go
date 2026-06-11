package cmd

import (
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
	Run: func(cmd *cobra.Command, args []string) {
		datadir := config.GitxDataDir()
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
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
