package main

import (
	"books/data"

	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

// Program is something that can return a bubbletea model. A program is chosen
// by the user on each run as a command line argument.
type Program interface {
	GetRootModel() (tea.Model, error)
}

// Programs contains all the known programs.
var Programs = map[string]Program{
	"hierarchy": Hierarchy{},
	"flat":      Flat{},
	"book":      Book{},
}

// main runs a program chosen by the user.
// The required syntax for running this application is
// ./books <dbPath> <program> [additionalParams...]
// To see the list of available programs you can run
// ./books <dbPath> list
// (in this case the dbPath is ignored).
func main() {
	var dbPath, programName string

	// Checks the arguments
	if len(os.Args) < 3 {
		Error("Usage: ./books <dbPath> <programName|\"list\"> [additionalParams...]")
	} else {
		dbPath = os.Args[1]
		programName = os.Args[2]
	}

	// Prints the list of available programs
	if programName == "list" {
		fmt.Println("The known programs are: ")
		for p, _ := range Programs {
			fmt.Println("-", p)
		}

		return
	}

	// Makes sure the program exists
	if _, ok := Programs[programName]; !ok {
		Error("Unrecognized program.")
	}

	// Initialize the database
	if err := data.InitDB(dbPath); err != nil {
		Error("Error initializing the database:", err)
	}

	// Run the desired program
	model, err := Programs[programName].GetRootModel()
	if err == nil {
		_, err = tea.NewProgram(model, tea.WithAltScreen()).Run()
	}

	if err != nil {
		Error(err)
	}
}

// Error shows an error then terminates.
func Error(data ...any) {
	fmt.Println(data...)
	os.Exit(1)
}
