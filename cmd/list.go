/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/SirVoly/bookkeeper/internal/handlers"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the entries from a table in the database",
	Long: `List the entries from a table in the database.
The command defaults to listing all books,
but this can be changed with the --type (-t) flag.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := handlers.ListBooksHandler(args)
		if err != nil {
			log.Fatalf("error in command create: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
