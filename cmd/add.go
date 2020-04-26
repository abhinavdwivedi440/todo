package cmd

import (
	"fmt"
	"github.com/abhinavdwivedi440/todo/db"
	"github.com/spf13/cobra"
	"strings"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list.",
	Long: "\"add\" command  adds a task to your task list.\n\nUsage:\n  todo add \"Your task\"",
	Run: func (cmd *cobra.Command, args []string) {
			task := strings.Join(args, " ")
			_, err := db.CreateTask(task)
			if err != nil {
				return
			}
			fmt.Printf("Added \"%s\" to your task list.\n", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}