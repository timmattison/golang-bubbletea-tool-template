package main_model

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/timmattison/golang-bubbletea-tool-template/support"
	"github.com/timmattison/golang-bubbletea-tool-template/tui/screen-1"
	"log/slog"
)

type SessionState int

const (
	NotInitialized SessionState = iota
	Screen1
)

type MainModel struct {
	SessionState  SessionState
	Screen1       screen_1.Screen1Model
	Quitting      bool
	NonFatalError support.NonFatalErrorMsg
	FatalError    support.FatalErrorMsg
	initialized   bool
}

func New() MainModel {
	screen1 := screen_1.New()

	return MainModel{
		SessionState: NotInitialized,
		Screen1:      screen1,
		initialized:  true,
	}
}

func (m MainModel) Init() tea.Cmd {
	if !m.initialized {
		slog.Error("Main model was not initialized properly, cannot continue")
		return tea.Quit
	}

	return tea.Batch(
		m.Screen1.Init(),
	)
}