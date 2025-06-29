package tui

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/kapitanov/git-todo/internal/application"
)

type listItemWrapper struct {
	Item *application.Item
}

func (w *listItemWrapper) FilterValue() string { return w.Item.Title() }

func wrapListItems(app *application.App) []list.Item {
	appItems := app.Items()

	listItems := make([]list.Item, len(appItems))
	for i, item := range appItems {
		listItems[i] = &listItemWrapper{Item: item}
	}
	return listItems
}

type listModelDelegate struct{}

func (*listModelDelegate) Height() int                             { return 1 }
func (*listModelDelegate) Spacing() int                            { return 0 }
func (*listModelDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (*listModelDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(*listItemWrapper)
	if !ok {
		return
	}

	isSelected := m.Index() == index
	str := formatListItem(i.Item, isSelected)

	_, _ = fmt.Fprint(w, str)
}

func formatListItem(item *application.Item, isSelected bool) string {
	var selector, marker, text string

	if isSelected {
		selector = "→ "
	} else {
		selector = "  "
	}

	if item.IsCompleted() {
		marker = " [✓] "
	} else {
		marker = " [ ] "
	}

	id := fmt.Sprintf("%s ", item.ID())
	text = strings.TrimSpace(item.Title())

	var selectorStyle, markerStyle, textStyle lipgloss.Style

	if !isSelected {
		if item.IsCompleted() {
			selectorStyle = nonSelectedListItemStyle
			markerStyle = nonSelectedListItemStyle
			textStyle = nonSelectedCompletedListItemStyle
		} else {
			selectorStyle = nonSelectedListItemStyle
			markerStyle = nonSelectedListItemStyle
			textStyle = nonSelectedListItemStyle
		}
	} else {
		if item.IsCompleted() {
			selectorStyle = selectedListItemStyle
			markerStyle = selectedListItemStyle
			textStyle = selectedCompletedListItemStyle
		} else {
			selectorStyle = selectedListItemStyle
			markerStyle = selectedListItemStyle
			textStyle = selectedListItemStyle
		}
	}

	idStyle := textStyle.Faint(true)
	
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		selectorStyle.Render(selector),
		markerStyle.Render(marker),
		idStyle.Render(id),
		textStyle.Render(text),
	)
}
