package cmd

import (
	"log"
	"fmt"

	"github.com/SirVoly/bookkeeper/internal/handlers"
	"github.com/spf13/cobra"
)

var createType string

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new entry into the database",
	Long: `Create a new entry into the database.
The command defaults to adding a book,
but this can by changed with the --type (-t) flag.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch createType {
		case "book", "books":
			err := handlers.CreateBookHandler(args)
			if err != nil {
				log.Fatalf("error in command create: %v", err)
			}
			return
		case "author", "authors":
			fmt.Println("Create author not yet implemented.")
			return
		default:
			log.Fatalf("unknown type: %s", createType)
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringVarP(&createType, "type", "t", "book", "Type to create")
}
