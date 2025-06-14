package cui

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/mattn/go-isatty"
	"github.com/mattn/go-tty"
	"github.com/rs/zerolog/log"
)

var (
	ItemIndexStyle         = lipgloss.NewStyle().Faint(true)
	ItemCheckboxStyle      = lipgloss.NewStyle().Bold(true)
	ItemTextStyle          = lipgloss.NewStyle()
	ItemCompletedTextStyle = lipgloss.NewStyle().Strikethrough(true)
)

func Confirm(message string) (bool, error) {
	_, _ = fmt.Fprintf(os.Stdout, "%s (y/n)? ", message)

	term, err := tty.Open()
	if err != nil {
		return false, err
	}
	defer func() { _ = term.Close() }()

	response, err := term.ReadString()
	if err != nil {
		return false, fmt.Errorf("failed to read input: %w", err)
	}

	response = strings.TrimSpace(strings.ToLower(response))
	if response == "y" || response == "yes" {
		return true, nil
	}

	return false, nil
}

func Edit(text, description string) (string, error) {
	if !IsInteractive() {
		return "", errors.New("unable to run editor: not running in an interactive terminal")
	}

	inputText := text
	if description != "" {
		inputText += "\n\n"
		for _, line := range strings.Split(description, "\n") {
			inputText += fmt.Sprintf("# %s\n", line)
		}
	}

	outputText, err := runEditor(inputText)
	if err != nil {
		return "", fmt.Errorf("failed to run editor: %w", err)
	}

	var outputLines []string
	for _, line := range strings.Split(outputText, "\n") {
		line = strings.TrimSpace(line)
		if line != "" && !strings.HasPrefix(line, "#") {
			outputLines = append(outputLines, line)
		}
	}

	return strings.Join(outputLines, " "), nil
}

func IsInteractive() bool {
	return isatty.IsTerminal(os.Stdin.Fd())
}

func runEditor(text string) (string, error) {
	return withTempFile(text, func(tmpfile string) error {
		editor := getSystemEditor()
		cmd := editor(tmpfile)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		log.Debug().Str("cmd", fmt.Sprintf("%s %s", cmd.Path, strings.Join(cmd.Args, " "))).Msg("running editor")
		if err := cmd.Start(); err != nil {
			return err
		}

		if err := cmd.Wait(); err != nil {
			var exitErr *exec.ExitError
			if errors.As(err, &exitErr) {
				log.Warn().Int("exitcode", exitErr.ExitCode()).Msg("failed to run editor")
			}
			return err
		}

		return nil
	})
}

func withTempFile(content string, fn func(string) error) (string, error) {
	tmpfile, err := os.CreateTemp("", "git-todo-*.txt")
	if err != nil {
		return "", err
	}
	defer func() { _ = os.Remove(tmpfile.Name()) }()

	_, err = tmpfile.WriteString(content)
	if err != nil {
		return "", err
	}
	if err := tmpfile.Close(); err != nil {
		return "", err
	}

	log.Debug().Str("file", tmpfile.Name()).Msg("created temp file")
	if err = fn(tmpfile.Name()); err != nil {
		return "", err
	}

	bs, err := os.ReadFile(tmpfile.Name())
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

func getSystemEditor() func(string) *exec.Cmd {
	if os.Getenv("EDITOR") != "" {
		editors := strings.Fields(os.Getenv("EDITOR"))
		if len(editors) > 0 && editors[0] != "" {
			return prepareEditorCommand(editors...)
		}
	}

	if runtime.GOOS == "windows" {
		return prepareEditorCommand("notepad")
	}

	if runtime.GOOS == "darwin" {
		return prepareEditorCommand("nano")
	}

	return prepareEditorCommand("vi")
}

func prepareEditorCommand(args ...string) func(path string) *exec.Cmd {
	return func(path string) *exec.Cmd {
		program := args[0]
		var programArgs []string
		programArgs = append(programArgs, args[1:]...)
		programArgs = append(programArgs, path)

		cmd := exec.Command(program, programArgs...)
		return cmd
	}
}
