package list

// ItemSupplier is something that can returns all the Items and check if
// an Item exists
type ItemSupplier interface {
	GetOne(id string) (Item, error)
	GetAll() []Item
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
