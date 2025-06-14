package commands

import (
	"github.com/spf13/cobra"

	"github.com/kapitanov/git-todo/internal/application"
	"github.com/kapitanov/git-todo/internal/installer"
)

func deinitCommand(c *commandContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deinit",
		Short: "uninstall git todo hooks",
		Long: `Uninstall git todo hooks from the current Git repository.
This command removes the git todo hooks that were installed in the current Git repository.`,
		Example: `  git todo deinit - removes git todo hooks from the current repository`,
		Args:    cobra.NoArgs,
	}

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		app, err := application.New()
		if err != nil {
			return c.HandleError(err)
		}

		err = installer.Uninstall(app.RepositoryRoot())
		if err != nil {
			return c.HandleError(err)
		}

		c.HumanReadablePrintf("Git hooks uninstalled successfully.\n")
		return nil
	}

	return cmd
}
