package screen_1

import (
	"github.com/muesli/reflow/indent"
	"strings"
)

func (m Screen1Model) View() string {
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

	return indent.String(output.String(), 2)
}
