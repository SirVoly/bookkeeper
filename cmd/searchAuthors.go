package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var authorSearchParams struct {
	FirstName string
	LastName  string
}

// searchAuthorsCmd represents the authors command
var searchAuthorsCmd = &cobra.Command{
	Use:   "authors",
    Short: "Search for authors",
	Run: func(cmd *cobra.Command, args []string) {
        // Call your search handler here
        fmt.Printf("Searching authors: first-name=%s, last-name=%s\n", authorSearchParams.FirstName, authorSearchParams.LastName)
	},
}

func init() {
	searchCmd.AddCommand(searchAuthorsCmd)

    searchAuthorsCmd.Flags().StringVar(&authorSearchParams.FirstName, "first-name", "", "First name to search for")
    searchAuthorsCmd.Flags().StringVar(&authorSearchParams.LastName, "last-name", "", "Last name to search for")
}
