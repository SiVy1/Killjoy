package cli

import (
	"strings"

	tea "charm.land/bubbletea/v2"
)

type model struct {
	tracks []string
	cursor int
}

func NewModel() model {
	return model{
		tracks: []string{
			"The five of us are dying",
			"My way home is through you",
			"Planetary (GO!)",
		},
		cursor: 0,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down":
			if m.cursor < len(m.tracks)-1 {
				m.cursor++
			}
		}
	}
	return m, nil
}

func (m model) View() tea.View {
	var b strings.Builder

	b.WriteString("Killjoy")
	b.WriteString("Library:\n")

	for i, t := range m.tracks {
		prefix := "  "
		if i == m.cursor {
			prefix = "> "
		}
		b.WriteString(prefix + t + "\n")
	}
	b.WriteString("\nq quit | up/down move\n")
	return tea.NewView(b.String())
}
