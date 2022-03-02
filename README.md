# homework-2-bakyazi

## build
```shell
$ go build -o bin/library cmd/hw2/main.go
```

## usage
```
Commands:
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
```

### list
```shell
$ ./bin/library list          
1. In Search of Lost Time [ID=1] [Author=Marcel Proust] [ISBN=6364822] [StockCode=5902] [Price=$91] [InStock=30]
2. Ulysses [ID=2] [Author=James Joyce] [ISBN=9366138] [StockCode=8196] [Price=$78] [InStock=160]
3. Don Quixote [ID=3] [Author=Miguel de Cervantes] [ISBN=3390522] [StockCode=1107] [Price=$95] [InStock=34]
4. One Hundred Years of Solitude [ID=4] [Author=Gabriel Garcia Marquez] [ISBN=9897522] [StockCode=2814] [Price=$145] [InStock=156]
5. The Great Gatsby [ID=5] [Author=F. Scott Fitzgerald] [ISBN=1608269] [StockCode=3326] [Price=$148] [InStock=141]
6. Moby Dick [ID=6] [Author=Herman Melville] [ISBN=4989705] [StockCode=3449] [Price=$93] [InStock=150]
7. War and Peace [ID=7] [Author=Leo Tolstoy] [ISBN=7812318] [StockCode=1764] [Price=$16] [InStock=48]
8. Hamlet [ID=8] [Author=William Shakespeare] [ISBN=3512050] [StockCode=2020] [Price=$120] [InStock=190]
.
.
.
51. The Lord of the Rings [ID=51] [Author=J. R. R. Tolkien] [ISBN=4878850] [StockCode=1870] [Price=$132] [InStock=116]
52. The Idiot [ID=52] [Author=Fyodor Dostoyevsky] [ISBN=8524762] [StockCode=1793] [Price=$48] [InStock=67]
```


### search

search (not found)
```shell
$ ./bin/library search fdksjfkjsdhfjksd  
Not found any book with respect to search criteria
```

search by name
```shell
$ ./bin/library search great
1. The Great Gatsby [ID=5] [Author=F. Scott Fitzgerald] [ISBN=9256034] [StockCode=2740] [Price=$110] [InStock=185]
2. Great Expectations [ID=26] [Author=Charles Dickens] [ISBN=2139827] [StockCode=9477] [Price=$147] [InStock=48]
```

search by author
```shell
$ ./bin/library search james
1. Ulysses [ID=2] [Author=James Joyce] [ISBN=1531236] [StockCode=5636] [Price=$35] [InStock=96]
2. A Portrait of the Artist as a Young Man [ID=47] [Author=James Joyce] [ISBN=5976354] [StockCode=8708] [Price=$36] [InStock=79]
```

Note: Since other attributes are random generated in every run. There is no example search by ISBN, Stock Code etc.

### delete

successful delete
```shell
$ ./bin/library delete 10   
book[ID=10] is successfully deleted
```

failed delete
```shell
$ ./bin/library delete 53
error occured during delete operation,
         - not found any book with given parameters
```

Note: Since book list auto generated in each run, I cannot exhibit example output for deleting already deleted book

### buy
successful buy
```shell
$ ./bin/library buy 10 4  
4 of the Book[ID=10] is bought. There are 112 left!
```
failed buy (exceed stock amount)
```shell
$ ./bin/library buy 10 487
error occured during buy operation, 
         - there is not enough stock to sell this book in demanded amount
```

failed buy (try to buy non-exist book)
```shell
$ ./bin/library buy 53 487
error occured during buyOperation operation, 
         - not found any book with given parameters
```

## Unit Tests
```shell
$ go test test/library_test.go
```

You can check unit tests [here](test/library_test.go)