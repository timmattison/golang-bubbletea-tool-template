package main_model

import (
	"errors"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/timmattison/golang-bubbletea-tool-template/support"
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
		m.FatalError = support.FatalErrorMsg{Err: fmt.Errorf("invalid session state")}
	}

	switch typedMessage := untypedMessage.(type) {
	case support.FatalErrorMsg:
		m.FatalError = typedMessage

	case support.NonFatalErrorMsg:
		m.NonFatalError = typedMessage
		cmds = append(cmds, support.ClearErrorAfter(2))

	case tea.KeyMsg:
		if typedMessage.Type == tea.KeyCtrlC || typedMessage.Type == tea.KeyEsc {
			m.Quitting = true
		}

	case support.ClearErrorMsg:
		m.NonFatalError = support.EmptyNonFatalErrorMsg
	}

	if !errors.Is(m.FatalError, support.EmptyFatalErrorMsg) || m.Quitting {
		cmds = append(cmds, tea.Quit)
	}

	return m, tea.Batch(cmds...)
}
