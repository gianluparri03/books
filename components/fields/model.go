package fields

import (
	"books/components/navigator"

	tea "github.com/charmbracelet/bubbletea"
	"strings"
)

// Model is a bubbletea model that shows a list of Fields with their labels
// and values. This model is not interactive.
type Model []Field

// New returns a new Model.
func New(fields ...Field) Model {
	return fields
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) NUpdate(msg tea.Msg) (tea.Model, tea.Cmd, navigator.Jump) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch key := msg.String(); key {
		case "ctrl+c", "q": // quits the application
			return m, tea.Quit, navigator.Jump{}

		case "b": // jumps back to the previous model
			return m, nil, navigator.Jump{Prev: true}
		}
	}

	return m, nil, navigator.Jump{}
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	nm, c, _ := m.NUpdate(msg)
	return nm, c
}

func (m Model) View() string {
	sb := strings.Builder{}

	for i, field := range m {
		sb.WriteString(labelStyle.Render(field.Label))

		if field.Inline {
			sb.WriteString(labelStyle.Render(": "))
		} else {
			sb.WriteString("\n")
		}

		sb.WriteString(field.Value)
		sb.WriteString("\n")

		if i < len(m)-1 {
			sb.WriteString("\n")
		}
	}

	return sb.String()
}
