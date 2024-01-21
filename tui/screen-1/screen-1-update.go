package screen_1

import (
	"log/slog"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/timmattison/golang-bubbletea-tool-template/global"
)

func (m *Screen1Model) Update(untypedMsg tea.Msg) (Screen1Model, tea.Cmd) {
	if !m.initialized {
		slog.Error("Screen 1 was not initialized with Init, cannot continue")
		return *m, tea.Quit
	}

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

	return *m, tea.Batch(cmds...)
}
