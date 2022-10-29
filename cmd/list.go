package cmd

import (
	"fmt"
	"os"
	"traininglog/db"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Displays your saved lifts.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
		lifts, err := db.AllLifts()
		if err != nil {
			fmt.Println("Something went wrong:", err)
			os.Exit(1)
		}
		if len(lifts) == 0 {
			fmt.Println("You have not saved any lifts.")
			return
		}
		fmt.Println("You have saved the following lifts:")
		for i, lift := range lifts {
			fmt.Printf("%d. %s\n", i+1, lift.Value)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
