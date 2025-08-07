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
	if len(os.Args) < 2 {
		fmt.Println("Please specify the database path.")
		os.Exit(1)
	}

	if err := data.InitDB(os.Args[1]); err != nil {
		fmt.Println("Error initializing the database:", err)
		os.Exit(1)
	}

	p := tea.NewProgram(librariesListNModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
