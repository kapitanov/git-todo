package tui

import (
	"errors"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/kapitanov/git-todo/internal/application"
)

type createModel struct {
	app    *application.App
	input  textinput.Model
	title  *titleBar
	err    *errorBar
	footer *footerBar
	height int
}

func newCreateModel(app *application.App) *createModel {
	m := &createModel{
		input: textinput.New(),
		title: newTitleBar("Edit the title for an new TODO item"),
		err:   newErrorBar(),
		footer: newFooterBar(
			footerBarItem{Key: "enter", Label: "save"},
			footerBarItem{Key: "esc", Label: "cancel"},
		),
		app: app,
	}
	m.input.Prompt = "| "
	m.input.PromptStyle = editPromptStyle
	m.input.SetValue("")
	m.input.Focus()
	m.input.CharLimit = application.MaxTitleLength

	return m
}

func (m *createModel) Init() tea.Cmd { return textinput.Blink }

func (m *createModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.input.Width = msg.Width
		m.title.Width = msg.Width
		m.err.Width = msg.Width
		m.footer.Width = msg.Width
		return m, nil

	case tea.KeyMsg:
		m.err.Error = nil
		switch msg.Type {
		case tea.KeyEnter:
			title := m.input.Value()
			if title != "" {
				_, err := m.app.NewItem(title)
				if err == nil {
					return nil, nil
				}

				m.err.Error = err
			} else {
				m.err.Error = errors.New("title cannot be empty")
			}

		case tea.KeyCtrlC, tea.KeyEsc:
			return nil, nil
		}
	}

	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

func (m *createModel) View() string {
	rows := []string{
		m.title.View(),
		"",
		m.input.View(),
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
