package tui

import tea "github.com/charmbracelet/bubbletea"

type rootModel struct {
	viewStack     []tea.Model
	width, height int
}

func newRootModel(view tea.Model) *rootModel {
	return &rootModel{
		viewStack: []tea.Model{view},
	}
}

func (m *rootModel) Init() tea.Cmd {
	view := m.currentView()
	if view == nil {
		return tea.Quit
	}

	return view.Init()
}

func (m *rootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if msg, ok := msg.(tea.WindowSizeMsg); ok {
		m.width, m.height = msg.Width, msg.Height
	}

	view := m.currentView()
	if view == nil {
		return m, tea.Quit
	}

	nextView, cmd := view.Update(msg)
	if nextView == nil {
		return m.popView()
	}

	if nextView == view {
		return m, cmd
	}

	return m.pushView(nextView, cmd)
}

func (m *rootModel) View() string {
	view := m.currentView()
	if view == nil {
		return ""
	}

	return view.View()
}

func (m *rootModel) currentView() tea.Model {
	if len(m.viewStack) == 0 {
		return nil
	}

	return m.viewStack[len(m.viewStack)-1]
}

func (m *rootModel) popView() (tea.Model, tea.Cmd) {
	if len(m.viewStack) < 2 {
		m.viewStack = m.viewStack[:0]
		return m, tea.Quit
	}

	view := m.viewStack[len(m.viewStack)-2]
	m.viewStack = m.viewStack[:len(m.viewStack)-1]

	view.Init()
	cmd := func() tea.Msg {
		return tea.WindowSizeMsg{Width: m.width, Height: m.height}
	}

	return m, cmd
}

func (m *rootModel) pushView(view tea.Model, cmd tea.Cmd) (tea.Model, tea.Cmd) {
	m.viewStack = append(m.viewStack, view)

	view.Init()
	initCmd := func() tea.Msg {
		return tea.WindowSizeMsg{Width: m.width, Height: m.height}
	}

	return m, tea.Batch(initCmd, cmd)
}
