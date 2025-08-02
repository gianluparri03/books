package main

import (
	"books/components/fields"
	"books/components/tabs"
	"books/data"

	tea "github.com/charmbracelet/bubbletea"
	"strings"
)

// NewBookViewer returns a tea model that shows the details of a book
func NewBookViewer(isbn string) tea.Model {
	book, _ := data.GetBook(isbn)

	return tabs.NewTabbedModel([]tabs.Tab{
		tabs.Tab{
			Title: "General",
			Model: fields.NewFieldsModel(
				fields.Field{Label: "Title", Value: book.Title},
				fields.Field{Label: "Authors", Value: strings.Join(book.Authors, "\n")},
				fields.Field{Label: "Publisher", Value: book.Publisher},
			),
		},

		tabs.Tab{
			Title: "Details",
			Model: fields.NewFieldsModel(
				fields.Field{Label: "Pages", Value: book.Pages},
				fields.Field{Label: "Price", Value: book.Price},
				fields.Field{Label: "Language", Value: book.Lang},
				fields.Field{Label: "ISBN", Value: book.Isbn},
			),
		},

		tabs.Tab{
			Title: "Location",
			Model: fields.NewFieldsModel(
				fields.Field{Label: "Library", Value: book.Library},
				fields.Field{Label: "Group", Value: book.Group},
				fields.Field{Label: "Saga", Value: book.Saga},
				fields.Field{Label: "Number", Value: book.Number},
			),
		},

		tabs.Tab{
			Title: "Status",
			Model: fields.NewFieldsModel(
				fields.Field{Label: "Bought on", Value: book.BoughtDate, Inline: true},
				fields.Field{Label: "Bought from", Value: book.BoughtShop, Inline: true},
				fields.Field{Label: "Started reading on", Value: book.StartedDate, Inline: true},
				fields.Field{Label: "Finished reading on", Value: book.FinishedDate, Inline: true},
				fields.Field{Label: "Status", Value: book.Status, Inline: true},
			),
		},
	})
}
