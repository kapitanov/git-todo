package commands

import (
	"github.com/spf13/cobra"

	"github.com/kapitanov/git-todo/internal/application"
)

func Check() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "check <index>...",
		Short: "mark a TODO item as completed",
		Long: `Mark one or more TODO items as completed by their index in the list.
You may find the indices using the "git todo ls" command.`,
		Example: `  git todo check 5     - marks TODO item with index 5 as completed
  git todo check 1 3 5 - marks TODO items with indices 1, 3, and 5 as completed`,
		Args: cobra.MinimumNArgs(1),
	}

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		app, err := application.New()
		if err != nil {
			return err
		}

		items, err := selectItemsByIndex(app, args)
		if err != nil {
			return err
		}

		for _, item := range items {
			if !item.IsCompleted() {
				err = item.SetIsCompleted(true)
				if err != nil {
					return err
				}

				cmd.PrintErrf("TODO item #%d has been checked as completed (%s)\n", item.ID(), item.Title())
			}
		}
		return nil
	}

	return cmd
}
