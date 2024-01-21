package screen_1

import (
	"log/slog"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Screen1Model struct {
	Message        string
	WaitingSpinner spinner.Model
	Err            error
	newed          bool
	initialized    bool
}

func New() Screen1Model {
	waitingSpinner := spinner.New()
	waitingSpinner.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	waitingSpinner.Spinner = spinner.Points

	return Screen1Model{
		Message:        "This is screen 1",
		WaitingSpinner: waitingSpinner,
		newed:          true,
	}
}

func (m *Screen1Model) Init() tea.Cmd {
	if !m.newed {
		slog.Error("Screen 1 was not created with New, cannot continue")
		return tea.Quit
	}

	m.initialized = true

	return m.WaitingSpinner.Tick
}
