package main

import (
	"books/components/fields"
	"books/components/list"
	"books/components/navigator"
	"books/components/tabs"
	"books/data"

	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

var (
	library data.Library
	group   data.Group
	saga    data.Saga
	book    data.Book
)

func HierarchyView() tea.Model {
	return newLibrariesList()
}

func newLibrariesList() tea.Model {
	var items []list.Item
	for _, l := range data.GetLibraries() {
		var groups strings.Builder
		for i, g := range data.GetGroups(l.Id) {
			if i > 0 {
				groups.WriteString(", ")
			}
			groups.WriteString(g.Name)
		}

		items = append(items, list.NewItem(l.Id, l.Name, groups.String()))
	}

	return navigator.New(
		nil,

		list.New(getTitle(), items, false, true),

		func(id string) tea.Model {
			library, _ = data.GetLibrary(id)
			return newGroupsList()
		},
	)
}

func newGroupsList() tea.Model {
	var items []list.Item
	for _, g := range data.GetGroups(library.Id) {
		var sagas strings.Builder
		for i, s := range data.GetSagas(g.Id) {
			if i > 0 {
				sagas.WriteString(", ")
			}
			sagas.WriteString(s.Name)
		}

		items = append(items, list.NewItem(g.Id, g.Name, sagas.String()))
	}

	return navigator.New(
		func() tea.Model {
			library = data.Library{}
			return newLibrariesList()
		},

		list.New(getTitle(), items, true, true),

		func(id string) tea.Model {
			group, _ = data.GetGroup(id)
			return newSagasList()
		},
	)
}

func newSagasList() tea.Model {
	var items []list.Item
	for _, s := range data.GetSagas(group.Id) {
		var books strings.Builder
		for i, b := range data.GetBooks(s.Id) {
			if i > 0 {
				books.WriteString(", ")
			}
			books.WriteString(b.Title)
		}

		items = append(items, list.NewItem(s.Id, s.Name, books.String()))
	}

	return navigator.New(
		func() tea.Model {
			group = data.Group{}
			return newGroupsList()
		},

		list.New(getTitle(), items, true, true),

		func(id string) tea.Model {
			saga, _ = data.GetSaga(id)
			return newBooksList()
		},
	)
}

func newBooksList() tea.Model {
	var items []list.Item
	for _, b := range data.GetBooks(saga.Id) {
		authors := strings.Join(b.Authors, ", ")
		items = append(items, list.NewItem(b.Isbn, b.Title, authors))
	}

	return navigator.New(
		func() tea.Model {
			saga = data.Saga{}
			return newSagasList()
		},

		list.New(getTitle(), items, true, true),

		func(id string) tea.Model {
			book, _ = data.GetBook(id)
			return newBookDetails()
		},
	)
}

func newBookDetails() tea.Model {
	return navigator.New(
		func() tea.Model {
			book = data.Book{}
			return newBooksList()
		},

		tabs.New([]tabs.Tab{
			tabs.Tab{
				Title: "General",
				Model: fields.New(
					fields.Field{Label: "Title", Value: book.Title},
					fields.Field{Label: "Authors", Value: strings.Join(book.Authors, "\n")},
					fields.Field{Label: "Publisher", Value: book.Publisher},
				),
			},

			tabs.Tab{
				Title: "Details",
				Model: fields.New(
					fields.Field{Label: "Pages", Value: book.Pages},
					fields.Field{Label: "Price", Value: book.Price},
					fields.Field{Label: "Language", Value: book.Lang},
					fields.Field{Label: "ISBN", Value: book.Isbn},
				),
			},

			tabs.Tab{
				Title: "Location",
				Model: fields.New(
					fields.Field{Label: "Library", Value: book.Library},
					fields.Field{Label: "Group", Value: book.Group},
					fields.Field{Label: "Saga", Value: book.Saga},
					fields.Field{Label: "Number", Value: book.Number},
				),
			},

			tabs.Tab{
				Title: "Status",
				Model: fields.New(
					fields.Field{Label: "Bought on", Value: book.BoughtDate, Inline: true},
					fields.Field{Label: "Bought from", Value: book.BoughtShop, Inline: true},
					fields.Field{Label: "Started reading on", Value: book.StartedDate, Inline: true},
					fields.Field{Label: "Finished reading on", Value: book.FinishedDate, Inline: true},
					fields.Field{Label: "Status", Value: book.Status, Inline: true},
				),
			},
		}),

		nil,
	)
}

func getTitle() string {
	var sb strings.Builder
	sb.WriteString("My Library")

	if library.Id != "" {
		sb.WriteString(" > ")
		sb.WriteString(library.Name)
	}

	if group.Id != "" {
		sb.WriteString(" > ")
		sb.WriteString(group.Name)
	}

	if saga.Id != "" {
		sb.WriteString(" > ")
		sb.WriteString(saga.Name)
	}

	return sb.String()
}
