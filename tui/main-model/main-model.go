package main_model

import (
	"log/slog"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/timmattison/golang-bubbletea-tool-template/global"
	"github.com/timmattison/golang-bubbletea-tool-template/tui/screen-1"
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
	NonFatalError global.NonFatalErrorMsg
	FatalError    global.FatalErrorMsg
	// newed       indicates if the model was created with New
	newed bool
	// initialized indicates if the model's Init function has been called
	initialized bool
}

func New() MainModel {
	screen1 := screen_1.New()

	return MainModel{
		SessionState: NotInitialized,
		Screen1:      screen1,
		newed:        true,
	}
}

func (m *MainModel) Init() tea.Cmd {
	if !m.newed {
		slog.Error("Main model was not created with New, cannot continue")
		return tea.Quit
	}

	m.initialized = true

	return tea.Batch(
		m.Screen1.Init(),
	)
}
