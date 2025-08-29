package cmd

import (
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
    Use:   "search",
    Short: "Search for entries in the database",
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
