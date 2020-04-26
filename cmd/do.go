package cmd

import (
	"github.com/abhinavdwivedi440/todo/colors"
	"github.com/abhinavdwivedi440/todo/db"
	"github.com/spf13/cobra"
	"strconv"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as complete",
	Long:  "\"do\" command marks a task as complete.\n\nUsage:\n  todo do 1 2 7 4...",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				colors.FgR.Println("Failed to pass the argument:", arg)
			} else {
				ids = append(ids, id)
			}
		}
		tasks, err := db.AllTasks()
		if err != nil {
			colors.FgR.Println("Something went wrong:", err)
			return
		}
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				colors.FgR.Println("Invalid task number:", id)
				continue
			}
			task := tasks[id-1]
			err := db.DeleteTask(task.Key)
			if err != nil {
				colors.FgR.Printf("Failed to mark \"%d\" as completed. Error: %s\n", id, err)
			} else {
				colors.G.Printf("Marked \"%d\" as completed.\n", id)

			}
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
