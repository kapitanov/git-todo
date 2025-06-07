package tui

import (
	"context"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/exp/utf8string"

	"github.com/kapitanov/git-todo/internal/application"
)

var (
	accentColor         = lipgloss.AdaptiveColor{Light: "#8DDBE0", Dark: "#8DDBE0"}
	accentContrastColor = lipgloss.AdaptiveColor{Light: "#1A1A1A", Dark: "#1A1A1A"}
	errorColor          = lipgloss.AdaptiveColor{Light: "#FF0000", Dark: "#FF0000"}

	titleStyle      = lipgloss.NewStyle().Background(accentColor).Foreground(accentContrastColor)
	footerStyle     = lipgloss.NewStyle().Faint(true)
	footerKeyStyle  = lipgloss.NewStyle().Faint(true).Bold(true)
	footerSeparator = " • "
	editPromptStyle = lipgloss.NewStyle().Faint(true)

	nonSelectedListItemStyle          = lipgloss.NewStyle()
	nonSelectedCompletedListItemStyle = lipgloss.NewStyle().Strikethrough(true)
	selectedListItemStyle             = lipgloss.NewStyle().Foreground(accentColor)
	selectedCompletedListItemStyle    = lipgloss.NewStyle().Foreground(accentColor).Strikethrough(true)

	errorStyle = lipgloss.NewStyle().Foreground(errorColor)
)

func Run(ctx context.Context, app *application.App) error {
	m := newListModel(app)
	root := newRootModel(m)

	program := tea.NewProgram(root, tea.WithContext(ctx), tea.WithAltScreen(), tea.WithInputTTY(), tea.WithoutSignalHandler())
	_, err := program.Run()
	if err != nil {
		return err
	}

	return nil
}

func padRight(s string, width int) string {
	n := utf8string.NewString(s).RuneCount()
	if n >= width {
		return s
	}

	s += strings.Repeat(" ", width-n)
	return s
}
