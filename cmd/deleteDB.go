package cmd

import (
	"fmt"
	"log"

	"github.com/SirVoly/bookkeeper/internal/handlers"
	"github.com/spf13/cobra"
)

var forced bool

// deleteDBCmd represents the deleteDB command
var deleteDBCmd = &cobra.Command{
	Use:   "deleteDB",
	Short: "Deletes the full database. Should only be used for debugging / testing",
	Long: `deleteDB clears the full database.
This command cannot be reversed and should be used with caution!
The command is provided for debugging and testing purposes only.`,
	Run: func(cmd *cobra.Command, args []string) {
		if forced {
			if err := handlers.DeleteDatabase(args); err != nil {
				log.Fatalf("error deleting database: %v\n", err)
			}
			fmt.Println("deleted database")
		} else {
			log.Fatalf("deletion not forced\n")
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteDBCmd)
	deleteDBCmd.Flags().BoolVarP(&forced, "force", "f", false, "Force deletion")
}
