package installer

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/kapitanov/git-todo/internal/logutil"
	"github.com/kapitanov/git-todo/internal/testutil"
)

func testHook() hookDefinition {
	return hookDefinition{
		Name:    "post-commit",
		Command: "# git-todo post-commit hook\ngit-todo githooks post-commit\n",
	}
}

// Tests for installing git hooks

func TestInstallHook_NotInstalled(t *testing.T) {
	logutil.WithTestLogger(t, func() {
		testutil.WithTempDir(t, func(dir string) {
			testutil.GitInit(t, dir)

			h := testHook()
			err := installHook(dir, h, false)
			require.NoError(t, err)

			filePath := filepath.Join(dir, ".git", "hooks", h.Name)

			verifyGitHook(t, filePath, func(v *verifier) {
				v.Contains(h.Command)
			})
		})
	})
}

func TestInstallHook_NotInstalledButHookExists(t *testing.T) {
	logutil.WithTestLogger(t, func() {
		testutil.WithTempDir(t, func(dir string) {
			testutil.GitInit(t, dir)

			h := testHook()
			filePath := filepath.Join(dir, ".git", "hooks", h.Name)

			require.NoError(t, os.MkdirAll(filepath.Dir(filePath), 0755))
			originalCommand := "echo 'dummy hook'\n"
			originalContent := "#!/bin/sh\n\necho 'dummy hook'\n"
			require.NoError(t, os.WriteFile(filePath, []byte(originalContent), 0644))

			err := installHook(dir, h, false)
			require.NoError(t, err)

			verifyGitHook(t, filePath, func(v *verifier) {
				v.Contains(originalCommand)
				v.Contains(h.Command)
			})
			assert.FileExists(t, filePath+".bak")
		})
	})
}

func TestInstallHook_AlreadyInstalled(t *testing.T) {
	logutil.WithTestLogger(t, func() {
		testutil.WithTempDir(t, func(dir string) {
			testutil.GitInit(t, dir)

			h := testHook()
			filePath := filepath.Join(dir, ".git", "hooks", h.Name)

			require.NoError(t, os.MkdirAll(filepath.Dir(filePath), 0755))
			originalCommand := "echo 'dummy hook'\n"
			originalContent := fmt.Sprintf("#!/bin/sh\n\n%s\n%s\n", originalCommand, h.Command)

			require.NoError(t, os.WriteFile(filePath, []byte(originalContent), 0644)) // NOTE: not executable!

			err := installHook(dir, h, false)
			require.NoError(t, err)

			verifyGitHook(t, filePath, func(v *verifier) {
				v.Contains(originalCommand)
				v.Contains(h.Command)
			})
			assert.NoFileExists(t, filePath+".bak")
		})
	})
}

func TestInstallHook_AlreadyInstalledForce(t *testing.T) {
	logutil.WithTestLogger(t, func() {
		testutil.WithTempDir(t, func(dir string) {
			testutil.GitInit(t, dir)

			h := testHook()
			filePath := filepath.Join(dir, ".git", "hooks", h.Name)

			require.NoError(t, os.MkdirAll(filepath.Dir(filePath), 0755))
			originalCommand := "echo 'dummy hook'\n"
			originalContent := fmt.Sprintf("#!/bin/sh\n\n%s\n%s\n", originalCommand, h.Command)

			require.NoError(t, os.WriteFile(filePath, []byte(originalContent), 0644))

			err := installHook(dir, h, true)
			require.NoError(t, err)

			verifyGitHook(t, filePath, func(v *verifier) {
				v.NotContains(originalCommand)
				v.Contains(h.Command)
			})
			assert.FileExists(t, filePath+".bak")
		})
	})
}

// Tests for uninstalling git hooks

func TestUninstallHook_NotInstalled(t *testing.T) {
	logutil.WithTestLogger(t, func() {
		testutil.WithTempDir(t, func(dir string) {
			testutil.GitInit(t, dir)

			h := testHook()
			err := uninstallHook(dir, h)
			require.NoError(t, err)

			filePath := filepath.Join(dir, ".git", "hooks", h.Name)

			assert.NoFileExists(t, filePath)
		})
	})
}

func TestUninstallHook_NotInstalledButHookExists(t *testing.T) {
	logutil.WithTestLogger(t, func() {
		testutil.WithTempDir(t, func(dir string) {
			testutil.GitInit(t, dir)

			h := testHook()
			filePath := filepath.Join(dir, ".git", "hooks", h.Name)

			require.NoError(t, os.MkdirAll(filepath.Dir(filePath), 0755))
			originalCommand := "echo 'dummy hook'"
			originalContent := "#!/bin/sh\n\necho 'dummy hook'\n"
			require.NoError(t, os.WriteFile(filePath, []byte(originalContent), 0755))

			err := uninstallHook(dir, h)
			require.NoError(t, err)

			verifyGitHook(t, filePath, func(v *verifier) {
				v.Contains(originalCommand)
				v.NotContains(h.Command)
			})
			assert.NoFileExists(t, filePath+".bak")
		})
	})
}

func TestUninstallHook_AlreadyInstalled(t *testing.T) {
	logutil.WithTestLogger(t, func() {
		testutil.WithTempDir(t, func(dir string) {
			testutil.GitInit(t, dir)

			h := testHook()
			filePath := filepath.Join(dir, ".git", "hooks", h.Name)

			require.NoError(t, os.MkdirAll(filepath.Dir(filePath), 0755))
			originalCommand := "echo 'dummy hook'"
			originalContent := fmt.Sprintf("#!/bin/sh\n\n%s\n%s\n", originalCommand, h.Command)

			require.NoError(t, os.WriteFile(filePath, []byte(originalContent), 0644)) // NOTE: not executable!

			err := uninstallHook(dir, h)
			require.NoError(t, err)

			verifyGitHook(t, filePath, func(v *verifier) {
				v.Contains(originalCommand)
				v.NotContains(h.Command)
			})
			assert.FileExists(t, filePath+".bak")
		})
	})
}

// Test helpers

type verifier struct {
	t       *testing.T
	path    string
	content string
}

func (v *verifier) Contains(command string) {
	assert.Containsf(v.t, v.content, command, "git hook must contain command %q", command)
}

func (v *verifier) NotContains(command string) {
	assert.NotContainsf(v.t, v.content, command, "git hook must not contain command %q", command)
}

func (v *verifier) HasShebang() {
	assert.True(v.t, strings.HasPrefix(v.content, "#!/bin/sh\n"), "git hook must start with shebang")
}

func (v *verifier) IsExecutable() {
	info, err := os.Stat(v.path)
	require.NoError(v.t, err)

	perm := info.Mode().Perm()
	assert.Truef(v.t, perm&0111 != 0, "git hook must be executable but got permissions: %o", perm)
}

func verifyGitHook(t *testing.T, filePath string, f func(*verifier)) {
	if !assert.FileExists(t, filePath) {
		return
	}

	bs, err := os.ReadFile(filePath)
	require.NoError(t, err)

	v := &verifier{t: t, path: filePath, content: string(bs)}
	v.HasShebang()
	v.IsExecutable()
	f(v)
}
