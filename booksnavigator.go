package main

import (
	"books/components/navigator"
	"books/data"

	"strings"
)

// BooksNavigator is NavigatorModel that lets the user choose
// a book, through a library, a group and a saga
var BooksNavigator = navigator.NewNavigatorModel(
	navigator.NewNavigator(
		"My Library",
		[]navigator.Layer{
			librariesLayer{},
			groupsLayer{},
			sagasLayer{},
			booksLayer{},
		},
	),
	NewBookViewer,
)

type librariesLayer struct{}

func (_ librariesLayer) parse(l data.Library) navigator.Item {
	return navigator.NewItem(l.Id, l.Name, "")
}

func (ll librariesLayer) GetOne(id string) (navigator.Item, error) {
	l, e := data.GetLibrary(id)
	return ll.parse(l), e
}

func (ll librariesLayer) GetAll(parent string) (items []navigator.Item) {
	for _, library := range data.GetLibraries() {
		items = append(items, ll.parse(library))
	}
	return
}

type groupsLayer struct{}

func (_ groupsLayer) parse(g data.Group) navigator.Item {
	return navigator.NewItem(g.Id, g.Name, "")
}

func (gl groupsLayer) GetOne(id string) (navigator.Item, error) {
	g, e := data.GetGroup(id)
	return gl.parse(g), e
}

func (gl groupsLayer) GetAll(parent string) (items []navigator.Item) {
	for _, group := range data.GetGroups(parent) {
		items = append(items, gl.parse(group))
	}
	return
}

type sagasLayer struct{}

func (_ sagasLayer) parse(s data.Saga) navigator.Item {
	return navigator.NewItem(s.Id, s.Name, "")
}

func (sl sagasLayer) GetOne(id string) (navigator.Item, error) {
	s, e := data.GetSaga(id)
	return sl.parse(s), e
}

func (sl sagasLayer) GetAll(parent string) (items []navigator.Item) {
	for _, saga := range data.GetSagas(parent) {
		items = append(items, sl.parse(saga))
	}
	return
}

type booksLayer struct{}

func (_ booksLayer) parse(b data.Book) navigator.Item {
	authors := strings.Join(b.Authors, ", ")
	return navigator.NewItem(b.Isbn, b.Title, authors)
}

func (bl booksLayer) GetOne(id string) (navigator.Item, error) {
	b, e := data.GetBook(id)
	return bl.parse(b), e
}

func (bl booksLayer) GetAll(parent string) (items []navigator.Item) {
	for _, book := range data.GetBooks(parent) {
		items = append(items, bl.parse(book))
	}
	return
}
