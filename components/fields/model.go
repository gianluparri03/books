package fields

import (
	"books/components/navigator"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

// Model is a collection of fields
type Model []Field

// New returns a new Model
func New(fields ...Field) Model {
	return fields
}

// Init is used by bubbletea
func (m Model) Init() tea.Cmd {
	return nil
}

// NUpdate is used by Navigator and bubbletea
func (m Model) NUpdate(msg tea.Msg) (tea.Model, tea.Cmd, navigator.Jump) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch key := msg.String(); key {
		case "ctrl+c", "q":
			return m, tea.Quit, navigator.Jump{}
		case "b":
			return m, nil, navigator.Jump{Prev: true}
		}
	}

	return m, nil, navigator.Jump{}
}

// Update is used by bubbletea
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	nm, c, _ := m.NUpdate(msg)
	return nm, c
}

// View is used by bubbletea
func (m Model) View() string {
	sb := strings.Builder{}

	labelStyle := lipgloss.NewStyle().Foreground(LabelColor)

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
