package data

import (
	"errors"
	"image"
	"image/jpeg"
	"os"
)

// GetLibraries returns all the libraries.
func GetLibraries() []Library {
	return libraries
}

// GetLibrary returns the library with that id.
func GetLibrary(id string) (Library, error) {
	for _, l := range libraries {
		if l.Id == id {
			return l, nil
		}
	}

	return Library{}, errors.New("library not found")
}

// GetGroups returns all the groups inside of a library.
func GetGroups(library string) []Group {
	var found []Group
	for _, g := range groups {
		if g.Library == library || library == "" {
			found = append(found, g)
		}
	}

	return found
}

// GetGroup returns the group with that id.
func GetGroup(id string) (Group, error) {
	for _, l := range groups {
		if l.Id == id {
			return l, nil
		}
	}

	return Group{}, errors.New("group not found")
}

// GetSagas returns all the sagas inside of a group.
func GetSagas(group string) []Saga {
	var found []Saga
	for _, s := range sagas {
		if s.Group == group || group == "" {
			found = append(found, s)
		}
	}

	return found
}

// GetSaga returns the saga with that id.
func GetSaga(id string) (Saga, error) {
	for _, s := range sagas {
		if s.Id == id {
			return s, nil
		}
	}

	return Saga{}, errors.New("saga not found")
}

// GetBooks returns all the books inside of a saga.
func GetBooks(saga string) []Book {
	var found []Book
	for _, b := range books {
		if b.Saga == saga || saga == "" {
			found = append(found, b)
		}
	}

	return found
}

// GetBook returns a book from the isbn. It also tries to fetch the thumbnail,
// but if some error occours during that phase it will return a nil one.
func GetBook(isbn string) (Book, error) {
	for _, b := range books {
		if b.Isbn == isbn {
			if thumb, err := getBookThumb(b.Isbn); err == nil {
				b.Thumbnail = thumb
			}

			return b, nil
		}
	}

	return Book{}, errors.New("book not found")
}

// getBookThumb returns the thumbnail of the book with the given isbn.
// It supposes that all the thumbnails are in the data/thumbs/ folder, the
// filename is the isbn and the extension is .jpg.
func getBookThumb(isbn string) (image.Image, error) {
	f, err := os.Open("data/thumbs/" + isbn + ".jpg")
	defer f.Close()
	if err != nil {
		return nil, err
	}

	return jpeg.Decode(f)
}
