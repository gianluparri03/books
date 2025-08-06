package tabs

import (
	"github.com/charmbracelet/bubbles/key"
)

// keyMap contains the key bindings
type keyMap struct {
	Back key.Binding
	Next key.Binding
	Prev key.Binding
	Quit key.Binding
}

// ShortHelp returns some of the key bindings
func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Back, k.Prev, k.Next, k.Quit}
}

// FullHelp returns all the key bindings
func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		[]key.Binding{k.Back, k.Quit},
		[]key.Binding{k.Prev, k.Next},
	}
}

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
