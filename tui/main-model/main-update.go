package main_model

import (
	"errors"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/timmattison/golang-bubbletea-tool-template/global"
)

func (m MainModel) Update(untypedMessage tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	if m.SessionState == NotInitialized {
		m.SessionState = Screen1
	}

	switch m.SessionState {
	// Make sure we send the updates to the correct nested model depending on our state
	case Screen1:
		m.Screen1, cmd = m.Screen1.Update(untypedMessage)
		cmds = append(cmds, cmd)
	default:
		m.FatalError = global.FatalErrorMsg{Err: fmt.Errorf("invalid session state")}
	}

	switch typedMessage := untypedMessage.(type) {
	case global.FatalErrorMsg:
		m.FatalError = typedMessage

	case global.NonFatalErrorMsg:
		m.NonFatalError = typedMessage
		cmds = append(cmds, global.ClearErrorAfter(2))

	case tea.KeyMsg:
		if typedMessage.Type == tea.KeyCtrlC || typedMessage.Type == tea.KeyEsc {
			m.Quitting = true
		}

	case global.ClearErrorMsg:
		m.NonFatalError = global.EmptyNonFatalErrorMsg
	}

	if !errors.Is(m.FatalError, global.EmptyFatalErrorMsg) || m.Quitting {
		cmds = append(cmds, tea.Quit)
	}

	return m, tea.Batch(cmds...)
}
