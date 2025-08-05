package data

import "errors"

// GetLibraries returns all the libraries
func GetLibraries() []Library {
	return libraries
}

// GetLibrary returns a library from the id
func GetLibrary(id string) (Library, error) {
	for _, l := range libraries {
		if l.Id == id {
			return l, nil
		}
	}

	return Library{}, errors.New("library not found")
}

// GetGroups returns all the groups of a library
func GetGroups(library string) []Group {
	var found []Group
	for _, g := range groups {
		if g.Library == library || library == "" {
			found = append(found, g)
		}
	}

	return found
}

// GetGroup returns a group from the id
func GetGroup(id string) (Group, error) {
	for _, l := range groups {
		if l.Id == id {
			return l, nil
		}
	}

	return Group{}, errors.New("group not found")
}

// GetSagas returns all the sagas of a group
func GetSagas(group string) []Saga {
	var found []Saga
	for _, s := range sagas {
		if s.Group == group || group == "" {
			found = append(found, s)
		}
	}

	return found
}

// GetSaga returns a saga from the id
func GetSaga(id string) (Saga, error) {
	for _, s := range sagas {
		if s.Id == id {
			return s, nil
		}
	}

	return Saga{}, errors.New("saga not found")
}

// GetBooks returns all the books of a saga
func GetBooks(saga string) []Book {
	var found []Book
	for _, b := range books {
		if b.Saga == saga || saga == "" {
			found = append(found, b)
		}
	}

	return found
}

// GetBook returns a book from the isbn
func GetBook(isbn string) (Book, error) {
	for _, b := range books {
		if b.Isbn == isbn {
			return b, nil
		}
	}

	return Book{}, errors.New("book not found")
}
