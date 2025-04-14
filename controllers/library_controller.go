package controllers

import (
	"bufio"
	"fmt"
	"library-management/models"
	"library-management/services"
	"os"
)

// For taking input to access services.
func TakeInput() {
	fmt.Println("***Library***")
	fmt.Println("Select one of the following options.")
	fmt.Println("1. Add a new book.")
	fmt.Println("2. Remove an existing book.")
	fmt.Println("3. Borrow a book.")
	fmt.Println("4. Return a book.")
	fmt.Println("5. List all available books.")
	fmt.Println("6. List all borrowed books.")
	fmt.Printf("Enter: ")

	var value int
	fmt.Scanln(&value)

	switch value {
	case 1: // Add book.
		// Get Book's Name
		fmt.Printf("Enter book name: ")
		bkName := bufio.NewScanner(os.Stdin)
		bkName.Scan()
		BookName := bkName.Text()

		// Get Author's Name
		fmt.Printf("Enter author's name: ")
		writer := bufio.NewScanner(os.Stdin)
		writer.Scan()
		author := writer.Text()

		var book = models.Book{ID: services.BookID, Title: BookName, Author: author, Status: "Available"}
		services.BookID++ // Update book ID
		services.AddBook(book)
	case 2: // Remove book.
		// Get Book ID
		var bookID int
		fmt.Printf("Enter book ID: ")
		fmt.Scan(&bookID)

		services.RemoveBook(bookID)
	case 3: // Borrow book.
		var bookID, memberID int

		// Get Book ID
		fmt.Printf("Enter book ID: ")
		fmt.Scan(&bookID)

		// Get Member ID
		fmt.Printf("Enter member ID: ")
		fmt.Scan(&memberID)

		services.BorrowBook(bookID, memberID)
	case 4: // Return book.
		var bookID, memberID int

		// Get Book ID
		fmt.Printf("Enter book ID: ")
		fmt.Scan(&bookID)

		// Get Member ID
		fmt.Printf("Enter member ID: ")
		fmt.Scan(&memberID)

		services.ReturnBook(bookID, memberID)
	case 5:
		services.ListAvailableBooks()
	case 6:
		var memberID int

		// Get Member ID
		fmt.Printf("Enter member ID: ")
		fmt.Scan(&memberID)

		services.ListBorrowedBooks(memberID)
	default:
		err := fmt.Errorf("sorry, you have entered the wrong input")
		fmt.Println(err.Error())
	}
}

/*

// All IDs (Global Scope)
var BookID int
var MemberID int

// METHODS
func AddBook(book *models.Book) {
	fmt.Printf("Enter Book Name: ")
	BookName := bufio.NewScanner(os.Stdin)
	BookName.Scan()
	book.Title = BookName.Text()

	fmt.Printf("Enter Author's Name: ")
	AuthorName := bufio.NewScanner(os.Stdin)
	AuthorName.Scan()
	book.Author = AuthorName.Text()

	book.Status = "Available"

	// Increment the value of book
	BookID += 1

	book.ID = BookID
	var Save services.Library
	Save.Books[BookID] = *book
}

func RemoveBook(book models.Book) {
	fmt.Printf("Enter Book ID: ")
	detail := bufio.NewScanner(os.Stdin)
	detail.Scan()
	delete(book, )
}
*/
