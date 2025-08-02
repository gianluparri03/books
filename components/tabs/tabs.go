package tabs

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

// BorderColor is the color of the tabs border
var BorderColor = lipgloss.Color("#874BFD")

// TabbedModel is a bubbletea model composed of tabs
type TabbedModel struct {
	tabs   []Tab
	active int
}

// Tab is a piece of the model, with a title and a content
type Tab struct {
	Title string
	Model tea.Model
}

// NewTabbedModel returns a new model made of the given tabs
func NewTabbedModel(tabs []Tab) tea.Model {
	return TabbedModel{tabs: tabs}
}

// Init is used by bubbletea
func (tm TabbedModel) Init() tea.Cmd {
	return tm.tabs[tm.active].Model.Init()
}

// Update is used by bubbletea
func (tm TabbedModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch key := msg.String(); key {
		case "ctrl+c", "q":
			return tm, tea.Quit

		// Selects the previous tab
		case "left", "shift+tab":
			if tm.active > 0 {
				tm.active--
				return tm, tm.tabs[tm.active].Model.Init()
			}

		// Selects the next tab
		case "right", "tab":
			if tm.active < len(tm.tabs)-1 {
				tm.active++
				return tm, tm.tabs[tm.active].Model.Init()
			}
		}
	}

	// Propagates the message to the inner model
	var cmd tea.Cmd
	tm.tabs[tm.active].Model, cmd = tm.tabs[tm.active].Model.Update(msg)
	return tm, cmd
}

// View is used by bubbletea
func (tm TabbedModel) View() string {
	var width int
	sb := strings.Builder{}

	tm.renderHeader(&sb, &width)
	sb.WriteString("\n")
	tm.renderContent(&sb, width)

	padding := lipgloss.NewStyle().Padding(1, 2)
	return padding.Render(sb.String())
}

// renderHeader writes the tabs header into a stringbuilder
func (tm TabbedModel) renderHeader(sb *strings.Builder, width *int) {
	var tabs []string

	for i, t := range tm.tabs {
		isFirst := i == 0
		isLast := i == len(tm.tabs)-1
		isActive := i == tm.active

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
func (tm TabbedModel) renderContent(sb *strings.Builder, width int) {
	style := lipgloss.NewStyle().
		Padding(1, 2, 0).
		Border(lipgloss.NormalBorder()).
		UnsetBorderTop().
		BorderForeground(BorderColor)

	model := tm.tabs[tm.active].Model
	width -= style.GetHorizontalFrameSize() - style.GetHorizontalPadding()
	content := style.Width(width).Render(model.View())
	sb.WriteString(content)
}

// getTabBorder returns the border around a tab's title
func getTabBorder(isFirst, isLast, isActive bool) lipgloss.Border {
	border := lipgloss.RoundedBorder()

	if isFirst && isActive {
		border.BottomLeft = "│"
		border.Bottom = " "
		border.BottomRight = "└"
	} else if isLast && isActive {
		border.BottomLeft = "┘"
		border.Bottom = " "
		border.BottomRight = "│"
	} else if isActive {
		border.BottomLeft = "┘"
		border.Bottom = " "
		border.BottomRight = "└"
	} else if isFirst {
		border.BottomLeft = "├"
		border.Bottom = "─"
		border.BottomRight = "┴"
	} else if isLast {
		border.BottomLeft = "┴"
		border.Bottom = "─"
		border.BottomRight = "┤"
	} else {
		border.BottomLeft = "┴"
		border.Bottom = "─"
		border.BottomRight = "┴"
	}

	return border
}
