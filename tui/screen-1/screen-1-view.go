package screen_1

import (
	"log/slog"
	"strings"
)

func (m *Screen1Model) View() string {
	if !m.initialized {
		slog.Error("Screen 1 was not initialized with Init, cannot continue")
		return ""
	}

	if m.Err != nil {
		return m.Err.Error()
	}

	var output strings.Builder

	output.WriteString("\n")

	if m.Err != nil {
		output.WriteString(m.Err.Error())
	} else {
		output.WriteString(m.WaitingSpinner.View())
		output.WriteString(" ")
		output.WriteString(m.Message)
		output.WriteString("\n")
	}

	return output.String()
}
