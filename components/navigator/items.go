package navigator

// Layer is something to choose from, with a GetAll method, to get all
// the options (of a parent), and a GetOne method, to get a single item
// and esnure the id is correct
type Layer interface {
	GetOne(id string) (Item, error)
	GetAll(parent string) []Item
}

// Item is an item that can be choosen from a list
type Item struct {
	id          string
	title       string
	description string
}

// NewItem creates a new item
func NewItem(id, title, description string) Item {
	return Item{id, title, description}
}

// Id returns the item's identifier
func (i Item) Id() string { return i.id }

// Title returns the item's title, shown in the list
func (i Item) Title() string { return i.title }

// Description returns the item's description, shown in the list
func (i Item) Description() string { return i.description }

// FilterValue returns a string used when filtering
func (i Item) FilterValue() string { return i.title + " " + i.description }
