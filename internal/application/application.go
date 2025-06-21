package application

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/kapitanov/git-todo/internal/application/discover"
	"github.com/kapitanov/git-todo/internal/application/idgen"
	"github.com/kapitanov/git-todo/internal/application/model"
	"golang.org/x/exp/utf8string"
)

const (
	MaxTitleLength = 256
)

type ItemAlreadyExistsError struct {
	Item *Item
}

func (e ItemAlreadyExistsError) Error() string {
	return fmt.Sprintf("item %q already exists", e.Item.Title())
}

type App struct {
	repositoryRoot string
	dataFilePath   string
	items          []*Item
	byID           map[string]*Item
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

	app := &App{
		repositoryRoot: repositoryRoot,
		dataFilePath:   dataFilePath,
		items:          make([]*Item, 0, len(m.Items)),
		byID:           make(map[string]*Item, len(m.Items)),
	}

	for _, item := range m.Items {
		i := newItem(item, app)
		app.items = append(app.items, i)
		app.byID[item.ID] = i
	}

	return app, nil
}

func (app *App) RepositoryRoot() string { return app.repositoryRoot }
func (app *App) Path() string           { return app.dataFilePath }

func (app *App) Item(id string) *Item {
	item, exists := app.byID[id]
	if exists {
		return item
	}

	for _, i := range app.items {
		if strings.HasPrefix(i.ID(), id) {
			if item != nil {
				// This partial ID matches to more than one item.
				return nil
			}
			item = i
		}
	}

	return item
}

func (app *App) Items() []*Item { return app.items }

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

	for _, item := range app.items {
		if item.Title() == title {
			return nil, ItemAlreadyExistsError{Item: item}
		}
	}

	var id string
	for id = range idgen.Generate(title) {
		if _, exists := app.byID[id]; !exists {
			break
		}
	}

	item := newItem(&model.Item{ID: id, Title: title, IsCompleted: false}, app)
	app.items = append(app.items, item)
	app.byID[id] = item
	err = app.save()
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (app *App) FindItem(title string) *Item {
	for _, item := range app.Items() {
		if item.Title() == title {
			return item
		}
	}

	return nil
}

func (app *App) ClearItems() error {
	app.items = make([]*Item, 0)
	app.byID = make(map[string]*Item)
	return app.save()
}

func (app *App) save() error {
	m := &model.Model{
		Items: make([]*model.Item, 0, len(app.items)),
	}
	for _, item := range app.items {
		m.Items = append(m.Items, item.item)
	}

	return m.Store(app.dataFilePath)
}

func (app *App) delete(item *Item) error {
	delete(app.byID, item.ID())
	app.items = slices.DeleteFunc(app.items, func(i *Item) bool { return i == item })

	return app.save()
}

type Item struct {
	item *model.Item
	app  *App
}

func newItem(item *model.Item, app *App) *Item {
	return &Item{
		item: item,
		app:  app,
	}
}

func (item *Item) ID() string        { return item.item.ID }
func (item *Item) Title() string     { return item.item.Title }
func (item *Item) IsCompleted() bool { return item.item.IsCompleted }

func (item *Item) SetTitle(val string) error {
	val, err := validateTitle(val)
	if err != nil {
		return err
	}

	existingItem := item.app.FindItem(val)
	if existingItem != nil && existingItem.ID() != item.ID() {
		return ItemAlreadyExistsError{Item: existingItem}
	}

	item.item.Title = val
	return item.app.save()
}

func (item *Item) SetIsCompleted(val bool) error {
	item.item.IsCompleted = val
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
