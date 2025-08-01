package navigator

// stack is a lifo data structure
type stack []Item

// Enter adds an item to the stack
func (s *stack) Enter(item Item) {
	*s = append(*s, item)
}

// Exit pops an item from the stack
func (s *stack) Exit() {
	if len(*s) > 0 {
		*s = (*s)[:len(*s)-1]
	}
}

// Print returns a visual representation of the stack
func (s stack) Print(head string) string {
	str := head

	for _, item := range s {
		str += " > " + item.Title()
	}

	return str
}

// Level returns the number of items on the stack
func (s stack) Level() int {
	return len(s)

}

// Last returns the last item on the stack
func (s stack) Last() Item {
	if s.Level() > 0 {
		return s[len(s)-1]
	} else {
		return Item{}
	}
}
