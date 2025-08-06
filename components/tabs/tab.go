package tabs

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Tab is a piece of the model, with a title and a content
type Tab struct {
	Title string
	Model tea.Model
}
