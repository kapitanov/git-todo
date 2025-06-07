package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/kapitanov/git-todo/internal/application"
	"github.com/kapitanov/git-todo/internal/commands/cui"
)

func List() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ls",
		Short: "list TODO items",
		Long: `List all TODO items in the current Git repository.
This command displays all TODO items in the current Git repository, showing their index, completion status, and title.

By default, TODO items are displayed in pretty, human-readable format,
but you can customize the output using various flags, as shown in the examples below.
`,
		Example: `  git todo ls              - lists all TODO items in the current repository
  git todo ls --completed  - lists only completed TODO items
  git todo ls --incomplete - lists only incomplete TODO items
  git todo ls --json       - lists TODO items in JSON format
  git todo ls --plain      - lists TODO items in a plain, script-friendly (space-separated) format
`,
		Args: cobra.NoArgs,
	}

	filter := configureItemsFilter(cmd)
	render := configureItemsRenderer(cmd)

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		app, err := application.New()
		if err != nil {
			return err
		}

		items, err := filter(app)
		if err != nil {
			return err
		}

		out, err := render(items)
		if err != nil {
			return err
		}

		_, err = os.Stdout.Write([]byte(out))
		if err != nil {
			return err
		}
		return nil
	}

	return cmd
}

func configureItemsFilter(cmd *cobra.Command) func(app *application.App) ([]*application.Item, error) {
	var (
		completedOnly, incompleteOnly bool
	)
	cmd.Flags().BoolVarP(&completedOnly, "completed", "c", false, "show only completed TODO items")
	cmd.Flags().BoolVarP(&incompleteOnly, "incomplete", "i", false, "show only incomplete TODO items")

	return func(app *application.App) ([]*application.Item, error) {
		if !onlyOneAllowed(completedOnly, incompleteOnly) {
			return nil, errors.New("cannot use both \"--completed\" and \"--incomplete\" flags at the same time")
		}

		if !completedOnly && !incompleteOnly {
			return app.Items(), nil
		}

		var items []*application.Item
		for _, item := range app.Items() {
			if completedOnly && item.IsCompleted() {
				items = append(items, item)
			} else if incompleteOnly && !item.IsCompleted() {
				items = append(items, item)
			}
		}

		return items, nil
	}
}

func configureItemsRenderer(cmd *cobra.Command) func(items []*application.Item) (string, error) {
	var (
		printJSON, printPlain bool
	)
	cmd.Flags().BoolVarP(&printJSON, "json", "j", false, "print TODO items in JSON format")
	cmd.Flags().BoolVarP(&printPlain, "plain", "p", false, "print TODO items in the plain format")

	return func(items []*application.Item) (string, error) {
		if !onlyOneAllowed(printJSON, printPlain) {
			return "", errors.New("cannot use \"--json\" and \"--plain\" flags at the same time")
		}

		if printJSON {
			return renderJSONList(items)
		}

		if printPlain {
			return renderPlainList(items), nil
		}

		return renderPrettyList(items), nil
	}
}

func renderPrettyList(items []*application.Item) string {
	if len(items) == 0 {
		return "No TODO items found.\n"
	}

	maxIndex := len(items)
	n := numDecimalPlaces(maxIndex)
	itemNumberFormat := fmt.Sprintf("%%%dd", n)

	var sb strings.Builder
	for _, item := range items {
		checkBox := "·"
		style := cui.ItemTextStyle
		if item.IsCompleted() {
			checkBox = "✓"
			style = cui.ItemCompletedTextStyle
		}

		itemNumber := cui.ItemIndexStyle.Render(fmt.Sprintf(itemNumberFormat, item.ID()))
		checkBox = cui.ItemCheckboxStyle.Render(checkBox)
		title := style.Render(item.Title())

		sb.WriteString(fmt.Sprintf("%s %s %s\n", itemNumber, checkBox, title))
	}

	return sb.String()
}

func renderPlainList(items []*application.Item) string {
	if len(items) == 0 {
		return "\n"
	}

	var sb strings.Builder
	for _, item := range items {
		checkBox := "X"
		if item.IsCompleted() {
			checkBox = "_"
		}

		sb.WriteString(fmt.Sprintf("%d %s %s\n", item.ID(), checkBox, item.Title()))
	}

	return sb.String()
}

func renderJSONList(items []*application.Item) (string, error) {
	type (
		Item struct {
			ID        int    `json:"id"`
			Completed bool   `json:"completed"`
			Title     string `json:"title"`
		}

		ItemList []Item
	)

	list := make(ItemList, 0, len(items))
	for _, item := range items {
		list = append(list, Item{
			ID:        item.ID(),
			Completed: item.IsCompleted(),
			Title:     item.Title(),
		})
	}

	bs, err := json.MarshalIndent(list, "", "    ")
	return string(bs) + "\n", err
}

func onlyOneAllowed(xs ...bool) bool {
	if len(xs) == 0 {
		return false
	}

	var count int
	for _, x := range xs {
		if x {
			count++
		}
	}

	return count <= 1
}

func numDecimalPlaces(num int) int {
	if num == 0 {
		return 1
	}
	count := 0
	for num > 0 {
		num /= 10
		count++
	}
	return count
}
