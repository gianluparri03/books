package tabs

import (
	"github.com/charmbracelet/lipgloss"
)

// BorderColor is the color of the tabs border.
var BorderColor = lipgloss.Color("#874BFD")

// getTabBorder returns the border around a tab's header, based on the isFirst,
// isActive and isLast flags.
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
