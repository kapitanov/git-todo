package installer

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	preCommitHook = hookDefinition{
		Name:    "pre-commit",
		Command: "# git-todo: \"pre-commit\" git hook\ngit-todo githooks pre-commit",
	}

	prePushHook = hookDefinition{
		Name:    "pre-push",
		Command: "# git-todo: \"pre-push\" git hook\ngit-todo githooks pre-push",
	}

	allHooks = []hookDefinition{preCommitHook, prePushHook}
)

func Install(rootDir string, force bool) error {
	for _, hook := range allHooks {
		if err := installHook(rootDir, hook, force); err != nil {
			return err
		}
	}

	return nil
}

func Uninstall(rootDir string) error {
	for _, hook := range allHooks {
		if err := uninstallHook(rootDir, hook); err != nil {
			return err
		}
	}

	return nil
}

const (
	dirPermissions  = 0755
	filePermissions = 0755
)

type hookDefinition struct {
	Name    string
	Command string
}

func installHook(rootDir string, h hookDefinition, force bool) error {
	i := newHookInstaller(h)
	return i.Install(rootDir, force)
}

func uninstallHook(rootDir string, h hookDefinition) error {
	i := newHookInstaller(h)
	return i.Uninstall(rootDir)
}

type hookInstaller struct {
	name    string
	command string
	log     zerolog.Logger
}

func newHookInstaller(d hookDefinition) hookInstaller {
	return hookInstaller{
		name:    d.Name,
		command: strings.TrimSpace(d.Command),
		log:     log.Logger.With().Str("hook", d.Name).Logger(),
	}
}

func (i *hookInstaller) Install(rootDir string, force bool) error {
	hookFilePath := filepath.Join(rootDir, ".git", "hooks", i.name)
	i.log.Info().Str("file", hookFilePath).Msg("installing git hook")

	hookDir := filepath.Dir(hookFilePath)
	if err := os.MkdirAll(hookDir, dirPermissions); err != nil {
		i.log.Error().Err(err).Str("dir", hookDir).Msg("failed to create git hook directory")
		return fmt.Errorf("failed to create hook directory %q: %w", hookDir, err)
	}

	if err := i.installHookCore(hookFilePath, force); err != nil {
		return err
	}

	if err := i.ensureHookExecutable(hookFilePath); err != nil {
		return err
	}

	return nil
}

func (i *hookInstaller) Uninstall(rootDir string) error {
	hookFilePath := filepath.Join(rootDir, ".git", "hooks", i.name)
	i.log.Info().Str("file", hookFilePath).Msg("uninstalling git hook")

	if _, err := os.Stat(hookFilePath); os.IsNotExist(err) {
		i.log.Info().Msg("git hook is not installed")
		return nil
	}

	if err := i.uninstallHookFromExisting(hookFilePath); err != nil {
		return err
	}

	return nil
}

func (i *hookInstaller) installHookCore(hookFilePath string, force bool) error {
	if _, err := os.Stat(hookFilePath); os.IsNotExist(err) {
		return i.installHookFromScratch(hookFilePath)
	}

	if force {
		return i.reinstallHookFromScratch(hookFilePath)
	}

	return i.installHookIntoExisting(hookFilePath)
}

func (i *hookInstaller) installHookFromScratch(hookFilePath string) error {
	content := fmt.Sprintf("#!/bin/sh\n\n%s\n", i.command)

	if err := os.WriteFile(hookFilePath, []byte(content), filePermissions); err != nil {
		i.log.Error().Err(err).Msg("failed to install git hook")
		return fmt.Errorf("failed to install git hook %q: %w", hookFilePath, err)
	}

	i.log.Info().Msg("installed git hook")
	return nil
}

func (i *hookInstaller) reinstallHookFromScratch(hookFilePath string) error {
	if err := i.backupHook(hookFilePath); err != nil {
		return err
	}

	return i.installHookFromScratch(hookFilePath)
}

func (i *hookInstaller) installHookIntoExisting(hookFilePath string) error {
	// Read the existing hook content
	content, err := os.ReadFile(hookFilePath)
	if err != nil {
		i.log.Error().Err(err).Msg("failed to read existing git hook")
		return fmt.Errorf("failed to read existing git hook %q: %w", hookFilePath, err)
	}

	if strings.Contains(string(content), i.command) {
		i.log.Info().Msg("git hook is already installed")
		return nil
	}

	// Create a backup of the existing hook file
	if err := i.backupHook(hookFilePath); err != nil {
		return err
	}

	// Append the new command to the existing content
	newContent := fmt.Sprintf("%s\n%s\n", strings.TrimSpace(string(content)), i.command)
	if err := os.WriteFile(hookFilePath, []byte(newContent), filePermissions); err != nil {
		i.log.Error().Err(err).Msg("failed to update git hook")
		return fmt.Errorf("failed to update git hook %q: %w", hookFilePath, err)
	}

	i.log.Info().Msg("updated git hook")
	return nil
}

func (i *hookInstaller) uninstallHookFromExisting(hookFilePath string) error {
	// Read the existing hook content
	content, err := os.ReadFile(hookFilePath)
	if err != nil {
		i.log.Error().Err(err).Msg("failed to read existing git hook")
		return fmt.Errorf("failed to read existing git hook %q: %w", hookFilePath, err)
	}

	if !strings.Contains(string(content), i.command) {
		i.log.Info().Msg("git hook is not installed")
		return nil
	}

	// Create a backup of the existing hook file
	if err := i.backupHook(hookFilePath); err != nil {
		return err
	}

	// Append the new command to the existing content
	newContent := strings.ReplaceAll(string(content), i.command, "")
	newContent = strings.TrimSpace(newContent)

	if err := os.WriteFile(hookFilePath, []byte(newContent), filePermissions); err != nil {
		i.log.Error().Err(err).Msg("failed to update git hook")
		return fmt.Errorf("failed to update git hook %q: %w", hookFilePath, err)
	}

	i.log.Info().Msg("updated git hook")

	if err := i.ensureHookExecutable(hookFilePath); err != nil {
		return err
	}

	return nil
}

func (i *hookInstaller) backupHook(hookFilePath string) error {
	content, err := os.ReadFile(hookFilePath)
	if err != nil {
		i.log.Error().Err(err).Msg("failed to read existing git hook")
		return fmt.Errorf("failed to read existing git hook %q: %w", hookFilePath, err)
	}

	backupFilePath := hookFilePath + ".bak"
	if err := os.WriteFile(backupFilePath, content, filePermissions); err != nil {
		i.log.Error().Err(err).Str("file", hookFilePath).Msg("failed to backup git hook")
		return fmt.Errorf("failed to backup git hook %q: %w", hookFilePath, err)
	}

	i.log.Info().Str("file", backupFilePath).Msg("created a backup of a git hook")
	return nil
}

func (i *hookInstaller) ensureHookExecutable(hookFilePath string) error {
	i.log.Debug().Str("file", hookFilePath).Int("perm", filePermissions).Msg("set git hook as executable")

	if err := os.Chmod(hookFilePath, filePermissions); err != nil {
		i.log.Error().Err(err).Msg("failed to set permissions for git hook")
		return fmt.Errorf("failed to set permissions for git hook %q: %w", hookFilePath, err)
	}

	return nil
}
