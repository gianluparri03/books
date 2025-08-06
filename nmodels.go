package main

import (
	"books/components/navigator"
	"books/data"

	tea "github.com/charmbracelet/bubbletea"
)

// librariesListNModel wraps librariesListModel
func librariesListNModel() tea.Model {
	return navigator.New(
		nil,

		librariesListModel(),

		func(id string) tea.Model {
			library, _ = data.GetLibrary(id)
			return groupsListNModel()
		},
	)
}

// groupsListNModel wraps groupsListModel
func groupsListNModel() tea.Model {
	return navigator.New(
		func() tea.Model {
			library = data.Library{}
			return librariesListNModel()
		},

		groupsListModel(),

		func(id string) tea.Model {
			group, _ = data.GetGroup(id)
			return sagasListNModel()
		},
	)
}

// sagasListNModel wraps sagasListModel
func sagasListNModel() tea.Model {
	return navigator.New(
		func() tea.Model {
			group = data.Group{}
			return groupsListNModel()
		},

		sagasListModel(),

		func(id string) tea.Model {
			saga, _ = data.GetSaga(id)
			return booksListNModel()
		},
	)
}

// booksListNModel wraps booksListModel
func booksListNModel() tea.Model {
	return navigator.New(
		func() tea.Model {
			saga = data.Saga{}
			return sagasListNModel()
		},

		booksListModel(),

		func(id string) tea.Model {
			book, _ = data.GetBook(id)
			return bookDetailsNModel()
		},
	)
}

// bookDetailsNModel wraps bookDetailsModel
func bookDetailsNModel() tea.Model {
	return navigator.New(
		func() tea.Model {
			book = data.Book{}
			return booksListNModel()
		},

		bookDetailsModel(),

		nil,
	)
}
