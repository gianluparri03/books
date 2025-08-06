package tabs

import (
	"books/components/navigator"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

// Model is a bubbletea model composed of tabs
type Model struct {
	tabs   []Tab
	active int
	help   help.Model
}

// New returns a new model
func New(tabs []Tab) tea.Model {
	return Model{tabs: tabs, help: help.New()}
}

// Init is used by bubbletea
func (m Model) Init() tea.Cmd {
	return m.tabs[m.active].Model.Init()
}

// NUpdate is used by Navigator and bubbletea
func (m Model) NUpdate(msg tea.Msg) (tea.Model, tea.Cmd, navigator.Jump) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		// Returns to the previous page
		case key.Matches(msg, DefaultKeys.Back):
			return m, nil, navigator.Jump{Prev: true}

		// Quits the application
		case key.Matches(msg, DefaultKeys.Quit):
			return m, tea.Quit, navigator.Jump{}

		// Selects the previous tab
		case key.Matches(msg, DefaultKeys.Prev):
			if m.active > 0 {
				m.active--
				return m, m.tabs[m.active].Model.Init(), navigator.Jump{}
			}

		// Selects the next tab
		case key.Matches(msg, DefaultKeys.Next):
			if m.active < len(m.tabs)-1 {
				m.active++
				return m, m.tabs[m.active].Model.Init(), navigator.Jump{}
			}
		}
	}

	// Propagates the message to the inner model
	var cmd tea.Cmd
	var jump navigator.Jump
	if tab, ok := m.tabs[m.active].Model.(navigator.NModel); ok {
		m.tabs[m.active].Model, cmd, jump = tab.NUpdate(msg)
	} else {
		m.tabs[m.active].Model, cmd = m.tabs[m.active].Model.Update(msg)
	}
	return m, cmd, jump
}

// Update is used by bubbletea
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	nm, c, _ := m.NUpdate(msg)
	return nm, c
}

// View is used by bubbletea
func (m Model) View() string {
	var width int
	sb := strings.Builder{}

	m.renderHeader(&sb, &width)
	sb.WriteString("\n")
	m.renderContent(&sb, width)
	sb.WriteString("\n\n")
	sb.WriteString(m.help.View(DefaultKeys))

	padding := lipgloss.NewStyle().Padding(1, 2)
	return padding.Render(sb.String())
}

// renderHeader writes the tabs header into a stringbuilder
func (m Model) renderHeader(sb *strings.Builder, width *int) {
	var tabs []string

	for i, t := range m.tabs {
		isFirst := i == 0
		isLast := i == len(m.tabs)-1
		isActive := i == m.active

		style := lipgloss.NewStyle().
			Padding(0, 1).
			Border(getTabBorder(isFirst, isLast, isActive), true).
			BorderForeground(BorderColor)

		tabs = append(tabs, style.Render(t.Title))
	}

	header := lipgloss.JoinHorizontal(lipgloss.Top, tabs...)
	*width = lipgloss.Width(header)
	sb.WriteString(header)
}

// renderContent writes the tabs content into a stringbuilder
func (m Model) renderContent(sb *strings.Builder, width int) {
	style := lipgloss.NewStyle().
		Padding(1, 2, 0).
		Border(lipgloss.NormalBorder()).
		UnsetBorderTop().
		BorderForeground(BorderColor)

	model := m.tabs[m.active].Model
	width -= style.GetHorizontalFrameSize() - style.GetHorizontalPadding()
	content := style.Width(width).Render(model.View())
	sb.WriteString(content)
}
