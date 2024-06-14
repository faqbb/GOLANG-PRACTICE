package main

import "fmt"

type BookStatus string

const (
	Available BookStatus = "disponible"
	Borrowed  BookStatus = "prestado"
)

type Book struct {
	Title  string
	Author string
	Genre  string
	Status BookStatus
}

type Library struct {
	books []Book
}

func (l *Library) FindBooks(query string) []*Book {
	foundBooks := []*Book{}
	for _, book := range l.books {
		if book.Title == query || book.Author == query {
			foundBooks = append(foundBooks, &book)
		}
	}
	return foundBooks
}

func CreateBook(title string, author string, genre string, status BookStatus) Book {
	newBook := Book{
		Title:  title,
		Author: author,
		Genre:  genre,
		Status: status,
	}
	return newBook
}

func (l *Library) AddBook(book Book) {
	l.books = append(l.books, book)
}

func (l *Library) UpdateStatus(title string, status BookStatus) {
	for _, book := range l.books {
		if title == book.Title {
			book.Status = status
		}
	}
}

func (l *Library) DeleteBook(title string) {
	for i, book := range l.books {
		if title == book.Title {
			l.books = append(l.books[:i], l.books[i+1:]...)
		}
	}
}

func (l *Library) displayBooks() {
	for _, book := range l.books {
		fmt.Printf("Titulo: %s\n Autor: %s\n Genero: %s\n Estado: %s", book.Title, book.Author, book.Genre, book.Status)
	}
}

func main() {
	library := &Library{}
	newBook := CreateBook("la granja", "jose", "terror", "disponible")
	library.AddBook(newBook)
	library.displayBooks()

}
