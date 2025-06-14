package model

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/rs/zerolog/log"
)

type Model struct {
	Items []*Item
}

type Item struct {
	IsCompleted bool
	Title       string
}

func newEmptyModel() *Model {
	return &Model{
		Items: []*Item{},
	}
}

func Load(path string) (*Model, error) {
	bs, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			log.Info().Str("path", path).Msg("file does not exist, falling back to empty model")
			return newEmptyModel(), nil // Return an empty model if the file does not exist
		}

		log.Error().Err(err).Str("path", path).Msg("failed to read model file")
		return nil, fmt.Errorf("failed to read file %q: %w", path, err)
	}

	m, err := parse(bs)
	if err != nil {
		log.Error().Err(err).Str("path", path).Msg("failed to parse model file")
		return nil, fmt.Errorf("failed to parse file %q: %w", path, err)
	}

	log.Debug().Err(err).Str("path", path).Msg("loaded model file")
	return m, nil
}

var (
	lineRegex = regexp.MustCompile(`^\s*(\[(.)\])\s*(.*)$`)
)

func parse(bs []byte) (*Model, error) {
	model := newEmptyModel()

	for line := range strings.Lines(string(bs)) {
		line = strings.TrimSpace(line)
		m := lineRegex.FindStringSubmatch(line)
		if len(m) == 0 {
			continue
		}

		isCompleted := m[2] != " " && m[2] != "_"
		title := strings.TrimSpace(m[3])

		item := &Item{
			IsCompleted: isCompleted,
			Title:       title,
		}
		model.Items = append(model.Items, item)
	}

	return model, nil
}

func (m *Model) Store(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file %q: %w", path, err)
	}

	defer func() { _ = f.Close() }()

	str := m.stringify()
	_, err = f.Write([]byte(str))
	return err
}

func (m *Model) stringify() string {
	var sb strings.Builder

	for _, item := range m.Items {
		if item.IsCompleted {
			sb.WriteString("[x] ")
		} else {
			sb.WriteString("[ ] ")
		}

		sb.WriteString(item.Title)
		sb.WriteString("\n")
	}

	str := sb.String()
	return str
}
