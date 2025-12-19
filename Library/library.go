package Library

import "fmt"

type Library struct {
	Books map[string]Book
}

func (l *Library) AddBook() {
	title := readLine("Please enter the Title of a book")
	author := readLine("Enter the Author of this book")

	book := Book{
		ID:         len(l.Books) + 1,
		Title:      title,
		Author:     author,
		IsBorrowed: false,
	}
	l.Books[book.Title] = book
	fmt.Printf("You have succesfully added the book %s", book.Title)
}

func (l *Library) BorrowBook() {
	title := readLine("Please enter a title of a book you want to borrow")
	book, ok := l.Books[title]
	if !ok {
		fmt.Println("There is no books with this title")
		return
	}

	if book.IsBorrowed {
		fmt.Println("This book is already borrowed")
		return
	}

	book.IsBorrowed = true

	l.Books[title] = book
	fmt.Printf("You have borrowed book %s \n", title)
}

func (l *Library) ReturnBook() {
	if len(l.Books) <= 0 {
		fmt.Println("There is no books to return")
		return
	}
	title := readLine("Please enter a title of a book you want to return")
	book, ok := l.Books[title]
	if !ok {
		fmt.Println("There is no book with this title")
		return
	}
	if book.IsBorrowed == false {
		fmt.Println("You cannot return a book that not borrowed")
		return
	}
	book.IsBorrowed = false
	l.Books[title] = book

	fmt.Printf("You have returned the book %s", title)
}

func (l *Library) ListAvailableBooks() {
	noAvailableBooks := "There is no available books"
	if len(l.Books) == 0 {
		fmt.Println(noAvailableBooks)
	}
	count := 0
	for k, v := range l.Books {
		if v.IsBorrowed == false {
			fmt.Println("-------")
			fmt.Println(k)
			fmt.Println("-------")
			count++
		}
	}

	if count == 0 {
		fmt.Println(noAvailableBooks)
	}
}

func NewLibrary() *Library {
	return &Library{
		Books: make(map[string]Book),
	}
}
