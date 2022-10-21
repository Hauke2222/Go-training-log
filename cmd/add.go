package cmd

import (
	"fmt"
	"strings"
	"time"
	"traininglog/db"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a lift to your log.",
	Long: `Adds a lift to your log. Example:
	Liftname: weight * reps * sets
	Squat: 200 x 5 x 5`,
	Run: func(cmd *cobra.Command, args []string) {
		date := time.Now().Local().String()
		args = append(args, date)
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
