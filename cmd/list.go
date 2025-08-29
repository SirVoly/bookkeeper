package cmd

import (
	"fmt"
	"log"

	"github.com/SirVoly/bookkeeper/internal/handlers"
	"github.com/spf13/cobra"
)

var listType string

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the entries from a table in the database",
	Long: `List the entries from a table in the database.
The command defaults to listing all books,
but this can be changed with the --type (-t) flag.`,
	Run: func(cmd *cobra.Command, args []string) {
		var handler handlers.Handler
		switch listType {
		case "book", "books":
			handler = handlers.ListBooksHandler
		case "author", "authors":
			fmt.Println("List authors not yet implemented.")
			return
		default:
			log.Fatalf("unknown type: %s\n", listType)
		}
		err := handler(args)
		if err != nil {
			log.Fatalf("error in command create: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().StringVarP(&listType, "type", "t", "book", "Table to list")
}
