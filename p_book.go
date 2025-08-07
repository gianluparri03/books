package main

import (
	"books/data"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// Book is a program that shows the detail of an already chosen book.
// The book isbn is fetched from an argument.
type Book struct{}

func (_ Book) GetRootModel() (tea.Model, error) {
	// Makes sure the isbn is specified
	if len(os.Args) < 4 {
		Error("Usage: ./books <dbPath> book <isbn>")
	}

	// Makes sure the book exists
	isbn := os.Args[3]
	b, err := data.GetBook(isbn)
	if err != nil {
		return nil, err
	}

	s, _ := data.GetSaga(b.Saga)
	g, _ := data.GetGroup(s.Group)
	l, _ := data.GetLibrary(g.Library)
	return bookDetailsModel(l, g, s, b, false), nil
}
