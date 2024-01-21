package main_model

import (
	"errors"
	"fmt"
	"strings"

	"github.com/muesli/reflow/indent"
	"github.com/timmattison/golang-bubbletea-tool-template/global"
)

func (m MainModel) View() string {
	var output strings.Builder

	if !errors.Is(m.FatalError, global.EmptyFatalErrorMsg) {
		output.WriteString(fmt.Sprintf("Fatal error: %v\n", m.FatalError.Err))
	}

	if !errors.Is(m.NonFatalError, global.EmptyNonFatalErrorMsg) {
		output.WriteString(fmt.Sprintf("Non-fatal error: %v\n", m.NonFatalError.Err))
	}

	switch m.SessionState {
	case NotInitialized:
		return ""

	case Screen1:
		output.WriteString(m.Screen1.View())

	default:
		panic(fmt.Sprintf("unhandled default case %v, this is a bug", m.SessionState))
	}
	return indent.String(output.String(), 2)
}
