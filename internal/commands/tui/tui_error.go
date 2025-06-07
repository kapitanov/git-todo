package tui

import "fmt"

type errorBar struct {
	Width int
	Error error
}

func newErrorBar() *errorBar {
	return &errorBar{}
}

func (t *errorBar) View() string {
	if t.Error == nil {
		return ""
	}

	str := fmt.Sprintf("Error: %s", t.Error.Error())
	str = padRight(str, t.Width)
	return errorStyle.Render(str)
}
