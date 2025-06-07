package discover

import (
	"path/filepath"

	"github.com/kapitanov/git-todo/internal/git"
)

func TODOPath() (repositoryRootDir, dataFilePath string, err error) {
	repositoryRootDir, err = git.RepositoryRoot()
	if err != nil {
		return
	}

	dataFilePath = filepath.Join(repositoryRootDir, ".git", "TODO")
	return
}
