package navigator

import (
	"github.com/charmbracelet/bubbles/key"
	bubblesList "github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

// NavigatorModel is a bubbletea model that lets the user choose from a nested
// list. It's works with a Navigator; the list model is the bubbles' one
type NavigatorModel struct {
	navigator Navigator
	model     bubblesList.Model

	then func(string) tea.Model
}

// NewNavigatorModel returns a new NavigatorModel based on the given navigator,
// that will render the model returned by then when the final choice will be
// made
func NewNavigatorModel(navigator Navigator, then func(string) tea.Model) tea.Model {
	if !navigator.Completed() {
		return newNavigatorModel(navigator, then)
	} else {
		return then(navigator.Last().Id())
	}
}

// newNavigatorModel returns an actual NavigatorModel
func newNavigatorModel(navigator Navigator, then func(string) tea.Model) tea.Model {
	nm := NavigatorModel{navigator: navigator, then: then}

	var bItems []bubblesList.Item
	for _, i := range navigator.GetOptions() {
		bItems = append(bItems, i)
	}

	nm.model = bubblesList.New(bItems, bubblesList.NewDefaultDelegate(), 0, 0)
	nm.model.Title = navigator.Print()
	nm.model.AdditionalShortHelpKeys = nm.KeyBindings
	nm.model.AdditionalFullHelpKeys = nm.KeyBindings
	return nm
}

// KeyBindings returns the available additional key bindings
func (nm NavigatorModel) KeyBindings() (kb []key.Binding) {
	if nm.navigator.CanGoBack() {
		kb = append(kb, key.NewBinding(
			key.WithKeys("b"),
			key.WithHelp("b", "Go back"),
		))
	}

	kb = append(kb, key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "Open section"),
	))

	return kb
}

// Init is used by bubbletea
func (nm NavigatorModel) Init() tea.Cmd {
	return tea.WindowSize()
}

// Update is used by bubbletea
func (nm NavigatorModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Resizes the window when necessary
	case tea.WindowSizeMsg:
		nm.model.SetSize(msg.Width, msg.Height)

	case tea.KeyMsg:
		switch key := msg.String(); key {

		// Quits the application when clicking 'ctrl+c'
		case "ctrl+c":
			return nm, tea.Quit

		// Tries to exit when clicking 'b'
		case "b":
			if nm.navigator.CanGoBack() {
				if nm.navigator.GoBack() {
					return NewNavigatorModel(nm.navigator, nm.then), tea.WindowSize()
				}
			}

		// Tries to enter when clicking 'enter'
		case "enter":
			if item, ok := nm.model.SelectedItem().(Item); ok {
				if nm.navigator.Choose(item) {
					return NewNavigatorModel(nm.navigator, nm.then), tea.WindowSize()
				}
			}
		}
	}

	// Otherwise updates and returns the current one
	var cmd tea.Cmd
	nm.model, cmd = nm.model.Update(msg)
	return nm, cmd
}

// View is used by bubbletea
func (nm NavigatorModel) View() string {
	return nm.model.View()
}
