package library

import "errors"

var (
	ErrBookAlreadyDeleted = errors.New("book is already deleted")
	ErrNotFoundBook       = errors.New("not found any book with given parameters")
	ErrBookOutOfStock     = errors.New("there is not enough stock to sell this book in demanded amount")
)
