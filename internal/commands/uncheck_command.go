package commands

import (
	"github.com/spf13/cobra"

	"github.com/kapitanov/git-todo/internal/application"
)

func uncheckCommand(c *commandContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "uncheck <index>...",
		Short: "mark a TODO item as incomplete",
		Long: `Mark one or more TODO items as incomplete by their IDs.
You may find the IDs using the "git-todo ls" command.

If an item is already marked as incomplete, this command won't do anything.`,
		Example: `  git-todo check 4e3eeecc                   - marks TODO item with index [4e3eeecc] as incomplete
  git-todo check 4e3eeecc 9612977c ae19ad18 - marks TODO items with indices [4e3eeecc], [9612977c], and [ae19ad18] as incomplete`,
		Args: cobra.MinimumNArgs(1),
	}

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		app, err := application.New()
		if err != nil {
			return c.HandleError(err)
		}

		for item, err := range selectItemsByID(app, args) {
			if err != nil {
				return c.HandleError(err)
			}

			if !item.IsCompleted() {
				c.HumanReadablePrintf("TODO item [%s] %q is not marked as completed\n", item.ID(), item.Title())
				continue
			}

			err = item.SetIsCompleted(false)
			if err != nil {
				return c.HandleError(err)
			}

			c.HumanReadablePrintf("TODO item [%s] %q has been marked as incomplete\n", item.ID(), item.Title())
			c.MachineReadablePrintf("%s\n", item.ID())
		}
		return nil
	}

	return cmd
}
