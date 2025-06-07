package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/kapitanov/git-todo/internal/application"
)

type deleteModel struct {
	item   *application.Item
	title  *titleBar
	err    *errorBar
	footer *footerBar
	height int
}

func newDeleteModel(item *application.Item) *deleteModel {
	m := &deleteModel{
		title: newTitleBar("Delete a TODO item"),
		err:   newErrorBar(),
		footer: newFooterBar(
			footerBarItem{Key: "y", Label: "delete"},
			footerBarItem{Key: "esc", Label: "cancel"},
		),
		item: item,
	}

	return m
}

func (m *deleteModel) Init() tea.Cmd { return textinput.Blink }

func (m *deleteModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.title.Width = msg.Width
		m.err.Width = msg.Width
		m.footer.Width = msg.Width
		return m, nil

	case tea.KeyMsg:
		m.err.Error = nil
		switch keypress := msg.String(); keypress {
		case "y":
			err := m.item.Delete()
			if err == nil {
				return nil, nil
			}

			m.err.Error = err

		case "q", "ctrl+c", "esc":
			return nil, nil
		}
	}

	return m, cmd
}

func (m *deleteModel) View() string {
	rows := []string{
		m.title.View(),
		"",
		"Please confirm deletion of the TODO item:",
		fmt.Sprintf("| %s", m.item.Title()),
		"",
		m.err.View(),
		"",
	}

	for len(rows) < m.height-1 {
		rows = append(rows, "")
	}

	rows = append(rows, m.footer.View())
	return lipgloss.JoinVertical(lipgloss.Top, rows...)
}
