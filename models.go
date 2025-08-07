package main

import (
	"books/components/fields"
	"books/components/list"
	"books/components/preview"
	"books/components/tabs"
	"books/data"

	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// librariesListModel is a List model used to choose a library.
func librariesListModel(title string, canGoBack, canEnter bool) tea.Model {
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

	return list.New(title, items, canGoBack, canEnter)
}

// groupsListModel is a List model used to choose a group.
func groupsListModel(title string, library data.Library, canGoBack, canEnter bool) tea.Model {
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

	return list.New(title, items, canGoBack, canEnter)
}

// sagasListModel is a List model used to choose a saga.
func sagasListModel(title string, group data.Group, canGoBack, canEnter bool) tea.Model {
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

	return list.New(title, items, canGoBack, canEnter)
}

// booksListModel is a List model used to choose a book.
func booksListModel(title string, saga data.Saga, canGoBack, canEnter bool) tea.Model {
	var items []list.Item
	for _, b := range data.GetBooks(saga.Id) {
		items = append(items, list.NewItem(b.Isbn, b.Title, b.Authors))
	}

	return list.New(title, items, canGoBack, canEnter)
}

// bookDetailsModel is made of a Preview model, a Tabbed model and some
// Fields model.
func bookDetailsModel(l data.Library, g data.Group, s data.Saga, b data.Book, canGoBack bool) tea.Model {
	return preview.New(
		b.Thumbnail,
		tabs.New(
			canGoBack,
			[]tabs.Tab{
				tabs.Tab{
					Title: "General",
					Model: fields.New(
						fields.Field{Label: "Title", Value: b.Title},
						fields.Field{Label: "Authors", Value: b.Authors},
						fields.Field{Label: "Publisher", Value: b.Publisher},
					),
				},

				tabs.Tab{
					Title: "Details",
					Model: fields.New(
						fields.Field{Label: "Pages", Value: b.Pages},
						fields.Field{Label: "Price", Value: b.Price},
						fields.Field{Label: "Language", Value: b.Lang},
						fields.Field{Label: "ISBN", Value: b.Isbn},
					),
				},

				tabs.Tab{
					Title: "Location",
					Model: fields.New(
						fields.Field{Label: "Library", Value: l.Name},
						fields.Field{Label: "Group", Value: g.Name},
						fields.Field{Label: "Saga", Value: s.Name},
						fields.Field{Label: "Number", Value: b.Number},
					),
				},

				tabs.Tab{
					Title: "Status",
					Model: fields.New(
						fields.Field{Label: "Bought on", Value: b.BoughtDate, Inline: true},
						fields.Field{Label: "Bought from", Value: b.BoughtShop, Inline: true},
						fields.Field{Label: "Started reading on", Value: b.StartedDate, Inline: true},
						fields.Field{Label: "Finished reading on", Value: b.FinishedDate, Inline: true},
						fields.Field{Label: "Status", Value: b.Status, Inline: true},
					),
				},
			},
		),
	)
}
