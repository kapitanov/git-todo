package tui

type titleBar struct {
	Width    int
	Title    string
	Subtitle string
}

func newTitleBar(title string) *titleBar {
	return &titleBar{
		Title: title,
	}
}

func (t *titleBar) View() string {
	if t.Subtitle == "" {
		str := padRight(t.Title, t.Width)
		return titleStyle.Render(str)
	} else {
		left, right := t.Title, t.Subtitle
		left = padRight(left, t.Width-len(right))

		str := left + right
		return titleStyle.Render(str)
	}
}
