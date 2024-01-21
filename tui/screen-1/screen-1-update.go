package screen_1

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/timmattison/golang-bubbletea-tool-template/global"
)

func (m Screen1Model) Update(untypedMsg tea.Msg) (Screen1Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	//switch typedMsg := untypedMsg.(type) {
	//}

	m.WaitingSpinner, cmd = m.WaitingSpinner.Update(untypedMsg)
	cmds = append(cmds, cmd)

	if m.Err != nil {
		cmds = append(cmds, func() tea.Msg {
			return global.FatalErrorMsg{Err: m.Err}
		})
		cmds = append(cmds, tea.Quit)
	}

	return m, tea.Batch(cmds...)
}
