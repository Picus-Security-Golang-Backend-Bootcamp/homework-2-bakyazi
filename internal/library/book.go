package library

import (
	"fmt"
	"sync"
)

// Deletable interface Book struct implements it
type Deletable interface {
	Delete() error
}

// Book struct is data structure to store information about books
type Book struct {
	sync.RWMutex
	ID        int
	Name      string
	Author    string
	StockCode string
	ISBN      string
	PageCount int
	Price     int
	Amount    int
	IsDeleted bool
}

// NewBook create new book with name & author
// fill attributes with map
// NOTE purpose of this function is to create not random book object
// for unit tests
func NewBook(book map[string]interface{}, id int) *Book {
	return &Book{
		ID:        id,
		Name:      book["name"].(string),
		Author:    book["author"].(string),
		StockCode: book["stock_code"].(string),
		ISBN:      book["isbn"].(string),
		PageCount: book["page_count"].(int),
		Price:     book["price"].(int),
		Amount:    book["amount"].(int),
		IsDeleted: false,
	}
}

// NewBookRandomFill create new book with name & author
// randomly generated StockCode, ISBN, PageCount, Price, Amount
func NewBookRandomFill(book []string, id int) *Book {
	return &Book{
		ID:        id,
		Name:      book[0],
		Author:    book[1],
		StockCode: generateRandomInteger(1000, 9999).String(),
		ISBN:      generateRandomInteger(1000000, 9999999).String(),
		PageCount: int(generateRandomInteger(100, 1000)),
		Price:     int(generateRandomInteger(5, 150)),
		Amount:    int(generateRandomInteger(20, 200)),
		IsDeleted: false,
	}
}

// Delete determines that book is deleted
func (b *Book) Delete() error {
	if b.IsDeleted {
		return ErrBookAlreadyDeleted
	}
	b.IsDeleted = true
	return nil
}

// DecreaseAmount decreases amount of book
func (b *Book) DecreaseAmount(amount int) (int, error) {
	// to handle case of concurrent clients try to buy this book
	b.Lock()
	defer b.Unlock()

	if b.Amount < amount {
		return -1, ErrBookOutOfStock
	}

	b.Amount -= amount
	return b.Amount, nil
}

// String makes Book object flat string
func (b *Book) String() string {
	return fmt.Sprintf("%d %s %s %s %s %d %d %d",
		b.ID,
		b.Name,
		b.Author,
		b.StockCode,
		b.ISBN,
		b.PageCount,
		b.Price,
		b.Amount)
}

// PrettyString creates human friendly string
func (b *Book) PrettyString() string {
	return fmt.Sprintf("%s [ID=%d] [Author=%s] [ISBN=%s] [StockCode=%s] [Price=$%d] [InStock=%d]",
		b.Name,
		b.ID,
		b.Author,
		b.ISBN,
		b.StockCode,
		b.Price,
		b.Amount)
}
