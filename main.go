package main

import (
	"log"

	tea "charm.land/bubbletea/v2"
	ui "github.com/SiVy1/Killjoy/internal/ui"
)

func main() {
	m := ui.NewModel()
	p := tea.NewProgram(m)

	_, err := p.Run()
	if err != nil {
		log.Fatal(err)
	}
}
