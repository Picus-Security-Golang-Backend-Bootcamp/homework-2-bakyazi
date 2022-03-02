package test

import (
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-2-bakyazi/internal/library"
	"testing"
)

var books = []map[string]interface{}{
	{
		"name":       "Moby Dick",
		"author":     "Herman Melville",
		"stock_code": "123123",
		"isbn":       "111111",
		"page_count": 100,
		"price":      10,
		"amount":     55,
	},
	{
		"name":       "War and Peace",
		"author":     "Leo Tolstoy",
		"stock_code": "123124",
		"isbn":       "222222",
		"page_count": 200,
		"price":      20,
		"amount":     65,
	},
	{
		"name":       "Hamlet",
		"author":     "William Shakespeare",
		"stock_code": "123125",
		"isbn":       "333333",
		"page_count": 300,
		"price":      30,
		"amount":     75,
	},
}

func TestList(t *testing.T) {
	bookList := loadLibrary()
	result := bookList.List()

	if len(result) != 3 {
		t.Fail()
	}
	var book *library.Book
	book = result[0]
	// first book check price
	if book.ID != 1 || book.Price != 10 {
		t.Fail()
	}
	book = result[1]
	// second book check ISBN
	if book.ID != 2 || book.ISBN != "222222" {
		t.Fail()
	}
	book = result[2]
	// third book check stock code
	if book.ID != 3 || book.StockCode != "123125" {
		t.Fail()
	}
}

func TestSearch(t *testing.T) {
	bookList := loadLibrary()
	var book *library.Book

	// test search with non-matching text
	result := bookList.Search([]string{"sdjkfsdhjkfhjksd"})

	if len(result) != 0 {
		t.Fail()
	}

	// test search with name
	result = bookList.Search([]string{"moby"})

	if len(result) != 1 {
		t.Fail()
	}
	book = result[0]
	if book.ID != 1 || book.Author != "Herman Melville" {
		t.Fail()
	}

	// search with author
	result = bookList.Search([]string{"tolstoy"})

	if len(result) != 1 {
		t.Fail()
	}
	book = result[0]
	if book.ID != 2 || book.Name != "War and Peace" {
		t.Fail()
	}

	// test search with isbn
	result = bookList.Search([]string{"333333"})

	if len(result) != 1 {
		t.Fail()
	}
	book = result[0]
	if book.ID != 3 || book.Name != "Hamlet" {
		t.Fail()
	}

}

func TestBuy(t *testing.T) {
	bookList := loadLibrary()
	// buy 5 of first book
	err := bookList.Buy(1, 5)
	if err != nil {
		t.Fail()
	}

	// search first book and check if its amount has been decreased or not
	result := bookList.Search([]string{"moby"})
	book := result[0]
	if book.Amount != 50 {
		t.Fail()
	}

	// buy second book and expect to get library.ErrBookOutOfStock
	// since there are only 65 second book
	err = bookList.Buy(2, 70)
	if err != library.ErrBookOutOfStock {
		t.Fail()
	}

	// buy non exist book and expect to get library.ErrNotFoundBook
	// since there are only 65 second book
	err = bookList.Buy(4, 1)
	if err != library.ErrNotFoundBook {
		t.Fail()
	}

}

func TestDelete(t *testing.T) {
	bookList := loadLibrary()

	// try to delete a book not exist
	err := bookList.Delete(4)
	if err != library.ErrNotFoundBook {
		t.Fail()
	}

	// delete third book
	err = bookList.Delete(3)
	if err != nil {
		t.Fail()
	}

	// try to delete third book again and expect ErrBookAlreadyDeleted
	err = bookList.Delete(3)
	if err != library.ErrBookAlreadyDeleted {
		t.Fail()
	}

	// try to get third book by search operation and expect to find noting since it is deleted
	result := bookList.Search([]string{"hamlet"})
	if len(result) != 0 {
		t.Fail()
	}

}

// loadLibrary create library.BookList with global variable books
func loadLibrary() *library.BookList {
	var list library.BookList

	for i, b := range books {
		book := library.NewBook(b, i+1)
		list = append(list, book)
	}
	return &list
}
