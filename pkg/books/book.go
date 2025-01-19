package books

type Book struct {
	Name   string
	Author string
}

func NewBook(name, author string) *Book {
	return &Book{
		Name:   name,
		Author: author,
	}
}
