package services

import (
	"fmt"
	"library-management/models"
)

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
}

type Library struct {
	Books   map[int]models.Book
	Members map[int]models.Member
}

// In memory data
var libraryDetail = Library{
	Books: map[int]models.Book{
		1: {ID: 1, Title: "Respectable Sins", Author: "Jerry Bridges", Status: "Borrowed"},
		2: {ID: 2, Title: "Dancers at the Gate of Death", Author: "Dr D. K. Olukoya", Status: "Available"},
		3: {ID: 3, Title: "The Prayer Ministry of the Church", Author: "Watchman Nee", Status: "Available"},
		4: {ID: 4, Title: "The Honorable Imposter", Author: "Gibert Morris", Status: "Available"},
		5: {ID: 5, Title: "Atomic Habits", Author: "James Clear", Status: "Available"},
	},
	Members: map[int]models.Member{
		1: {ID: 1, Name: "Ade", BorrowedBooks: []models.Book{{ID: 1, Title: "Respectable Sins", Author: "Jerry Bridges", Status: "Borrowed"}}},
		2: {ID: 2, Name: "Bayo", BorrowedBooks: []models.Book{}},
		3: {ID: 3, Name: "Ola", BorrowedBooks: []models.Book{}},
		4: {ID: 4, Name: "Oluwa", BorrowedBooks: []models.Book{}},
		5: {ID: 5, Name: "Tunde", BorrowedBooks: []models.Book{}},
	},
}

// Global Scope
var BookID int = 6

// METHODS

// Adds book to the library
func AddBook(book models.Book) {
	libraryDetail.Books[book.ID] = book

	// List out all books
	ListAvailableBooks()
}

// Removes books from the library.
func RemoveBook(bookID int) {
	delete(libraryDetail.Books, bookID)

	// List out all books
	ListAvailableBooks()
}

// Check if book exists, is available to be borrows,
// and lends it if it is available.
func BorrowBook(bookID int, memberID int) error {

	book, bookExists := libraryDetail.Books[bookID]
	member, memberExists := libraryDetail.Members[memberID]

	// Find Book and check if it is available.
	if !bookExists {
		return fmt.Errorf("sorry, this book does not exist")
	}

	// Check if the book is available.
	if book.Status != "Available" {
		return fmt.Errorf("sorry, this book is not available at the moment")
	}

	// Check if member exists.
	if !memberExists {
		return fmt.Errorf("sorry, this member does not exist")
	}

	// Update book status.
	book.Status = "Borrowed"
	libraryDetail.Books[bookID] = book

	// Add book to list of books borrowed by the member.
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	libraryDetail.Members[memberID] = member

	// List out all books
	ListAvailableBooks()

	// List out all books borrowed by this particular member.
	ListBorrowedBooks(memberID)
	return nil

}

// Returns borrowed books
func ReturnBook(bookID int, memberID int) error {

	// Check if the member is an existing member.
	for _, member := range libraryDetail.Members {
		if member.ID == memberID {

			// Check if this member actually borrowed this book.
			for idx, book := range member.BorrowedBooks {
				if book.ID == bookID {
					member.BorrowedBooks = append(member.BorrowedBooks[:idx], member.BorrowedBooks[idx+1:]...)
					book.Status = "Available"
					libraryDetail.Books[bookID] = book

					// List out all books
					ListAvailableBooks()
					return nil
				}
			}
			return fmt.Errorf("this book was not borrowed by this user")
		}
	}
	return fmt.Errorf("this user does not exist")
}

// Gives a list of available books.
func ListAvailableBooks() []models.Book {
	var availableBooks []models.Book

	for _, book := range libraryDetail.Books {
		if book.Status == "Available" {
			availableBooks = append(availableBooks, book)
		}
	}
	fmt.Println(availableBooks)
	return availableBooks
}

func ListBorrowedBooks(memberID int) []models.Book {
	for _, member := range libraryDetail.Members {
		if member.ID == memberID {
			fmt.Println(member.BorrowedBooks)
			return member.BorrowedBooks
		}
	}
	return []models.Book{}
}
