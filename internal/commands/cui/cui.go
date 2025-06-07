package cui

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	ItemIndexStyle         = lipgloss.NewStyle().Faint(true)
	ItemCheckboxStyle      = lipgloss.NewStyle().Bold(true)
	ItemTextStyle          = lipgloss.NewStyle()
	ItemCompletedTextStyle = lipgloss.NewStyle().Strikethrough(true)
)

func Confirm(message string) (bool, error) {
	fmt.Printf("%s (y/n)? ", message)
	var response string
	_, err := fmt.Scanln(&response)
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

func runEditor(text string) (string, error) {
	return withTempFile(text, func(tmpfile string) error {
		editor := getSystemEditor()
		cmd := exec.Command(editor, tmpfile)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
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

	if err = fn(tmpfile.Name()); err != nil {
		return "", err
	}

	bs, err := os.ReadFile(tmpfile.Name())
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

func getSystemEditor() string {
	if os.Getenv("EDITOR") != "" {
		editors := strings.Fields(os.Getenv("EDITOR"))
		if len(editors) > 0 && editors[0] != "" {
			return editors[0]
		}
	}

	if runtime.GOOS == "windows" {
		return "notepad"
	}

	if runtime.GOOS == "darwin" {
		return "nano"
	}

	return "vi"
}
