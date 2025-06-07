package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/paginator"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/kapitanov/git-todo/internal/application"
)

type listModel struct {
	app    *application.App
	list   list.Model
	title  *titleBar
	footer *footerBar
}

func newListModel(app *application.App) *listModel {
	const (
		defaultWidth = 50
		listHeight   = 10
	)
	m := &listModel{
		app:   app,
		list:  list.New(nil, &listModelDelegate{}, defaultWidth, listHeight),
		title: newTitleBar("Git TODO"),
		footer: newFooterBar(
			footerBarItem{Key: "↑/k", Label: "up"},
			footerBarItem{Key: "↓/j", Label: "down"},
			footerBarItem{Key: "n", Label: "new item"},
			footerBarItem{Key: "space/t", Label: "toggle item"},
			footerBarItem{Key: "enter/e", Label: "edit item"},
			footerBarItem{Key: "d", Label: "delete item"},
			footerBarItem{Key: "x", Label: "clear all items"},
			footerBarItem{Key: "q", Label: "quit"},
		),
	}

	m.list.Styles.ArabicPagination = footerStyle

	m.list.SetShowTitle(false)
	m.list.SetShowStatusBar(false)
	m.list.SetShowHelp(false)
	m.list.Paginator.Type = paginator.Arabic
	m.list.SetFilteringEnabled(false)

	return m
}

func (m *listModel) updateItems() {
	items := wrapListItems(m.app)
	m.list.SetItems(items)

	var total, completed int
	for _, item := range items {
		if i, ok := item.(*listItemWrapper); ok {
			total++
			if i.Item.IsCompleted() {
				completed++
			}
		}
	}

	m.title.Subtitle = fmt.Sprintf("%d/%d items completed", completed, total)
}

func (m *listModel) Init() tea.Cmd {
	m.updateItems()
	return nil
}

func (m *listModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.title.Width = msg.Width
		m.footer.Width = msg.Width
		m.list.SetWidth(msg.Width)
		m.list.SetHeight(msg.Height - 4)
		return m, nil

	case tea.KeyMsg:
		res, cmd := m.onKeyPress(msg)
		if res != nil {
			return res, cmd
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m *listModel) onKeyPress(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch keypress := msg.String(); keypress {
	case "q", "ctrl+c":
		return nil, tea.Quit

	case " ", "t":
		err := m.toggleItem()
		if err != nil {
			return nil, tea.Quit
		}
		m.updateItems()
		return m, nil

	case "e", "enter":
		i, ok := m.list.SelectedItem().(*listItemWrapper)
		if ok {
			return newEditModel(i.Item), nil
		}

	case "n":
		return newCreateModel(m.app), nil

	case "x":
		return newClearModel(m.app), nil

	case "d":
		i, ok := m.list.SelectedItem().(*listItemWrapper)
		if ok {
			return newDeleteModel(i.Item), nil
		}
	}

	return nil, nil
}

func (m *listModel) toggleItem() error {
	i, ok := m.list.SelectedItem().(*listItemWrapper)
	if !ok {
		return nil
	}

	return i.Item.SetIsCompleted(!i.Item.IsCompleted())
}

func (m *listModel) View() string {
	return lipgloss.JoinVertical(lipgloss.Top, m.title.View(), "", m.list.View(), "", m.footer.View())
}
