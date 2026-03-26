package main

import (
	"log"

	tea "charm.land/bubbletea/v2"
	"github.com/SiVy1/Killjoy/cli"
)

func main() {
	m := cli.NewModel()
	p := tea.NewProgram(m)

	_, err := p.Run()
	if err != nil {
		log.Fatal(err)
	}
}
