package navigator

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Navigator keeps track of what has been chosen, and what needs to be.
// It is composed of a stack and an array of layers
type Navigator struct {
	stack  stack
	layers []Layer
}

// NewNavigator creates a new Navigator for the given layers
func NewNavigator(l []Layer) Navigator {
	return Navigator{stack: stack{}, layers: l}
}

// GetModel returns the current bubbletea model to be rendered
func (n Navigator) GetModel() tea.Model {
	layer := n.layers[n.stack.Level()]
	items := layer.GetAll(n.stack.Last().Id())

	return newWindow(n, n.stack.Print("My Library"), items)
}

// CanEnter returns true if the displayed items can be choosen
func (n Navigator) CanEnter() bool {
	if len(n.layers)-1 > n.stack.Level() {
		return true
	}

	return false
}

// Enter tries to select the current item and jump to the next layer.
// It returns a new model, or nil if something goes wrong
func (n *Navigator) Enter(item Item) tea.Model {
	if n.CanEnter() {
		if _, err := n.layers[n.stack.Level()].GetOne(item.Id()); err == nil {
			n.stack.Enter(item)
			return n.GetModel()
		}
	}

	return nil
}

// CanExit returns true if there is a layer before the current one
func (n Navigator) CanExit() bool {
	if n.stack.Level() > 0 {
		return true
	}

	return false
}

// Exit tries to get to the previous layer. It returns the new model, or nil
// on errors
func (n *Navigator) Exit() tea.Model {
	if n.CanExit() {
		n.stack.Exit()
		return n.GetModel()
	}

	return nil
}
