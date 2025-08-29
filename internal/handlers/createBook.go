package handlers

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/SirVoly/bookkeeper/internal/database"
	"github.com/google/uuid"
)

func CreateBookHandler(_ []string) error {
	queries, err := database.InitDatabase()
	if err != nil {
		return fmt.Errorf("error getting database: %v", err)
	}
	book, err := queries.CreateBook(context.Background(), database.CreateBookParams{
		ID:              uuid.NewString(),
		Title:           "Heroes",
		Isbn:            9780718188726,
		PublicationDate: sql.NullString{String: "November 2, 2017", Valid: true},
		NumberOfPages:   sql.NullInt64{Int64: 416, Valid: true},
	})
	if err != nil {
		return fmt.Errorf("error creating book: %v", err)
	}
	fmt.Printf("Book created: %s (%s)\n", book.Title, book.ID)
	return nil
}