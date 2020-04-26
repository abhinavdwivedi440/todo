package cmd

import (
	"github.com/abhinavdwivedi440/todo/colors"
	"github.com/abhinavdwivedi440/todo/db"
	"github.com/spf13/cobra"
	"strings"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list.",
	Long:  "\"add\" command  adds a task to your task list.\n\nUsage:\n  todo add \"Your task\"\nHere quotes are optional\n  todo add Your task\n",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			colors.FgR.Println("You have not provided any task to add in your task list\nPlease provide a task to add")
			return
		}
		task := strings.Join(args, " ")
		tasks, err := db.AllTasks()
		if err != nil {
			colors.FgR.Println("Something went wrong.")
			return
		}
		for _, t := range tasks {
			if t.Value == task {
				colors.FgR.Println("You already have this task in your list")
				return
			}
		}
		_, err = db.CreateTask(task)
		if err != nil {
			colors.FgR.Println("Something went wrong:", err)
			return
		}
		colors.M.Print("Added ")
		colors.Y.Printf("\"%s\"", task)
		colors.M.Print(" to your task list.\n")
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
