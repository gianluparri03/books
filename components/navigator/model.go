package navigator

import tea "github.com/charmbracelet/bubbletea"

// Model is a bubbletea model that works as a wrapper for other models that can
// implements the NModel interface.
type Model struct {
	prev    func() tea.Model
	current tea.Model
	next    func(string) tea.Model
}

// New returns a new Model. Current is the current model, prev a function
// returning the previous model and next another function with a string
// parameter that returns the next model.
func New(prev func() tea.Model, current tea.Model, next func(string) tea.Model) tea.Model {
	return Model{prev: prev, current: current, next: next}
}

func (m Model) Init() tea.Cmd {
	return m.current.Init()
}

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

func (m Model) View() string {
	return m.current.View()
}
