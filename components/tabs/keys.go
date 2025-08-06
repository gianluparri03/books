package tabs

import (
	"github.com/charmbracelet/bubbles/key"
)

// keyMap contains the some bindings.
type keyMap struct {
	// Back returns to the previous model
	Back key.Binding

	// Prev opens the previous tab
	Prev key.Binding

	// Next opens the next tab
	Next key.Binding

	// Quit quits the application
	Quit key.Binding
}

// ShortHelp returns some of the key bindings.
func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Back, k.Prev, k.Next, k.Quit}
}

// FullHelp returns all the key bindings, divided in columns.
func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		[]key.Binding{k.Back, k.Quit},
		[]key.Binding{k.Prev, k.Next},
	}
}

// DefaultKeys are the keys used by the Model.
var DefaultKeys = keyMap{
	Back: key.NewBinding(
		key.WithKeys("b"),
		key.WithHelp("b", "Go back"),
	),

	Prev: key.NewBinding(
		key.WithKeys("left", "shift+tab"),
		key.WithHelp("←/shift+tab", "Prev tab"),
	),

	Next: key.NewBinding(
		key.WithKeys("right", "tab"),
		key.WithHelp("→/tab", "Next tab"),
	),

	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "Quit"),
	),
}
