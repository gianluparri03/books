package main

import (
	"books/data"

	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

var (
	// chosen items
	library data.Library
	group   data.Group
	saga    data.Saga
	book    data.Book
)

func main() {
	p := tea.NewProgram(librariesListNModel(), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
