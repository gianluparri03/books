package list

// Item is an item that can be choosen from the list.
type Item struct {
	id          string
	title       string
	description string
}

func NewItem(id, title, description string) Item {
	return Item{id, title, description}
}

func (i Item) Id() string { return i.id }

func (i Item) Title() string { return i.title }

func (i Item) Description() string { return i.description }

// FilterValue returns the value on which the filters may be applied
func (i Item) FilterValue() string { return i.title + " " + i.description }
