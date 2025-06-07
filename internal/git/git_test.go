package git

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/kapitanov/git-todo/internal/logutil"
	"github.com/kapitanov/git-todo/internal/testutil"
)

// Tests for RepositoryRoot()

func TestRepositoryRoot_OutsideOfRepository(t *testing.T) {
	logutil.WithTestLogger(t, func() {
		testutil.WithTempDir(t, func(dir string) {
			actual, err := RepositoryRoot()

			t.Logf("RepositoryRoot() -> (%q, %v)", actual, err)
			if assert.Error(t, err) {
				assert.ErrorIs(t, err, ErrNoGitRepository)
			}
		})
	})
}

func TestRepositoryRoot_DefaultPath(t *testing.T) {
	logutil.WithTestLogger(t, func() {
		testutil.WithTempDir(t, func(dir string) {
			testutil.GitInit(t, dir)

			subDir := filepath.Join(dir, "subdir")
			require.NoError(t, os.MkdirAll(subDir, 0755))
			require.NoError(t, os.Chdir(subDir))

			expected := dir

			actual, err := RepositoryRoot()
			t.Logf("RepositoryRoot() -> (%q, %v)", actual, err)
			require.NoError(t, err)

			assert.Equal(t, expected, actual)
		})
	})
}

func TestRepositoryRoot_ExistingPath(t *testing.T) {
	logutil.WithTestLogger(t, func() {
		testutil.WithTempDir(t, func(dir string) {
			testutil.GitInit(t, dir)

			subDir := filepath.Join(dir, "subdir")
			require.NoError(t, os.MkdirAll(subDir, 0755))
			require.NoError(t, os.Chdir(subDir))

			expected := dir

			actual, err := RepositoryRoot()
			t.Logf("RepositoryRoot() -> (%q, %v)", actual, err)
			require.NoError(t, err)

			assert.Equal(t, expected, actual)
		})
	})
}

// Tests for CurrentBranch()

func TestCurrentBranch_OutsideOfRepository(t *testing.T) {
	logutil.WithTestLogger(t, func() {
		testutil.WithTempDir(t, func(dir string) {
			actual, err := CurrentBranch()
			t.Logf("CurrentBranch() -> (%q, %v)", actual, err)

			if assert.Error(t, err) {
				assert.ErrorIs(t, err, ErrNoGitRepository)
			}
		})
	})
}

func TestCurrentBranch_EmptyRepository(t *testing.T) {
	logutil.WithTestLogger(t, func() {
		testutil.WithTempDir(t, func(dir string) {
			testutil.GitInit(t, dir)

			actual, err := CurrentBranch()
			t.Logf("CurrentBranch() -> (%q, %v)", actual, err)

			if assert.Error(t, err) {
				assert.ErrorIs(t, err, ErrNoGitRepository)
			}
		})
	})
}

func TestCurrentBranch_NonEmptyRepository(t *testing.T) {
	logutil.WithTestLogger(t, func() {
		testutil.WithTempDir(t, func(dir string) {
			const branch = "master"
			gitInitialCommit(t, dir, branch)

			actual, err := CurrentBranch()
			t.Logf("CurrentBranch() -> (%q, %v)", actual, err)

			assert.NoError(t, err)
			assert.Equal(t, branch, actual)
		})
	})
}

func gitInitialCommit(t *testing.T, dir, branch string) {
	t.Helper()

	{
		cmd := exec.Command("git", "init", "--initial-branch="+branch)
		cmd.Dir = dir
		require.NoError(t, cmd.Run())

		t.Logf("Initialized git repository in %q (initial branch %q)", dir, branch)
	}

	{
		cmd := exec.Command("git", "commit", "--allow-empty", "-m", "Initial commit")
		cmd.Dir = dir
		require.NoError(t, cmd.Run())

		t.Logf("Created an initial commit in git in %q (branch %q)", dir, branch)
	}
}
