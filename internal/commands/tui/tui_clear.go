package tui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/kapitanov/git-todo/internal/application"
)

type clearModel struct {
	app    *application.App
	title  *titleBar
	err    *errorBar
	footer *footerBar
	height int
}

func newClearModel(app *application.App) *clearModel {
	m := &clearModel{
		title: newTitleBar("Clear all TODO items"),
		err:   newErrorBar(),
		footer: newFooterBar(
			footerBarItem{Key: "y", Label: "clear"},
			footerBarItem{Key: "esc", Label: "cancel"},
		),
		app: app,
	}

	return m
}

func (m *clearModel) Init() tea.Cmd { return textinput.Blink }

func (m *clearModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			err := m.app.Clear()
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

func (m *clearModel) View() string {
	rows := []string{
		m.title.View(),
		"",
		"Please confirm deletion of all TODO items:",
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
