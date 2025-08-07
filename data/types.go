package data

import "image"

// Library is the first layer of books grouping.
type Library struct {
	Id   string
	Name string
}

// Group is the second layer of books grouping.
type Group struct {
	Library string // Library id

	Id   string
	Name string
}

// Group is the third layer of books grouping.
type Saga struct {
	Group string // Group Id

	Id   string
	Name string
}

// Book is a book.
type Book struct {
	Saga string // Saga Id

	Isbn      string
	Title     string
	Number    string // progressive number inside the saga
	Authors   string
	Publisher string
	Lang      string
	Pages     string
	Price     string

	Status       string // Read, ToBeRead, Reading, Abandoned...
	BoughtShop   string `db:"boughtShop"`   // the shop from which the book has been bought
	BoughtDate   string `db:"boughtDate"`   // the day the book has been bought
	StartedDate  string `db:"startedDate"`  // the day the reading started
	FinishedDate string `db:"finishedDate"` // the day the reading finished

	Thumbnail image.Image // optional
}
