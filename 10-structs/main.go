package main

import (
	"fmt"
)

type Author struct {
	Name    string
	Country string
}

type Book struct {
	Author Author
	Title  string
	Pages  int
	Year   int
}

func (b Book) String() string {
	return fmt.Sprintf("%s by %s, published in %d", b.Title, b.Author.Name, b.Year)
}

type Library struct {
	Books []Book
}

func (l *Library) AddBook(b Book) {
	l.Books = append(l.Books, b)
}
func (l Library) FindByAuthor(name string) []Book {
	var results []Book
	for _, book := range l.Books {
		if book.Author.Name == name {
			results = append(results, book)
		}
	}
	return results
}

func main() {
	author1 := Author{Name: "George Orwell", Country: "United Kingdom"}
	author2 := Author{Name: "Jane Austen", Country: "United Kingdom"}
	book1 := Book{Author: author1, Title: "1984", Pages: 328, Year: 1949}
	book2 := Book{Author: author1, Title: "Animal Farm", Pages: 112, Year: 1945}
	book3 := Book{Author: author2, Title: "Pride and Prejudice", Pages: 432, Year: 1813}

	library := Library{}
	library.AddBook(book1)
	library.AddBook(book2)
	library.AddBook(book3)

	fmt.Println("All books in the library:")
	for _, book := range library.Books {
		fmt.Println(book)
	}

	fmt.Println("Books by George Orwell:")
	for _, book := range library.FindByAuthor("George Orwell") {
		fmt.Println(book)
	}

}
