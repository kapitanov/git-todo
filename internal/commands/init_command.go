package commands

import (
	"github.com/spf13/cobra"

	"github.com/kapitanov/git-todo/internal/application"
	"github.com/kapitanov/git-todo/internal/installer"
)

func Init() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "install git todo hooks",
		Long: `Install git todo hooks in the current Git repository.
This command sets up the necessary hooks to integrate git todo functionality into your Git workflow:

  - pre-commit hook: to ensure TODO items are completed before commiting changes to the main branch.
    This hook prevents commits if there are any TODO items that are not marked as completed
    and the current branch is the main branch ("main" or "master"),
    but you still can commit your changes - just confirm the changes when the hook asks you to.

  - post-commit hook: to check for TODO items before committing.
    This hook prints a warning if there are any TODO items.

  - pre-push hook: to ensure TODO items are checked before pushing changes.
    This hook prevents pushing if there are any TODO items that are not marked as completed,
    but you still can push your changes - just confirm the changes when the hook asks you to.
`,
		Example: `  git todo init - installs git todo hooks in the current repository`,
		Args:    cobra.NoArgs,
	}

	var force bool
	cmd.Flags().BoolVarP(&force, "force", "f", false, "force githooks installation, even if hooks are already installed (will overwrite existing hooks)")

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		app, err := application.New()
		if err != nil {
			return err
		}

		err = installer.Install(app.RepositoryRoot(), force)
		if err != nil {
			return err
		}

		cmd.PrintErr("Git hooks installed successfully.\n")
		return nil
	}

	return cmd
}
