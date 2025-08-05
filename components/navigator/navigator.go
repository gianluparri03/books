package navigator

import tea "github.com/charmbracelet/bubbletea"

// Jump is used by NModels to signal the navigator that it needs to jump
// to the previous or next model
type Jump struct {
	Prev    bool
	Next    bool
	NextArg string
}

// NModel is a bubbletea model that also implements NUpdate, which can return
// a jump alongside the model and the command
type NModel interface {
	tea.Model
	NUpdate(tea.Msg) (tea.Model, tea.Cmd, Jump)
}

// Navigator is a wrapper for a model that has (or can have) a previous and
// a next one
type Navigator struct {
	prev    func() tea.Model
	current tea.Model
	next    func(string) tea.Model
}

// NewNavigator returns a new NavigatorModel
func NewNavigator(prev func() tea.Model, current tea.Model, next func(string) tea.Model) tea.Model {
	return Navigator{prev, current, next}
}

// Init is used by bubbletea
func (n Navigator) Init() tea.Cmd {
	return n.current.Init()
}

// Update is used by bubbletea
func (n Navigator) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	var j Jump

	// Executes the NUpdate of the current model, if it exists
	if current, ok := n.current.(NModel); ok {
		m, c, j = current.NUpdate(msg)
	} else {
		m, c = n.current.Update(msg)
	}

	if j.Prev && n.prev != nil {
		// Returns the previous model
		m = n.prev()
		c = m.Init()
		return m, c
	} else if j.Next && n.next != nil {
		// Returns the next model
		m = n.next(j.NextArg)
		c = m.Init()
		return m, c
	} else {
		// Returns the (updated) current model, wrapped in the navigator
		n.current = m
		return n, c
	}
}

// View is used by bubbletea
func (n Navigator) View() string {
	return n.current.View()
}
