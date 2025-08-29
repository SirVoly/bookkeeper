package handlers

import (
	"context"
	"fmt"

	"github.com/SirVoly/bookkeeper/internal/database"
)

func ListBooksHandler(_ []string) error {
	queries, err := database.InitDatabase()
	if err != nil {
		return fmt.Errorf("error getting database: %v", err)
	}
	books,err := queries.ListBooks(context.Background())
	if err != nil {
		return fmt.Errorf("error listing books: %v", err)
	}
	for _, book := range books {
		fmt.Printf("%s\n\n", book.Print())
	}
	return nil
}