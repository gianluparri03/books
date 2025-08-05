package fields

import (
	"books/components/navigator"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

// LabelColor is used to higlight the labels
var LabelColor = lipgloss.Color("#EE6FF8")

// Field is one of the visualized fields
type Field struct {
	Label  string
	Value  string
	Inline bool
}

// FieldsModel is a collection of fields
type FieldsModel []Field

// NewFieldsModel returns a FieldsModel
func NewFieldsModel(fields ...Field) FieldsModel {
	return fields
}

// Init is used by bubbletea
func (fm FieldsModel) Init() tea.Cmd {
	return nil
}

// NUpdate is used by Navigator and bubbletea
func (fm FieldsModel) NUpdate(msg tea.Msg) (tea.Model, tea.Cmd, navigator.Jump) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch key := msg.String(); key {
		case "ctrl+c", "q":
			return fm, tea.Quit, navigator.Jump{}
		case "b":
			return fm, nil, navigator.Jump{Prev: true}
		}
	}

	return fm, nil, navigator.Jump{}
}

// Update is used by bubbletea
func (fm FieldsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m, c, _ := fm.NUpdate(msg)
	return m, c
}

// View is used by bubbletea
func (fm FieldsModel) View() string {
	sb := strings.Builder{}

	labelStyle := lipgloss.NewStyle().Foreground(LabelColor)

	for i, field := range fm {
		sb.WriteString(labelStyle.Render(field.Label))

		if field.Inline {
			sb.WriteString(labelStyle.Render(": "))
		} else {
			sb.WriteString("\n")
		}

		sb.WriteString(field.Value)
		sb.WriteString("\n")

		if i < len(fm)-1 {
			sb.WriteString("\n")
		}
	}

	return sb.String()
}
