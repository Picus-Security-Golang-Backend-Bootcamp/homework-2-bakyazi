package main

import (
	"fmt"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-2-bakyazi/internal/library"
	"os"
	"strconv"
)

var bookList *library.BookList

// init loads books slice as a library.BookList
func init() {
	bookList = library.NewBookList(books)
}

func main() {
	args := os.Args

	if len(args) == 1 {
		printUsage()
		return
	}

	command := args[1]

	switch command {
	case "search":
		searchOperation(args)
	case "list":
		listOperation()
	case "delete":
		deleteOperation(args)
	case "buy":
		buyOperation(args)
	default:
		printUsage()
	}

}

// listOperation operates list command
func listOperation() {
	books := bookList.List()
	if len(books) == 0 {
		fmt.Printf("There is no book in library!\n")
		return
	}
	printBooks(books)
	return
}

// searchOperation operates search command
func searchOperation(args []string) {
	if len(args) < 3 {
		printUsage()
		return
	}

	books := bookList.Search(args[2:])
	if len(books) == 0 {
		fmt.Printf("Not found any book with respect to search criteria\n")
		return
	}
	printBooks(books)
	return
}

// buyOperation operates buy command
func buyOperation(args []string) {
	if len(args) != 4 {
		printUsage()
		return
	}

	id, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Printf("given id (%s) is not valid! please enter valid integer! \n", args[2])
		return
	}

	amount, err := strconv.Atoi(args[3])
	if err != nil {
		fmt.Printf("given amount (%s) is not valid! please enter valid integer! \n", args[3])
		return
	}

	err = bookList.Buy(id, amount)
	if err != nil {
		fmt.Printf("error occured during buy operation, \n\t - %s\n", err.Error())
		return
	}
	return
}

// deleteOperation operates delete command
func deleteOperation(args []string) {
	if len(args) != 3 {
		printUsage()
		return
	}

	id, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Printf("given id cannot be converted to integer %s\n", args[2])
		return
	}
	err = bookList.Delete(id)
	if err != nil {
		fmt.Printf("error occured during delete operation,\n\t - %s\n", err.Error())
		return
	}
	fmt.Printf("book[ID=%d] is successfully deleted\n", id)
	return
}

//printUsage prints usage
func printUsage() {
	usage := `Commands: 
- search => to search and list books with specified arguments, arguments are searched in books' name, author, isdn, stockCode and ID attributes
	e.g:
		$ ./bin/library search moby dick
- list => to show list of all books
	e.g:
		$ ./bin/library list
- delete => to delete the book by specified ID
	e.g:
		$ ./bin/library delete 5
- buy => to buy the book specified by the ID in the specified amount, first argument is ID of the book and second argument is the amount desired to be bought
	e.g:
		$ ./bin/library buy 5 10
`
	fmt.Println(usage)
}

// printBooks prints book list prettier
func printBooks(books library.BookList) {
	for i, book := range books {
		fmt.Printf("%d. %s\n", i+1, book.PrettyString())
	}
}

var books = [][]string{
	{"In Search of Lost Time", "Marcel Proust"},
	{"Ulysses", "James Joyce"},
	{"Don Quixote", "Miguel de Cervantes"},
	{"One Hundred Years of Solitude", "Gabriel Garcia Marquez"},
	{"The Great Gatsby", "F. Scott Fitzgerald"},
	{"Moby Dick", "Herman Melville"},
	{"War and Peace", "Leo Tolstoy"},
	{"Hamlet", "William Shakespeare"},
	{"The Odyssey", "Homer"},
	{"Madame Bovary", "Gustave Flaubert"},
	{"The Divine Comedy", "Dante Alighieri"},
	{"Lolita", "Vladimir Nabokov"},
	{"The Brothers Karamazov", "Fyodor Dostoyevsky"},
	{"Crime and Punishment", "Fyodor Dostoyevsky"},
	{"The Catcher in the Rye", "J. D. Salinger"},
	{"Pride and Prejudice", "Jane Austen"},
	{"The Adventures of Huckleberry Finn", "Mark Twain"},
	{"Anna Karenina", "Leo Tolstoy"},
	{"Alice's Adventures in Wonderland", "Lewis Carroll"},
	{"The Iliad", "Homer"},
	{"To the Lighthouse", "Virginia Woolf"},
	{"Catch-22", "Joseph Heller"},
	{"Heart of Darkness", "Joseph Conrad"},
	{"The Sound and the Fury", "William Faulkner"},
	{"Nineteen Eighty Four", "George Orwell"},
	{"Great Expectations", "Charles Dickens"},
	{"The Grapes of Wrath", "John Steinbeck"},
	{"Absalom, Absalom!", "William Faulkner"},
	{"Invisible Man", "Ralph Ellison"},
	{"To Kill a Mockingbird", "Harper Lee"},
	{"The Trial", "Franz Kafka"},
	{"The Red and the Black", "Stendhal"},
	{"Middlemarch", "George Eliot"},
	{"Gulliver's Travels", "Jonathan Swift"},
	{"Beloved", "Toni Morrison"},
	{"Mrs. Dalloway", "Virginia Woolf"},
	{"The Stories of Anton Chekhov", "Anton Chekhov"},
	{"The Stranger", "Albert Camus"},
	{"Jane Eyre", "Charlotte Bronte"},
	{"The Aeneid", "Virgil"},
	{"Collected Fiction", "Jorge Luis Borges"},
	{"The Sun Also Rises", "Ernest Hemingway"},
	{"David Copperfield", "Charles Dickens"},
	{"Tristram Shandy", "Laurence Sterne"},
	{"Leaves of Grass", "Walt Whitman"},
	{"The Magic Mountain", "Thomas Mann"},
	{"A Portrait of the Artist as a Young Man", "James Joyce"},
	{"Midnight's Children", "Salman Rushdie"},
	{"Oedipus the King", "Sophocles"},
	{"Candide", "Voltaire"},
	{"The Lord of the Rings", "J. R. R. Tolkien"},
	{"The Idiot", "Fyodor Dostoyevsky"},
}
