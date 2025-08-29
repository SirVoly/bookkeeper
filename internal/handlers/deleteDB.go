package handlers

import (
	"context"
	"fmt"

	"github.com/SirVoly/bookkeeper/internal/database"
)

func DeleteDatabase(_ []string) error {
	queries, err := database.InitQueries()
	if err != nil {
		return fmt.Errorf("error getting database: %v", err)
	}

	if err := queries.DeleteBooks(context.Background()); err != nil {
		return fmt.Errorf("error deleting books: %v", err)
	}
	return nil
}
