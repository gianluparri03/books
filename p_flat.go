package main

import (
	"books/components/navigator"
	"books/data"

	tea "github.com/charmbracelet/bubbletea"
)

// Flat is a program that lets the user choose a book from the list of all
// books.
type Flat struct {
	b data.Book
}

func (f Flat) GetRootModel() (tea.Model, error) {
	return f.getListModel(), nil
}

// getListModel wraps booksListModel
func (f *Flat) getListModel() tea.Model {
	return navigator.New(
		nil,

		booksListModel("My Library", data.Saga{}, false, true),

		func(id string) tea.Model {
			f.b, _ = data.GetBook(id)
			return f.getDetailsModel()
		},
	)
}

// getDetailsModel wraps bookDetailsModel
func (f *Flat) getDetailsModel() tea.Model {
	s, _ := data.GetSaga(f.b.Saga)
	g, _ := data.GetGroup(s.Group)
	l, _ := data.GetLibrary(g.Library)

	return navigator.New(
		func() tea.Model {
			f.b = data.Book{}
			return f.getListModel()
		},

		bookDetailsModel(l, g, s, f.b, true),

		nil,
	)
}
