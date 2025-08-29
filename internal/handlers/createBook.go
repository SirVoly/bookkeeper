package handlers

import (
	"context"
	"fmt"

	"github.com/SirVoly/bookkeeper/internal/database"
	"github.com/google/uuid"
)

func CreateBookHandler(_ []string) error {
	var bookParams database.CreateBookParams
	bookParams.ID = uuid.NewString()

	if err := promptUser(&bookParams); err != nil {
		return fmt.Errorf("error prompting user: %v", err)
	}
	
	book, err := insertBook(bookParams); 
	if err != nil {
		return fmt.Errorf("error inserting book: %v", err)
	}

	fmt.Println("Book added successfully:")
	fmt.Println(book.Print())

	return nil
}

func promptUser(book *database.CreateBookParams) error {
	var err error
	fmt.Println("Fields marked with * can be left empty (will be set to null).")

	book.Title, err = promptString("Book Title")
	if err != nil {
		return fmt.Errorf("error getting book title: %v", err)
	}

	book.Isbn, err = promptInt("ISBN")
	if err != nil {
		return fmt.Errorf("error getting book ISBN: %v", err)
	}

	book.PublicationDate = promptNullableString("Publication Date*")
	book.NumberOfPages = promptNullableInt("Number of Pages*")

	return nil
}

func insertBook(book database.CreateBookParams) (database.Book, error) {
	queries, err := database.InitDatabase()
	if err != nil {
		return database.Book{}, fmt.Errorf("error getting database: %v", err)
	}
	createdBook, err := queries.CreateBook(context.Background(), book)
	if err != nil {
		return database.Book{}, fmt.Errorf("error creating book: %v", err)
	}
	return createdBook, nil
}
