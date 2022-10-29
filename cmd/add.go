package cmd

import (
	"fmt"
	"strings"
	"traininglog/db"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a lift to your log.",
	Long: `Adds a lift to your log. Example:
	Date Liftname: weight * reps * sets
	22-10-2022 Squat: 200 x 5 x 5`,
	Run: func(cmd *cobra.Command, args []string) {
		lift := strings.Join(args, " ")
		_, err := db.CreateLift(lift)
		if err != nil {
			fmt.Println("Something went wrong:", err)
		}
		fmt.Printf("Added \"%s\" to your list.\n", lift)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
