package screen_1

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"log/slog"
)

type Screen1Model struct {
	Message        string
	WaitingSpinner spinner.Model
	Err            error
	initialized    bool
}

func New() Screen1Model {
	waitingSpinner := spinner.New()
	waitingSpinner.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	waitingSpinner.Spinner = spinner.Points

	return Screen1Model{
		Message:        "This is screen 1",
		WaitingSpinner: waitingSpinner,
		initialized:    true,
	}
}

func (m Screen1Model) Init() tea.Cmd {
	if !m.initialized {
		slog.Error("Screen 1 was not initialized properly, cannot continue")
		return tea.Quit
	}

	return m.WaitingSpinner.Tick
}
