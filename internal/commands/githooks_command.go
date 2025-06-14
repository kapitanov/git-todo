package commands

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/kapitanov/git-todo/internal/application"
	"github.com/kapitanov/git-todo/internal/commands/cui"
	"github.com/kapitanov/git-todo/internal/git"
)

func gitHooksCommand(c *commandContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:    "githooks",
		Hidden: true,
	}

	cmd.AddCommand(preCommitGitHook(c))
	cmd.AddCommand(prePushGitHook(c))

	return cmd
}

func preCommitGitHook(c *commandContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:    "pre-commit",
		Hidden: true,
	}

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		app, err := application.New()
		if err != nil {
			return c.HandleError(err)
		}

		items := app.IncompleteItems()
		if len(items) == 0 {
			return nil
		}

		branch, err := git.CurrentBranch()
		if err != nil {
			return c.HandleError(err)
		}
		if branch != git.Master && branch != git.Main {
			return nil
		}

		gitHookPrintItems(cmd, items)
		cmd.Println()

		confirmMessage := fmt.Sprintf("Are you sure you want to commit these changes to the %q branch", branch)
		confirmed, err := cui.Confirm(confirmMessage)
		if err != nil {
			return c.HandleError(err)
		}
		if !confirmed {
			return errors.New("commit aborted due to incomplete TODO items")
		}
		return nil
	}

	return cmd
}

func prePushGitHook(c *commandContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:    "pre-push",
		Hidden: true,
	}

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		app, err := application.New()
		if err != nil {
			return c.HandleError(err)
		}

		items := app.IncompleteItems()
		if len(items) == 0 {
			return nil
		}

		gitHookPrintItems(cmd, items)
		cmd.Println()

		confirmMessage := "Are you sure you want to push these changes"
		confirmed, err := cui.Confirm(confirmMessage)
		if err != nil {
			return c.HandleError(err)
		}
		if !confirmed {
			return errors.New("push aborted due to incomplete TODO items")
		}
		return nil
	}

	return cmd
}

func gitHookPrintItems(cmd *cobra.Command, items []*application.Item) {
	if len(items) == 1 {
		cmd.Print("You still have a TODO item to resolve:\n")
	} else {
		cmd.Print("You still have some TODO items to resolve:\n")
	}
	for _, item := range items {
		cmd.Printf(" - %s\n", item.Title())
	}
}
