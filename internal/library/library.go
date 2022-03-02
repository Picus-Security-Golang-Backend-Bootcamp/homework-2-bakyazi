package library

import (
	"fmt"
	"strings"
)

// BookList type alias for Book slice
type BookList []*Book

// NewBookList creates BookList with book name and author other fields filled randomly
func NewBookList(books [][]string) *BookList {
	var bookList BookList = make(BookList, len(books))
	for i, book := range books {
		bookList[i] = NewBookRandomFill(book, i+1)
	}
	return &bookList
}

// Search return list of books contains given text pattern
func (b *BookList) Search(texts []string) []*Book {
	var result []*Book

	searchPattern := strings.ToLower(strings.Join(texts, " "))

	for _, book := range *b {
		if book.IsDeleted {
			// pass if it is deleted
			continue
		}

		if strings.Contains(strings.ToLower(book.String()), searchPattern) {
			result = append(result, book)
		}
	}
	return result
}

// List lists all books not deleted
func (b *BookList) List() []*Book {
	var result []*Book
	for _, book := range *b {
		if book.IsDeleted {
			// pass if it is deleted
			continue
		}
		result = append(result, book)
	}
	return result
}

// Delete delete book with given id
func (b *BookList) Delete(id int) error {
	book, err := b.findByID(id)
	if err != nil {
		return err
	}
	return book.Delete()
}

// Buy buys book having given ID and decrease amount of it by given amount
func (b *BookList) Buy(id, amount int) error {
	book, err := b.findByID(id)
	if err != nil {
		return err
	}
	remaining, err := book.DecreaseAmount(amount)
	if err != nil {
		return err
	}
	fmt.Printf("%d of the Book[ID=%d] is bought. There are %d left!\n", amount, id, remaining)
	return nil
}

// findByID returns book with given ID
func (b *BookList) findByID(id int) (*Book, error) {
	for _, book := range *b {
		if book.ID == id {
			return book, nil
		}
	}
	return nil, ErrNotFoundBook
}
