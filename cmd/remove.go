package cmd

import (
	"fmt"
	"os"

	"gitx/internal/extensions"

	"charm.land/huh/v2"
	"charm.land/lipgloss/v2/list"
	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove [extension]",
	Short: "Remove git extensions.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		extension := args[0]

		removeInfo := extensions.Remove(extension, func(deleteList []string) bool {
			var confirmed bool
			form := huh.NewForm(
				huh.NewGroup(
					huh.NewNote().
						Title("This will remove the following").
						Description(list.New(deleteList).String()),
					huh.NewConfirm().
						Title("Continue?").
						Affirmative("Yes").
						Negative("No").
						Value(&confirmed),
				),
			)
			form.Run()
			return confirmed
		})

		if len(removeInfo.Failed) > 0 {
			fmt.Fprintf(os.Stdout, "Deleted %d files. The following files failed to remove\n%s", len(removeInfo.Success), list.New(removeInfo.Failed))
			return
		}

		if len(removeInfo.Success) > 0 {
			fmt.Fprintf(os.Stdout, "Successfully removed %s\n", extension)
		} else if len(removeInfo.DeleteList) == 0 {
			fmt.Fprintf(os.Stdout, "Extension %s not found\n", extension)
		}
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
