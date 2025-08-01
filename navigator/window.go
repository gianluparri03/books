package navigator

import (
	"github.com/charmbracelet/bubbles/key"
	bubblesList "github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

// window is a window that lets the user choose from a list.
// It's made of a Navigator and a bubbles list
type window struct {
	nav   Navigator
	model bubblesList.Model
}

// newWindow returns a new window with the given navigator, title and items
func newWindow(nav Navigator, title string, items []Item) (w window) {
	var bItems []bubblesList.Item
	for _, i := range items {
		bItems = append(bItems, i)
	}

	w.nav = nav
	w.model = bubblesList.New(bItems, bubblesList.NewDefaultDelegate(), 0, 0)
	w.model.Title = title
	w.model.AdditionalShortHelpKeys = w.KeyBindings
	w.model.AdditionalFullHelpKeys = w.KeyBindings

	return
}

// KeyBindings returns the available additional key bindings
func (w window) KeyBindings() (kb []key.Binding) {
	if w.nav.CanExit() {
		kb = append(kb, key.NewBinding(
			key.WithKeys("b"),
			key.WithHelp("b", "Go back"),
		))
	}

	if w.nav.CanEnter() {
		kb = append(kb, key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "Open section"),
		))
	}

	return kb
}

// Init is used by bubbletea
func (w window) Init() tea.Cmd {
	return nil
}

// Update is used by bubbletea
func (w window) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var newModel tea.Model
	var cmd tea.Cmd

	// Updates the model
	w.model, cmd = w.model.Update(msg)

	switch msg := msg.(type) {

	// Resizes the window when necessary
	case tea.WindowSizeMsg:
		w.model.SetSize(msg.Width, msg.Height)

	case tea.KeyMsg:
		switch key := msg.String(); key {

		// Quits the application when clicking 'ctrl+c'
		case "ctrl+c":
			return w, tea.Quit

		// Tries to exit when clicking 'b'
		case "b":
			if w.nav.CanExit() {
				newModel = w.nav.Exit()
			}

		// Tries to enter when clicking 'enter'
		case "enter":
			if w.nav.CanEnter() {
				if item, ok := w.model.SelectedItem().(Item); ok {
					newModel = w.nav.Enter(item)
				}
			}
		}
	}

	// If a new model is available, returns it; otherwise returns the current one
	if newModel != nil {
		return newModel.Update(tea.WindowSizeMsg{w.model.Width(), w.model.Height()})
	} else {
		return w, cmd
	}
}

// View is used by bubbletea
func (w window) View() string {
	return w.model.View()
}
