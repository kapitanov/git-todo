package git

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"github.com/rs/zerolog/log"
)

var (
	ErrNoGitRepository = errors.New("no git repository found")
)

func RepositoryRoot() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	log.Debug().Str("cmd", fmt.Sprintf("%s %s", cmd.Path, strings.Join(cmd.Args, " "))).Msg("executing command")

	output, err := cmd.Output()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			log.Warn().Str("stderr", string(exitError.Stderr)).Int("exitcode", exitError.ExitCode()).Msg("command exited with error")

			const noRepositoryExitCode = 128
			if exitError.ExitCode() == noRepositoryExitCode {
				return "", ErrNoGitRepository
			}

			return "", fmt.Errorf("commmand '%s %s' exited with exit code %d", cmd.Path, strings.Join(cmd.Args, " "), exitError.ExitCode())
		}

		log.Error().Err(err).Msg("command failed to execute")
		return "", fmt.Errorf("commmand '%s %s' failed: %w", cmd.Path, strings.Join(cmd.Args, " "), err)
	}

	rootDir := string(output)
	rootDir = strings.TrimSpace(rootDir)

	log.Info().Str("root", rootDir).Msg("discovered git repository root")
	return rootDir, nil
}

const (
	Master = "master"
	Main   = "main"
)

func CurrentBranch() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	log.Debug().Str("cmd", fmt.Sprintf("%s %s", cmd.Path, strings.Join(cmd.Args, " "))).Msg("executing command")

	output, err := cmd.Output()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			log.Warn().Str("stderr", string(exitError.Stderr)).Int("exitcode", exitError.ExitCode()).Msg("command exited with error")

			const noRepositoryExitCode = 128
			if exitError.ExitCode() == noRepositoryExitCode {
				return "", ErrNoGitRepository
			}

			return "", fmt.Errorf("commmand '%s %s' exited with exit code %d", cmd.Path, strings.Join(cmd.Args, " "), exitError.ExitCode())
		}

		log.Error().Err(err).Msg("command failed to execute")
		return "", fmt.Errorf("commmand '%s %s' failed: %w", cmd.Path, strings.Join(cmd.Args, " "), err)
	}

	branchName := string(output)
	branchName = strings.TrimSpace(branchName)

	log.Info().Str("branch", branchName).Msg("discovered current git branch")
	return branchName, nil
}
