package main

import (
	"books/components/navigator"
	"books/data"

	tea "github.com/charmbracelet/bubbletea"
	"strings"
)

// Hierarchy is a program that lets the user choose a book through the
// hierarchy, therefore it'll have to choose a library, a group, a saga and
// finally the book to see the details.
type Hierarchy struct {
	l data.Library
	g data.Group
	s data.Saga
	b data.Book
}

func (h Hierarchy) GetRootModel() (tea.Model, error) {
	return h.getLibrariesModel(), nil
}

// getLibrariesModel wraps librariesListModel.
func (h *Hierarchy) getLibrariesModel() tea.Model {
	return navigator.New(
		nil,

		librariesListModel(h.getTitle(), false, true),

		func(id string) tea.Model {
			h.l, _ = data.GetLibrary(id)
			return h.getGroupsModel()
		},
	)
}

// getGroupsModel wraps groupsListModel.
func (h *Hierarchy) getGroupsModel() tea.Model {
	return navigator.New(
		func() tea.Model {
			h.l = data.Library{}
			return h.getLibrariesModel()
		},

		groupsListModel(h.getTitle(), h.l, true, true),

		func(id string) tea.Model {
			h.g, _ = data.GetGroup(id)
			return h.getSagasModel()
		},
	)
}

// getSagasModel wraps sagasListModel.
func (h *Hierarchy) getSagasModel() tea.Model {
	return navigator.New(
		func() tea.Model {
			h.g = data.Group{}
			return h.getGroupsModel()
		},

		sagasListModel(h.getTitle(), h.g, true, true),

		func(id string) tea.Model {
			h.s, _ = data.GetSaga(id)
			return h.getBooksModel()
		},
	)
}

// getBooksModel wraps booksListModel.
func (h *Hierarchy) getBooksModel() tea.Model {
	return navigator.New(
		func() tea.Model {
			h.s = data.Saga{}
			return h.getSagasModel()
		},

		booksListModel(h.getTitle(), h.s, true, true),

		func(id string) tea.Model {
			h.b, _ = data.GetBook(id)
			return h.getDetailsModel()
		},
	)
}

// getDetailsModel wraps bookDetailsModel.
func (h *Hierarchy) getDetailsModel() tea.Model {
	return navigator.New(
		func() tea.Model {
			h.b = data.Book{}
			return h.getBooksModel()
		},

		bookDetailsModel(h.l, h.g, h.s, h.b, true),

		nil,
	)
}

// getTitle returns a title containing the chosen items.
func (h Hierarchy) getTitle() string {
	var sb strings.Builder
	sb.WriteString("My Library")

	if h.l.Id != "" {
		sb.WriteString(" > ")
		sb.WriteString(h.l.Name)

		if h.g.Id != "" {
			sb.WriteString(" > ")
			sb.WriteString(h.g.Name)

			if h.s.Id != "" {
				sb.WriteString(" > ")
				sb.WriteString(h.s.Name)
			}
		}
	}

	return sb.String()
}
