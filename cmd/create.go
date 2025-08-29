/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/SirVoly/bookkeeper/internal/handlers"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new entry into the database",
	Long: `Create a new entry into the database.
The command defaults to adding a book,
but this can by changed with the --type (-t) flag.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := handlers.CreateBookHandler(args)
		if err != nil {
			log.Fatalf("error in command create: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
