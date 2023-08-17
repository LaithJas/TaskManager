package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createUserCmd)
}

// /*
// Takes a password as a string, and returns a string representation of the hash
// */
// func hashPassword(password string) string {
// 	hash, err := bcrypt.GenerateFromPassword([]byte(password), 8)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return string(hash)
// }

// // prompt the user to enter a password
// // this is used instead of a command line arugment because it
// // doen't show the password on the screen when user inputs it
// func promptPassword() string {
// 	fmt.Println("Enter your password here (Enter for no password): ")
// 	bytePassword, _ := term.ReadPassword(syscall.Stdin)
// 	return string(bytePassword)
// }

/*
Adds a user to the list of users. collects Name, Username, and password.
The password is immediately hashed and send to the JSON file as a hash.
The data is then Marshaled and sent to JSON file.
*/
func create(usrInfo []string) {
	fileName := "tasks.json"
	jsonData := GetUsers(fileName)

	// checks if User name already exist
	_, err := IsUser(jsonData, usrInfo[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	usr := User{}
	usr.Name = usrInfo[0]
	usr.UserName = usrInfo[1]

	// hides password input
	password := promptPassword()
	if len(password) > 0 {
		usr.Password = hashPassword(password)
	}

	// marchalling and sending the data to json file
	jsonData.Users = append(jsonData.Users, usr)
	data, err := json.MarshalIndent(jsonData, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	err = os.WriteFile(fileName, data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

var createUserCmd = &cobra.Command{
	Use:   "createUser [Name] [UserName]",
	Short: "Create a new user",
	Long: `Creates a new user.
	The new user can have to include a Name and Username. They can choose to omit the password.
	If The user include a password, they need to provide it when they view or change their tasks`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		create(args)
	},
}
