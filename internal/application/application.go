package application

import (
	"errors"
	"slices"
	"strings"

	"golang.org/x/exp/utf8string"

	"github.com/kapitanov/git-todo/internal/application/discover"
	"github.com/kapitanov/git-todo/internal/application/model"
)

const (
	MaxTitleLength = 256
)

type App struct {
	model          *model.Model
	repositoryRoot string
	dataFilePath   string
}

func New() (*App, error) {
	repositoryRoot, dataFilePath, err := discover.TODOPath()
	if err != nil {
		return nil, err
	}

	m, err := model.Load(dataFilePath)
	if err != nil {
		return nil, err
	}

	return &App{model: m, repositoryRoot: repositoryRoot, dataFilePath: dataFilePath}, nil
}

func (app *App) RepositoryRoot() string { return app.repositoryRoot }
func (app *App) Path() string           { return app.dataFilePath }

func (app *App) Item(i int) *Item {
	if i < 1 || i > len(app.model.Items) {
		return nil
	}

	return &Item{
		id:    i,
		model: app.model.Items[i-1],
		app:   app,
	}
}

func (app *App) Items() []*Item {
	items := make([]*Item, len(app.model.Items))
	for i, item := range app.model.Items {
		items[i] = &Item{
			id:    i + 1,
			model: item,
			app:   app,
		}
	}
	return items
}

func (app *App) IncompleteItems() []*Item {
	var incompleteItems []*Item
	for _, item := range app.Items() {
		if !item.IsCompleted() {
			incompleteItems = append(incompleteItems, item)
		}
	}
	return incompleteItems
}

func (app *App) NewItem(title string) (*Item, error) {
	title, err := validateTitle(title)
	if err != nil {
		return nil, err
	}

	for _, item := range app.model.Items {
		if item.Title == title {
			return nil, errors.New("item already exists")
		}
	}

	item := &model.Item{
		Title:       title,
		IsCompleted: false,
	}
	app.model.Items = append(app.model.Items, item)
	err = app.save()
	if err != nil {
		return nil, err
	}

	items := app.Items()
	return items[len(items)-1], nil
}

func (app *App) Clear() error {
	app.model.Items = app.model.Items[:0]
	return app.save()
}

func (app *App) save() error {
	return app.model.Store(app.dataFilePath)
}

func (app *App) delete(item *Item) error {
	app.model.Items = slices.DeleteFunc(app.model.Items, func(i *model.Item) bool { return i == item.model })
	return app.save()
}

type Item struct {
	id    int
	model *model.Item
	app   *App
}

func (item *Item) ID() int           { return item.id }
func (item *Item) Title() string     { return item.model.Title }
func (item *Item) IsCompleted() bool { return item.model.IsCompleted }

func (item *Item) SetTitle(val string) error {
	val, err := validateTitle(val)
	if err != nil {
		return err
	}

	item.model.Title = val
	return item.app.save()
}

func (item *Item) SetIsCompleted(val bool) error {
	item.model.IsCompleted = val
	return item.app.save()
}

func (item *Item) Delete() error {
	return item.app.delete(item)
}

func validateTitle(title string) (string, error) {
	title = strings.TrimSpace(title)

	if title == "" {
		return "", errors.New("title cannot be empty")
	}

	titleStr := utf8string.NewString(title)

	if titleStr.RuneCount() > MaxTitleLength {
		title = titleStr.Slice(0, MaxTitleLength)
	}

	return title, nil
}
