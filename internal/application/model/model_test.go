package model

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/kapitanov/git-todo/internal/logutil"
)

func TestLoad(t *testing.T) {
	type TestCase struct {
		Name     string
		Path     string
		Expected *Model
	}

	testCases := []TestCase{
		{
			Name: "Valid file",
			Path: "testdata/valid.yaml",
			Expected: &Model{
				Items: []*Item{
					{ID: "a", IsCompleted: false, Title: "Item 1"},
					{ID: "b", IsCompleted: false, Title: "Item 2"},
					{ID: "c", IsCompleted: true, Title: "Item 3"},
					{ID: "d", IsCompleted: false, Title: "Item 4"},
					{ID: "e", IsCompleted: true, Title: "Item 5"},
				},
			},
		},
		{
			Name: "Invalid file",
			Path: "testdata/invalid.yaml",
			Expected: &Model{
				Items: []*Item{},
			},
		},
		{
			Name: "Empty file",
			Path: "testdata/empty.yaml",
			Expected: &Model{
				Items: []*Item{},
			},
		},
		{
			Name: "Non existing file",
			Path: "testdata/non-existing.txt",
			Expected: &Model{
				Items: []*Item{},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.Name, func(t *testing.T) {
			logutil.WithTestLogger(t, func() {
				actual, err := Load(tc.Path)
				require.NoError(t, err)

				assert.Equal(t, tc.Expected, actual)
			})
		})
	}
}

func TestStore(t *testing.T) {
	type TestCase struct {
		Name     string
		Model    *Model
		Expected string
	}

	testCases := []TestCase{
		{
			Name: "Valid model",
			Model: &Model{
				Items: []*Item{
					{ID: "a", IsCompleted: false, Title: "Item 1"},
					{ID: "b", IsCompleted: true, Title: "Item 2"},
				},
			},
			Expected: `items:
    - id: a
      title: Item 1
    - id: b
      done: true
      title: Item 2
`,
		},
		{
			Name:     "Empty model",
			Model:    &Model{Items: []*Item{}},
			Expected: "items: []\n",
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.Name, func(t *testing.T) {
			logutil.WithTestLogger(t, func() {
				f, err := os.CreateTemp("", "model_test_*.txt")
				require.NoError(t, err)
				defer func() { _ = os.Remove(f.Name()) }()

				path := f.Name()
				_ = f.Close()

				err = tc.Model.Store(path)
				require.NoError(t, err)

				bs, err := os.ReadFile(path)
				require.NoError(t, err)

				assert.Equal(t, tc.Expected, string(bs))
			})
		})
	}
}

func TestStoreLoad(t *testing.T) {
	type TestCase struct {
		Name  string
		Model *Model
	}

	testCases := []TestCase{
		{
			Name: "Valid model",
			Model: &Model{
				Items: []*Item{
					{IsCompleted: false, Title: "Item 1"},
					{IsCompleted: true, Title: "Item 2"},
				},
			},
		},
		{
			Name:  "Empty model",
			Model: &Model{Items: []*Item{}},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.Name, func(t *testing.T) {
			logutil.WithTestLogger(t, func() {
				f, err := os.CreateTemp("", "model_test_*.txt")
				require.NoError(t, err)
				defer func() { _ = os.Remove(f.Name()) }()

				path := f.Name()
				_ = f.Close()

				err = tc.Model.Store(path)
				require.NoError(t, err)

				model, err := Load(path)
				require.NoError(t, err)

				assert.Equal(t, tc.Model, model)
			})
		})
	}
}
