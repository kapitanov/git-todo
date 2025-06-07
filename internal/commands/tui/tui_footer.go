package tui

import (
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/exp/utf8string"
)

type footerBar struct {
	Width int
	Items []footerBarItem
}

func newFooterBar(items ...footerBarItem) *footerBar {
	return &footerBar{
		Items: items,
	}
}

type footerBarItem struct {
	Key   string
	Label string
}

func (i footerBarItem) measure() int {
	var width int
	if i.Key != "" {
		width += utf8string.NewString(i.Key).RuneCount() + 1 // +1 for the space after the key
	}

	width += utf8string.NewString(i.Label).RuneCount()
	return width
}

func (f *footerBar) View() string {
	width := f.measure()

	var items []string
	for i, item := range f.Items {
		if i > 0 {
			items = append(items, footerStyle.Render(footerSeparator))
		}

		if item.Key != "" {
			items = append(items, footerKeyStyle.Render(item.Key)+" ")
		}

		items = append(items, footerStyle.Render(item.Label))
	}

	items = append(items, footerStyle.Render(padRight("", f.Width-width)))
	return lipgloss.JoinHorizontal(lipgloss.Left, items...)
}

func (f *footerBar) measure() int {
	var width int

	for i, item := range f.Items {
		if i > 0 {
			width += 3 // for the separator
		}

		width += item.measure()
	}

	return width
}
