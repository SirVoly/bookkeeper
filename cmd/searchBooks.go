package cmd

import (
	"fmt"

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
        // Call your search handler here
        fmt.Printf("Searching books: title=%s, isbn=%d\n", bookSearchParams.Title, bookSearchParams.ISBN)
	},
}

func init() {
	searchCmd.AddCommand(searchBooksCmd)

    searchBooksCmd.Flags().StringVar(&bookSearchParams.Title, "title", "", "Title to search for")
    searchBooksCmd.Flags().Int64Var(&bookSearchParams.ISBN, "isbn", -1, "ISBN to search for")
}
