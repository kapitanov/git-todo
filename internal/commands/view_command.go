package commands

import (
	"os"

	"github.com/kapitanov/git-todo/internal/application"
	"github.com/spf13/cobra"
)

func viewCommand(c *commandContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "view <id>...",
		Short: "view a TODO item",
		Long: `Print a TODO item from the current Git repository.
This command displays a selected TODO item in the current Git repository, showing its ID, completion status, and title.

By default, TODO items are displayed in pretty, human-readable format,
but you can customize the output using various flags, as shown in the examples below.
`,
		Example: `  git todo view 4e3eeecc                    - print a TODO item
  git todo view 4e3eeecc 9612977c ae19ad18  - print few TODO items
  git todo view --json 4e3eeecc             - print a TODO item in JSON format
  git todo view --plain 4e3eeecc            - print a TODO item in a plain, script-friendly (space-separated) format`,
		Args: cobra.MinimumNArgs(1),
	}

	render := configureItemsRenderer(cmd)

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		app, err := application.New()
		if err != nil {
			return c.HandleError(err)
		}

		var items []*application.Item
		for item, err := range selectItemsByID(app, args) {
			if err != nil {
				return c.HandleError(err)
			}
			items = append(items, item)
		}

		out, err := render(items, c)
		if err != nil {
			return c.HandleError(err)
		}

		_, err = os.Stdout.Write([]byte(out))
		return err
	}

	return cmd
}
