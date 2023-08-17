package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

/*
This command Adds a task to a list of tasks. the task is added to a specific user
Tasks are ordered in a acending order. if a task is deleted or completed, another
task takes its place
*/
func init() {
	rootCmd.AddCommand(addCmd)
}

/*
method adds a task with the task ID specified.
The task that is added has a name, and an ID, the ID is assigned automatically
the name, or description is assigned by the user.

TODO: check if the name is empty. a minimum amount of chars needs to be set.
*/
func addTask(args []string) {
	fileName := "tasks.json"

	userIdx, err := GetUser(fileName, args[0])
	if err != nil {
		log.Panic("User is not found")
	}

	//checking for password info
	users := GetUsers(fileName) // Getting all users
	if users.Users[userIdx].Password != "" {
		pass := promptPassword()
		err1 := checkPassword(users.Users[userIdx].Password, pass)
		if err1 != nil {
			log.Panic("Fatal: password does not match")
		}
	}

	tempTask := Task{Id: len(users.Users[userIdx].Tasks) + 1, Note: args[1]}
	users.Users[userIdx].Tasks = append(users.Users[userIdx].Tasks, tempTask)

	data, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	err = os.WriteFile(fileName, data, 0644)
	if err != nil {
		log.Fatal(err)
	}

}

var addCmd = &cobra.Command{
	Use:   "add [UserName] [\"Task_in_qoutes\"]",
	Short: "Add a new task to your tasks list",
	Long: `Adds a task to your task list. The task is added 
	with a task number. The task number is like an ID number for the task
	that can be used to delete or change the task.
	If a has an assigned password, he will be prompted to enter it.`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		addTask(args)
	},
}
