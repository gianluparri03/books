package data

// UNCOMMENT these lines to make it compile
// var libraries []Library{}
// var groups []Group{}
// var sagas = []Saga{}
// var books = []Book{}

// Library is the first layer of books grouping 
type Library struct {
	Id   string
	Name string
}

// Group is the second layer of books grouping 
type Group struct {
	Library string

	Id      string
	Name    string
}

// Group is the third layer of books grouping 
type Saga struct {
	Library string
	Group   string

	Id      string
	Name    string
}

// Book is a book
type Book struct {
	Library      string
	Group        string
	Saga         string

	Isbn         string
	Title        string
	Number       int // progressive number inside the saga
	Authors      []string
	Publisher    string
	Lang         string
	Pages        int
	Price        string

	BoughtShop   string
	BoughtDate   string
	StartedDate  string
	FinishedDate string
	Status       string
}
