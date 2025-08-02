package navigator

// Navigator contains all the possible options for each layer, and keeps
// track of what has been chosen, and what needs to be.
// It is composed of a stack, an array of layers and
// and a final model rendered
// when the final choice has been made
type Navigator struct {
	rootName string
	stack    []Item
	layers   []Layer
}

// NewNavigator creates a new Navigator for the given layers
func NewNavigator(rootName string, l []Layer) Navigator {
	return Navigator{rootName: rootName, layers: l}
}

// GetOptions returns the options to choose from, in the current layer
func (n Navigator) GetOptions() []Item {
	return n.layers[len(n.stack)].GetAll(n.Last().Id())
}

// Choose tries to select the current item and jump to the next layer.
// Returns true if no errors occurred
func (n *Navigator) Choose(item Item) bool {
	if _, err := n.layers[len(n.stack)].GetOne(item.Id()); err == nil {
		n.stack = append(n.stack, item)
		return true
	}

	return false
}

// CanGoBack returns true if there is a layer before the current one
func (n Navigator) CanGoBack() bool {
	return len(n.stack) > 0
}

// GoBack tries to get to the previous layer.
// Returns true if no errors occurred
func (n *Navigator) GoBack() bool {
	if n.CanGoBack() {
		n.stack = n.stack[:len(n.stack)-1]
		return true
	}

	return false
}

// Completed returns true if all the choices has been made
func (n Navigator) Completed() bool {
	return len(n.layers) == len(n.stack)
}

// Print returns a visual representation of the choices
func (n Navigator) Print() string {
	str := n.rootName

	for _, item := range n.stack {
		str += " > " + item.Title()
	}

	return str
}

// Last returns the last choosen item
func (n Navigator) Last() Item {
	if len(n.stack) > 0 {
		return n.stack[len(n.stack)-1]
	} else {
		return Item{}
	}
}
