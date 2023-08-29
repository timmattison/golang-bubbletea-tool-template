package support

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"time"
)

var (
	Program *tea.Program

	HelpStyle             = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render
	BoldStyle             = lipgloss.NewStyle().Bold(true)
	EmptyNonFatalErrorMsg = NonFatalErrorMsg{}
	EmptyFatalErrorMsg    = FatalErrorMsg{}
)

type FatalErrorMsg struct {
	Err error
}

func (e FatalErrorMsg) Error() string {
	return e.Err.Error()
}

type NonFatalErrorMsg struct {
	Err error
}

func (e NonFatalErrorMsg) Error() string {
	return e.Err.Error()
}

type ClearErrorMsg struct{}

func ClearErrorAfter(t time.Duration) tea.Cmd {
	return tea.Tick(t, func(_ time.Time) tea.Msg {
		return ClearErrorMsg{}
	})
}

type NonFatalErrorCmd tea.Cmd

func ToNonFatalErrorCmd(err NonFatalErrorMsg) tea.Cmd {
	return func() tea.Msg {
		return err
	}
}

func ToCmd(msg tea.Msg) tea.Cmd {
	return func() tea.Msg {
		return msg
	}
}
