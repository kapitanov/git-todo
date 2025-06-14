package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/kapitanov/git-todo/internal/application"
)

func pathCommand(c *commandContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "path",
		Short:   "print the path to the git todo data file",
		Long:    "Print the path to the git todo data file used by the application.",
		Example: `  git todo path - prints the path to the git todo data file`,
		Args:    cobra.NoArgs,
	}

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		app, err := application.New()
		if err != nil {
			return c.HandleError(err)
		}

		path := app.Path()
		_, err = fmt.Fprintf(os.Stdout, "%s\n", path)
		return err
	}

	return cmd
}
