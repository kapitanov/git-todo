package commands

import (
	"github.com/spf13/cobra"

	"github.com/kapitanov/git-todo/internal/application"
	"github.com/kapitanov/git-todo/internal/commands/cui"
)

func Clear() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "clear",
		Short: "remove all TODO items",
		Long: `Remove all TODO items from the list.

This action cannot be undone!`,
		Example: `  git todo clear         - removes all TODO items from the list
  git todo clear --force - removes all TODO items without confirmation`,

		Args: cobra.NoArgs,
	}

	var force bool
	cmd.Flags().BoolVarP(&force, "force", "f", false, "force remove the TODO item without confirmation")

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		app, err := application.New()
		if err != nil {
			return err
		}

		if !force {
			confirmed, err := cui.Confirm("Are you sure you want to remove all TODO items")
			if err != nil {
				return err
			}
			if !confirmed {
				cmd.PrintErr("Cancelled by user\n")
				return nil
			}
		}

		err = app.Clear()
		if err != nil {
			return err
		}

		cmd.PrintErr("All TODO items have been deleted\n")
		return nil
	}

	return cmd
}
