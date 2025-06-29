package model

import (
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
)

type Model struct {
	Items []*Item `yaml:"items"`
}

type Item struct {
	ID          string `yaml:"id"`
	IsCompleted bool   `yaml:"done,omitempty"`
	Title       string `yaml:"title"`
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
		return newEmptyModel(), nil
	}

	log.Debug().Err(err).Str("path", path).Msg("loaded model file")
	return m, nil
}

func parse(bs []byte) (*Model, error) {
	model := newEmptyModel()
	err := yaml.Unmarshal(bs, model)
	return model, err
}

func (m *Model) Store(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file %q: %w", path, err)
	}

	defer func() { _ = f.Close() }()

	bs, err := yaml.Marshal(m)
	if err != nil {
		return fmt.Errorf("failed to marshal model: %w", err)
	}

	_, err = f.Write(bs)
	return err
}
