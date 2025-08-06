package data

import "image"

// libraries contains all the libraries.
var libraries []Library

// Library is the first layer of books grouping.
type Library struct {
	Id   string
	Name string
}

// groups contains all the groups.
var groups []Group

// Group is the second layer of books grouping.
type Group struct {
	Library string // Library id

	Id   string
	Name string
}

// sagas contains all the groups.
var sagas []Saga

// Group is the third layer of books grouping.
type Saga struct {
	Library string // Library id
	Group   string // Group Id

	Id   string
	Name string
}

// books contains all the books.
var books []Book

// Book is a book.
type Book struct {
	Library string // Library id
	Group   string // Group Id
	Saga    string // Saga Id

	Isbn      string
	Title     string
	Number    string // progressive number inside the saga
	Authors   []string
	Publisher string
	Lang      string
	Pages     string
	Price     string

	BoughtShop   string // the shop from which the book has been bought
	BoughtDate   string // the day the book has been bought
	StartedDate  string // the day the reading started
	FinishedDate string // the day the reading finished
	Status       string // Read, ToBeRead, Reading, Abandoned...

	Thumbnail image.Image // may be nil
}
