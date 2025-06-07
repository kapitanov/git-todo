package testutil

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func WithTempDir(t *testing.T, fn func(dir string)) {
	t.Helper()

	dir, err := os.MkdirTemp("", "git-todo-test-*")
	require.NoError(t, err)

	dir, err = filepath.EvalSymlinks(dir)
	require.NoError(t, err)

	t.Logf("Temporary directory created: %q", dir)

	defer func() { _ = os.RemoveAll(dir) }()

	require.NoError(t, os.Chdir(dir))

	fn(dir)
}

func GitInit(t *testing.T, dir string) {
	t.Helper()

	cmd := exec.Command("git", "init")
	cmd.Dir = dir
	require.NoError(t, cmd.Run())

	t.Logf("Initialized git repository in %q", dir)
}
