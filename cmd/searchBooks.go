package cmd

import (
	"log"

	"github.com/SirVoly/bookkeeper/internal/handlers"
	"github.com/spf13/cobra"
)

var bookSearchParams struct {
	Title string
	ISBN  int64
}

// searchBooksCmd represents the books command
var searchBooksCmd = &cobra.Command{
	Use:   "books",
	Short: "Search for books",
	Run: func(cmd *cobra.Command, args []string) {
		params := make(map[string]interface{})

		if cmd.Flags().Changed("title") {
			params["title"] = bookSearchParams.Title
		}
		if cmd.Flags().Changed("isbn") {
			params["isbn"] = bookSearchParams.ISBN
		}

		err := handlers.SearchBookHandler(params)
		if err != nil {
			log.Fatalf("error in command create: %v\n", err)
		}
	},
}

func init() {
	searchCmd.AddCommand(searchBooksCmd)

	searchBooksCmd.Flags().StringVar(&bookSearchParams.Title, "title", "", "Filter by title")
	searchBooksCmd.Flags().Int64Var(&bookSearchParams.ISBN, "isbn", 0, "Filter by ISBN")
}
