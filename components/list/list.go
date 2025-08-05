package list

import (
	"books/components/navigator"

	bubblesList "github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

// ListModel is a bubbletea model that lets the user choose from a list
type ListModel struct {
	supplier ItemSupplier
	model    bubblesList.Model
	selected Item
}

// NewListModel returns a new ListModel
func NewListModel(title string, supplier ItemSupplier) tea.Model {
	lm := ListModel{supplier: supplier}

	var bItems []bubblesList.Item
	for _, i := range supplier.GetAll() {
		bItems = append(bItems, i)
	}

	lm.model = bubblesList.New(bItems, bubblesList.NewDefaultDelegate(), 0, 0)
	lm.model.Title = title
	return lm
}

// Init is used by bubbletea
func (lm ListModel) Init() tea.Cmd {
	return tea.WindowSize()
}

// NUpdate is used by Navigator and bubbletea
func (lm ListModel) NUpdate(msg tea.Msg) (tea.Model, tea.Cmd, navigator.Jump) {
	switch msg := msg.(type) {

	// Resizes the window when necessary
	case tea.WindowSizeMsg:
		lm.model.SetSize(msg.Width, msg.Height)

	case tea.KeyMsg:
		switch key := msg.String(); key {

		// Quits the application when clicking 'ctrl+c'
		case "ctrl+c":
			return lm, tea.Quit, navigator.Jump{}

		// Tries to exit when clicking 'b'
		case "b":
			return lm, nil, navigator.Jump{Prev: true}

		// Tries to enter when clicking 'enter'
		case "enter":
			if item, ok := lm.model.SelectedItem().(Item); ok {
				return lm, nil, navigator.Jump{Next: true, NextArg: item.Id()}
			}
		}
	}

	// Otherwise updates and returns the current one
	var cmd tea.Cmd
	lm.model, cmd = lm.model.Update(msg)
	return lm, cmd, navigator.Jump{}
}

// Update is used by bubbletea
func (lm ListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m, c, _ := lm.NUpdate(msg)
	return m, c
}

// View is used by bubbletea
func (lm ListModel) View() string {
	return lm.model.View()
}
