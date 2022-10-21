package cmd

import (
	"fmt"
	"strconv"
	"traininglog/db"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse the argument:", arg)
			} else {
				ids = append(ids, id)
			}
		}
		lifts, err := db.AllLifts()
		if err != nil {
			fmt.Println("Something went wrong:", err)
			return
		}
		for _, id := range ids {
			if id <= 0 || id > len(lifts) {
				fmt.Println("Invalid lift number:", id)
				continue
			}
			lift := lifts[id-1]
			err := db.DeleteLift(lift.Key)
			if err != nil {
				fmt.Printf("Failed to delete \"%d\" . Error: %s\n", id, err)
			} else {
				fmt.Printf("Deleted: \"%s\" ", lift.Value)
			}
		}
		fmt.Println(ids)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
