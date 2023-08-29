package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/timmattison/golang-bubbletea-tool-template/support"
	main_model "github.com/timmattison/golang-bubbletea-tool-template/tui/main-model"
	"os"
)

func main() {
	myModel := main_model.New()
	//myModel := main_model.MainModel{}

	support.Program = tea.NewProgram(&myModel)

	if _, err := support.Program.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
