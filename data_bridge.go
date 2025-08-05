package main

import (
	"books/components/list"
	"books/data"

	"strings"
)

type librariesSupplier struct{}

func (_ librariesSupplier) parse(l data.Library) list.Item {
	return list.NewItem(l.Id, l.Name, "")
}

func (ls librariesSupplier) GetOne(id string) (list.Item, error) {
	l, e := data.GetLibrary(id)
	return ls.parse(l), e
}

func (ls librariesSupplier) GetAll() (items []list.Item) {
	for _, library := range data.GetLibraries() {
		items = append(items, ls.parse(library))
	}
	return
}

type groupsSupplier struct{ library string }

func (_ groupsSupplier) parse(g data.Group) list.Item {
	return list.NewItem(g.Id, g.Name, "")
}

func (gs groupsSupplier) GetOne(id string) (list.Item, error) {
	g, e := data.GetGroup(id)
	return gs.parse(g), e
}

func (gs groupsSupplier) GetAll() (items []list.Item) {
	for _, group := range data.GetGroups(gs.library) {
		items = append(items, gs.parse(group))
	}
	return
}

type sagasSupplier struct{ group string }

func (_ sagasSupplier) parse(s data.Saga) list.Item {
	return list.NewItem(s.Id, s.Name, "")
}

func (ss sagasSupplier) GetOne(id string) (list.Item, error) {
	s, e := data.GetSaga(id)
	return ss.parse(s), e
}

func (ss sagasSupplier) GetAll() (items []list.Item) {
	for _, saga := range data.GetSagas(ss.group) {
		items = append(items, ss.parse(saga))
	}
	return
}

type booksSupplier struct{ saga string }

func (_ booksSupplier) parse(b data.Book) list.Item {
	authors := strings.Join(b.Authors, ", ")
	return list.NewItem(b.Isbn, b.Title, authors)
}

func (bs booksSupplier) GetOne(id string) (list.Item, error) {
	b, e := data.GetBook(id)
	return bs.parse(b), e
}

func (bs booksSupplier) GetAll() (items []list.Item) {
	for _, book := range data.GetBooks(bs.saga) {
		items = append(items, bs.parse(book))
	}
	return
}
