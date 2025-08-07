package data

import (
	"bytes"
	"errors"
	"image/jpeg"
)

// GetLibraries returns all the libraries.
func GetLibraries() (l []Library) {
	db.Select(&l, `SELECT * FROM libraries;`)
	return l
}

// GetLibrary returns the library with that id.
func GetLibrary(id string) (l Library, err error) {
	if err = db.Get(&l, `SELECT * FROM libraries WHERE id=?;`, id); err == nil {
		return l, nil
	} else {
		return Library{}, errors.New("Library not found")
	}
}

// GetGroups returns all the groups inside of a library.
func GetGroups(library string) (g []Group) {
	db.Select(&g, `SELECT * FROM groups WHERE library=?;`, library)
	return g
}

// GetGroup returns the group with that id.
func GetGroup(id string) (g Group, err error) {
	if err = db.Get(&g, `SELECT * FROM groups WHERE id=?;`, id); err == nil {
		return g, nil
	} else {
		return Group{}, errors.New("Group not found")
	}
}

// GetSagas returns all the sagas inside of a group.
func GetSagas(group string) (s []Saga) {
	db.Select(&s, `SELECT * FROM sagas WHERE "group"=?;`, group)
	return s
}

// GetSaga returns the saga with that id.
func GetSaga(id string) (s Saga, err error) {
	if err = db.Get(&s, `SELECT * FROM sagas WHERE id=?;`, id); err == nil {
		return s, nil
	} else {
		return Saga{}, errors.New("Saga not found")
	}
}

// GetBooks returns all the books inside of a saga. Only the isbn, title and
// authors fields will be fetched. If saga is empty, all the books are fetched.
func GetBooks(saga string) (b []Book) {
	if saga != "" {
		db.Select(&b, `SELECT isbn, title, authors FROM books WHERE saga=?;`, saga)
	} else {
		db.Select(&b, `SELECT isbn, title, authors FROM books;`)
	}

	return b
}

// GetBook returns a book from the isbn. It also tries to fetch the thumbnail,
// but if some error occours during that phase it will return a nil one.
func GetBook(isbn string) (b Book, err error) {
	err = db.Get(&b, `SELECT * FROM books WHERE isbn=?;`, isbn)
	if err != nil {
		return Book{}, errors.New("Book not found")
	}

	// Fetches and loads the image
	var data []byte
	err = db.Get(&data, `SELECT data FROM thumbnails WHERE isbn=?;`, isbn)
	if err == nil {
		if img, err := jpeg.Decode(bytes.NewReader(data)); err == nil {
			b.Thumbnail = img
		}
	}

	return b, nil
}
