/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// finishCmd represents the finish command
var finishCmd = &cobra.Command{
	Use:   "finish [UserName] [Task-ID]",
	Short: "Finishes/Deletes a Task by ID number.",
	Long: `Finish or Deletes a specific task using its ID number.
a UserName must be provided, along with a password (if required) to complete a task.`,
	Args: cobra.ExactArgs(2),

	Run: func(cmd *cobra.Command, args []string) {
		finish(args)
	},
}

/*
method finished or deletes a task from user's task list. the first arg is the
user name, the second arg is the task id number.

TODO: fix some bugs
 1. the way variable num is being converted, this should be forced to be an int by the commandline args
 2. redundent use of code. can be mergered with add.go code
 3. fix the order of the tasks such that they are always updated (1,2,3..)
*/
func finish(args []string) {
	fileName := "tasks.json"

	userIdx, err := GetUser(fileName, args[0])
	if err != nil {
		log.Panic("User is not found")
	}

	users := GetUsers(fileName) // Getting all users
	if users.Users[userIdx].Password != "" {
		pass := promptPassword()
		err1 := checkPassword(users.Users[userIdx].Password, pass)
		if err1 != nil {
			log.Panic("Fatal: password does not match")
		}
	}

	num, _ := strconv.Atoi(args[1])

	tasks := users.Users[userIdx].Tasks
	for idx, task := range tasks {
		if task.Id == num {
			users.Users[userIdx].Tasks = append(users.Users[userIdx].Tasks[:idx],
				users.Users[userIdx].Tasks[idx+1:]...)
			break
		}
	}

	data, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	err = os.WriteFile(fileName, data, 0644)
	if err != nil {
		log.Fatal(err)
	}

}

func init() {
	rootCmd.AddCommand(finishCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// finishCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// finishCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
