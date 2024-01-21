package main

import (
	"fmt"
	"os"
	"runtime/debug"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/timmattison/golang-bubbletea-tool-template/global"
	"github.com/timmattison/golang-bubbletea-tool-template/tui/main-model"
)

var commit = func() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			if setting.Key == "vcs.revision" {
				return fmt.Sprintf("commit - %s", setting.Value)
			}
		}
	}
	return "NO COMMIT HASH FOUND - DEBUG VERSION ONLY - NOT FOR PRODUCTION USE"
}()

func main() {
	fmt.Println(commit)

	myModel := main_model.New()
	// myModel := main_model.MainModel{}

	global.Program = tea.NewProgram(&myModel)

	if _, err := global.Program.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
