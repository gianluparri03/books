package main

import (
	"books/components/fields"
	"books/components/list"
	"books/components/preview"
	"books/components/tabs"
	"books/data"

	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// librariesListModel is a List model used to choose a library.
func librariesListModel() tea.Model {
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

	title := fmt.Sprintf("My library")
	return list.New(title, items, false, true)
}

// groupsListModel is a List model used to choose a group.
func groupsListModel() tea.Model {
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

	title := fmt.Sprintf("My library > %s", library.Name)
	return list.New(title, items, true, true)
}

// sagasListModel is a List model used to choose a saga.
func sagasListModel() tea.Model {
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

	title := fmt.Sprintf("My library > %s > %s", library.Name, group.Name)
	return list.New(title, items, true, true)
}

// booksListModel is a List model used to choose a book.
func booksListModel() tea.Model {
	var items []list.Item
	for _, b := range data.GetBooks(saga.Id) {
		items = append(items, list.NewItem(b.Isbn, b.Title, b.Authors))
	}

	title := fmt.Sprintf(
		"My library > %s > %s > %s",
		library.Name, group.Name, saga.Name,
	)
	return list.New(title, items, true, true)
}

// bookDetailsModel is made of a Preview model, a Tabbed model and some
// Fields model.
func bookDetailsModel() tea.Model {
	return preview.New(
		book.Thumbnail,
		tabs.New([]tabs.Tab{
			tabs.Tab{
				Title: "General",
				Model: fields.New(
					fields.Field{Label: "Title", Value: book.Title},
					fields.Field{Label: "Authors", Value: book.Authors},
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
					fields.Field{Label: "Library", Value: library.Name},
					fields.Field{Label: "Group", Value: group.Name},
					fields.Field{Label: "Saga", Value: saga.Name},
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
	)
}
