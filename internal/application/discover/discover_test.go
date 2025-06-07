package discover

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/kapitanov/git-todo/internal/git"
	"github.com/kapitanov/git-todo/internal/logutil"
	"github.com/kapitanov/git-todo/internal/testutil"
)

func TestTODOPath_OutsideOfRepository(t *testing.T) {
	logutil.WithTestLogger(t, func() {
		testutil.WithTempDir(t, func(dir string) {
			repositoryRoot, dataFilePath, err := TODOPath()
			t.Logf("TODOPath() -> (%q, %q %v)", repositoryRoot, dataFilePath, err)

			if assert.Error(t, err) {
				assert.ErrorIs(t, err, git.ErrNoGitRepository)
			}
		})
	})
}

func TestTODOPath_DefaultPath(t *testing.T) {
	logutil.WithTestLogger(t, func() {
		testutil.WithTempDir(t, func(dir string) {
			testutil.GitInit(t, dir)

			subDir := filepath.Join(dir, "subdir")
			require.NoError(t, os.MkdirAll(subDir, 0755))
			require.NoError(t, os.Chdir(subDir))

			repositoryRoot, dataFilePath, err := TODOPath()
			t.Logf("TODOPath() -> (%q, %q %v)", repositoryRoot, dataFilePath, err)

			require.NoError(t, err)
			assert.Equal(t, dir, repositoryRoot)
			assert.Equal(t, filepath.Join(dir, ".git", "TODO"), dataFilePath)
		})
	})
}

func TestTODOPath_ExistingPath(t *testing.T) {
	logutil.WithTestLogger(t, func() {
		testutil.WithTempDir(t, func(dir string) {
			testutil.GitInit(t, dir)

			subDir := filepath.Join(dir, "subdir")
			require.NoError(t, os.MkdirAll(subDir, 0755))
			require.NoError(t, os.Chdir(subDir))

			repositoryRoot, dataFilePath, err := TODOPath()
			t.Logf("TODOPath() -> (%q, %q %v)", repositoryRoot, dataFilePath, err)

			require.NoError(t, err)
			assert.Equal(t, dir, repositoryRoot)
			assert.Equal(t, filepath.Join(dir, ".git", "TODO"), dataFilePath)
		})
	})
}
