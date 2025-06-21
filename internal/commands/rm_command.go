package commands

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/kapitanov/git-todo/internal/application"
	"github.com/kapitanov/git-todo/internal/commands/cui"
)

func removeCommand(c *commandContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rm",
		Short: "remove a TODO item",
		Long: `Remove a TODO item by its ID.
You may find the IDs using the "git todo ls" command.

This action cannot be undone!

This command requires a confirmation before removing the item.
If you want to skip the confirmation, use the --force flag.
`,
		Example: `  git todo rm 4e3eeecc                   - removes TODO item with index [4e3eeecc]
  git todo rm 4e3eeecc 9612977c ae19ad18 - removes TODO items [4e3eeecc], [9612977c], and [ae19ad18]
  git todo rm --force 4e3eeecc           - removes TODO items [4e3eeecc] without confirmation`,
		Args: cobra.MinimumNArgs(1),
	}

	var force bool
	cmd.Flags().BoolVarP(&force, "force", "f", false, "force remove the TODO item without confirmation")

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		if !force && !c.IsRunningInInteractiveMode() {
			return errors.New("this command requires either an interactive confirmation prompt or a \"--force\" flag")
		}

		app, err := application.New()
		if err != nil {
			return c.HandleError(err)
		}

		for item, err := range selectItemsByID(app, args) {
			if err != nil {
				return c.HandleError(err)
			}

			if !force {
				confirmed, err := cui.Confirm(fmt.Sprintf("Are you sure you want to remove TODO item [%s] %q", item.ID(), item.Title()))
				if err != nil {
					return c.HandleError(err)
				}
				if !confirmed {
					c.HumanReadablePrintf("Canceled removal of TODO item [%s] %q\n", item.ID(), item.Title())
					continue
				}
			}

			err = item.Delete()
			if err != nil {
				return c.HandleError(err)
			}

			c.HumanReadablePrintf("TODO item [%s] %q has been removed\n", item.ID(), item.Title())
			c.MachineReadablePrintf("%s\n", item.ID())
		}
		return nil
	}

	return cmd
}
