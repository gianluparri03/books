package navigator

import tea "github.com/charmbracelet/bubbletea"

// Model is a wrapper for a model that has (or can have) a previous and a next one
type Model struct {
	prev    func() tea.Model
	current tea.Model
	next    func(string) tea.Model
}

// New returns a new Model
func New(prev func() tea.Model, current tea.Model, next func(string) tea.Model) tea.Model {
	return Model{prev: prev, current: current, next: next}
}

// Init is used by bubbletea
func (m Model) Init() tea.Cmd {
	return m.current.Init()
}

// Update is used by bubbletea
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Executes the NUpdate of the current model
	nm, c, j := TryNUpdate(m.current, msg)

	if j.Prev && m.prev != nil {
		// Returns the previous model
		nm = m.prev()
		c = nm.Init()
		return nm, c
	} else if j.Next && m.next != nil {
		// Returns the next model
		nm = m.next(j.NextArg)
		c = nm.Init()
		return nm, c
	} else {
		// Returns the (updated) current model, wrapped in the navigator model
		m.current = nm
		return m, c
	}
}

// View is used by bubbletea
func (m Model) View() string {
	return m.current.View()
}
