package main

import (
	"books/components/fields"
	"books/components/list"
	"books/components/navigator"
	"books/components/tabs"
	"books/data"

	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

var (
	library data.Library
	group   data.Group
	saga    data.Saga
	book    data.Book
)

func newLibrariesList() tea.Model {
	return navigator.NewNavigator(
		nil,

		list.NewListModel(getTitle(), librariesSupplier{}),

		func(id string) tea.Model {
			library, _ = data.GetLibrary(id)
			return newGroupsList()
		},
	)
}

func newGroupsList() tea.Model {
	return navigator.NewNavigator(
		func() tea.Model {
			library = data.Library{}
			return newLibrariesList()
		},

		list.NewListModel(getTitle(), groupsSupplier{library: library.Id}),

		func(id string) tea.Model {
			group, _ = data.GetGroup(id)
			return newSagasList()
		},
	)
}

func newSagasList() tea.Model {
	return navigator.NewNavigator(
		func() tea.Model {
			group = data.Group{}
			return newGroupsList()
		},

		list.NewListModel(getTitle(), sagasSupplier{group: group.Id}),

		func(id string) tea.Model {
			saga, _ = data.GetSaga(id)
			return newBooksList()
		},
	)
}

func newBooksList() tea.Model {
	return navigator.NewNavigator(
		func() tea.Model {
			saga = data.Saga{}
			return newSagasList()
		},

		list.NewListModel(
			fmt.Sprintf("My Library > %s > %s > %s", library.Name, group.Name, saga.Name),
			booksSupplier{saga: saga.Id},
		),

		func(id string) tea.Model {
			book, _ = data.GetBook(id)
			return newBookDetails()
		},
	)
}

func newBookDetails() tea.Model {
	return navigator.NewNavigator(
		func() tea.Model {
			book = data.Book{}
			return newBooksList()
		},

		tabs.NewTabbedModel([]tabs.Tab{
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
