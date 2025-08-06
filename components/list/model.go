package list

import (
	"books/components/navigator"

	"github.com/charmbracelet/bubbles/key"
	bList "github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

// Model is a bubbletea model that lets the user choose from a bubbles list.
// If wrapped inside a navigator.Model, when an item will be chosen, the NUpdate
// method will return a next jump containing the id of the chosen item.
type Model struct {
	items []Item
	model bList.Model

	canGoBack bool
	canEnter  bool
}

// New returns a new Model. canGoBack and canEnter indicates whether the two
// jumps are enabled
func New(title string, items []Item, canGoBack, canEnter bool) tea.Model {
	m := Model{canGoBack: canGoBack, canEnter: canEnter}

	var bItems []bList.Item
	for _, i := range items {
		bItems = append(bItems, i)
	}

	m.model = bList.New(bItems, bList.NewDefaultDelegate(), 0, 0)
	m.model.Title = title
	m.model.AdditionalShortHelpKeys = m.KeyBindings
	m.model.AdditionalFullHelpKeys = m.KeyBindings
	return m
}

// KeyBindings returns the available additional key bindings.
func (m Model) KeyBindings() (kb []key.Binding) {
	if m.canGoBack {
		kb = append(kb, key.NewBinding(
			key.WithKeys("b"),
			key.WithHelp("b", "Go back"),
		))
	}

	if m.canEnter {
		kb = append(kb, key.NewBinding(
			key.WithKeys("o", "enter"),
			key.WithHelp("o/enter", "Open"),
		))
	}

	return kb
}

func (m Model) Init() tea.Cmd {
	return tea.WindowSize()
}

func (m Model) NUpdate(msg tea.Msg) (tea.Model, tea.Cmd, navigator.Jump) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg: // Resizes the window when necessary
		m.model.SetSize(msg.Width, msg.Height)

	case tea.KeyMsg:
		if m.model.FilterState() != bList.Filtering {
			switch key := msg.String(); key {

			// Quits the application
			case "q", "ctrl+c":
				return m, tea.Quit, navigator.Jump{}

			// Jumps to the previous model
			case "b":
				if m.canGoBack {
					return m, nil, navigator.Jump{Prev: true}
				}

			// Jumps to the next model
			case "enter":
				if m.canEnter {
					if item, ok := m.model.SelectedItem().(Item); ok {
						return m, nil, navigator.Jump{Next: true, NextArg: item.Id()}
					}
				}
			}
		}
	}

	// Updates the inner model
	var cmd tea.Cmd
	m.model, cmd = m.model.Update(msg)
	return m, cmd, navigator.Jump{}
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	nm, c, _ := m.NUpdate(msg)
	return nm, c
}

func (m Model) View() string {
	return m.model.View()
}
