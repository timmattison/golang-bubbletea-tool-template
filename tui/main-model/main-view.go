package main_model

import (
	"errors"
	"fmt"
	"github.com/muesli/reflow/indent"
	"github.com/timmattison/golang-bubbletea-tool-template/support"
	"strings"
)

func (m MainModel) View() string {
	var output strings.Builder

	if !errors.Is(m.FatalError, support.EmptyFatalErrorMsg) {
		output.WriteString(fmt.Sprintf("Fatal error: %v\n", m.FatalError.Err))
	}

	if !errors.Is(m.NonFatalError, support.EmptyNonFatalErrorMsg) {
		output.WriteString(fmt.Sprintf("Non-fatal error: %v\n", m.NonFatalError.Err))
	}

	switch m.SessionState {
	case NotInitialized:
		return ""

	case Screen1:
		return m.Screen1.View()

	default:
		panic(fmt.Sprintf("unhandled default case %v, this is a bug", m.SessionState))
	}

	return indent.String(output.String(), 2)
}
