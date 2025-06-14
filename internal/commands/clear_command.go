package commands

import (
	"errors"

	"github.com/kapitanov/git-todo/internal/application"
	"github.com/kapitanov/git-todo/internal/commands/cui"
	"github.com/spf13/cobra"
)

func clearCommand(c *commandContext) *cobra.Command {
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
			return c.HandleError(err)
		}

		if !force {
			if c.IsRunningInInteractiveMode() {
				confirmed, err := cui.Confirm("Are you sure you want to remove all TODO items")
				if err != nil {
					return c.HandleError(err)
				}
				if !confirmed {
					return nil
				}
			} else {
				return errors.New("operation is not confirmed")
			}
		}

		err = app.Clear()
		if err != nil {
			return c.HandleError(err)
		}

		c.HumanReadablePrintf("All TODO items have been deleted\n")
		return nil
	}

	return cmd
}
