package handlers

import (
	"context"
	"fmt"
	"strings"

	"github.com/SirVoly/bookkeeper/internal/database"
)

func SearchBookHandler(params map[string]interface{}) error {
	db, err := database.InitDatabase()
	if err != nil {
		return fmt.Errorf("error initializing database: %v", err)
	}

	baseQuery := "SELECT * FROM books"
	conditions := []string{}
	args := []interface{}{}

	if title, ok := params["title"]; ok {
		conditions = append(conditions, fmt.Sprintf("title = $%d", len(args)+1))
		args = append(args, title)
	}
	if isbn, ok := params["isbn"]; ok {
		conditions = append(conditions, fmt.Sprintf("isbn = $%d", len(args)+1))
		args = append(args, isbn)
	}
	// Add more fields as needed...

	query := baseQuery
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	rows, err := db.QueryContext(context.Background(), query, args...)
	if err != nil {
		return fmt.Errorf("error executing search: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var book database.Book
		if err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Isbn,
			&book.PublicationDate,
			&book.NumberOfPages,
		); err != nil {
			return fmt.Errorf("error scanning row: %v", err)
		}
		fmt.Println(book.Print())
	}

	return nil
}
