package cmd

import (
	"github.com/abhinavdwivedi440/todo/colors"
	"github.com/abhinavdwivedi440/todo/db"
	"github.com/spf13/cobra"
	"os"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all of your tasks.",
	Long:  "\"list\" command lists all of your tasks.\n\nUsage:\n  todo list",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		if err != nil {
			colors.FgR.Println("Something went wrong:", err)
			os.Exit(1)
		}
		if len(tasks) == 0 {
			colors.G.Println("You have no tasks to complete! Why not take a vacation ?")
			return
		}
		colors.G.Println("You have the following tasks:")
		for i, task := range tasks {
			colors.Y.Printf("%d. %s\n", i+1, task.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
