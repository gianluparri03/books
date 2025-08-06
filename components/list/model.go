package list

import (
	"books/components/navigator"

	"github.com/charmbracelet/bubbles/key"
	bList "github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

// Model is a bubbletea model that lets the user choose from a list
type Model struct {
	items []Item
	model bList.Model

	canGoBack bool
	canEnter  bool
}

// New returns a new Model
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

// KeyBindings returns the available additional key bindings
func (m Model) KeyBindings() (kb []key.Binding) {
	if m.canGoBack {
		kb = append(kb, key.NewBinding(
			key.WithKeys("b"),
			key.WithHelp("b", "Go back"),
		))
	}

	if m.canEnter {
		kb = append(kb, key.NewBinding(
			key.WithKeys("enter", "o"),
			key.WithHelp("enter/o", "Open"),
		))
	}

	return kb
}

// Init is used by bubbletea
func (m Model) Init() tea.Cmd {
	return tea.WindowSize()
}

// NUpdate is used by Navigator and bubbletea
func (m Model) NUpdate(msg tea.Msg) (tea.Model, tea.Cmd, navigator.Jump) {
	switch msg := msg.(type) {

	// Resizes the window when necessary
	case tea.WindowSizeMsg:
		m.model.SetSize(msg.Width, msg.Height)

	case tea.KeyMsg:
		switch key := msg.String(); key {

		// Quits the application when clicking 'ctrl+c'
		case "ctrl+c":
			return m, tea.Quit, navigator.Jump{}

		// Tries to exit when clicking 'b'
		case "b":
			if m.canGoBack {
				return m, nil, navigator.Jump{Prev: true}
			}

		// Tries to enter when clicking 'enter'
		case "enter":
			if m.canEnter {
				if item, ok := m.model.SelectedItem().(Item); ok {
					return m, nil, navigator.Jump{Next: true, NextArg: item.Id()}
				}
			}
		}
	}

	// Otherwise updates and returns the current one
	var cmd tea.Cmd
	m.model, cmd = m.model.Update(msg)
	return m, cmd, navigator.Jump{}
}

// Update is used by bubbletea
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	nm, c, _ := m.NUpdate(msg)
	return nm, c
}

// View is used by bubbletea
func (m Model) View() string {
	return m.model.View()
}
