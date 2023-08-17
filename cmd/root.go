/*
Copyright © 2023 Laith Abou Saleh
*/
package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"syscall"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/term"
)

type Users struct {
	Users []User `json:"users"`
}

// defines a single user and their tasks.
type User struct {
	Name     string `json:"name"`
	UserName string `json:"username"`
	Password string `json:"password,omitempty"`
	Tasks    []Task `json:"task,omitempty"`
}

type Task struct {
	Id   int
	Note string
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "TaskManager",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

/*
check if a user exisit in a list of users
*/
func IsUser(users Users, userName string) (bool, error) {
	for _, user := range users.Users {
		// fmt.Println(user.UserName)
		if user.UserName == userName {
			return true, errors.New("User Exist")
		}
	}
	return false, nil
}

/*
Helper function, used to get Users that exist inside JSON file.
*/
func GetUsers(fileName string) Users {

	//check if json file exist, if not create one
	if _, err := os.Stat(fileName); err != nil {
		os.OpenFile(fileName, os.O_CREATE, 0644)
	}

	// reading json file
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	// turning bytes into json format
	var users Users
	json.Unmarshal(data, &users)
	return users
}

/*
Get a single user. the fuction takes in the User name. if the
user is not found, a message is shown. function returns user and error if error exisit
*/
func GetUser(fileName, userName string) (int, error) {

	users := GetUsers(fileName)
	for i, user := range users.Users {
		if user.UserName == userName {
			return i, nil
		}
	}
	return -1, errors.New("User Not Found")
}

/*
Takes a password as a string, and returns a string representation of the hash
*/
func hashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		log.Fatal(err)
	}
	return string(hash)
}

// prompt the user to enter a password
// this is used instead of a command line arugment because it
// doen't show the password on the screen when user inputs it
func promptPassword() string {
	fmt.Println("Enter your password here (Enter for no password): ")
	bytePassword, _ := term.ReadPassword(syscall.Stdin)
	return string(bytePassword)
}

// checks if the password entered is correct
func checkPassword(hashed string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.TaskManager.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
