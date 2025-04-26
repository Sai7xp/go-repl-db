/*
* Created on 26 Feb 2024
* @author Sai Sumanth
 */

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
// REPL - Read, Eval, Print, Loop

# A REPL shell based db built using Go Lang.
Takes single user input and reads it, evaluates it, print the result and then loop back to wait for the next input.

It uses a single file to store everything, the goal is to have persistent data. The database can
store integer arrays as of now and perform some operations on them.

Resources Used -
 1. https://www.digitalocean.com/community/tutorials/what-is-repl
*/
func main() {

	/// start a infinite loop
	fmt.Println("Command: `wkn new` to create new db\nCommand: `wkn` to start the REPL\nCommand: `wkn --db-path ./<path_to_file>` to check whether file is valid or not")

	for {
		scaner := bufio.NewScanner(os.Stdin)
		fmt.Print(">")
		scaner.Scan()
		userInput := scaner.Text()

		/// process user input
		inputSlice := strings.Fields(userInput)
		if userInput == "wkn new" {
			// to create new .wkn file
			checkAndCreateNewDb()
		} else if userInput == "wkn" {
			// to start the REPL
			if DoesFileExists(".wkn") {
				// 🎊 user can interact with REPL db now
				startREPL()
				break
			} else {
				fmt.Println("Create new db using `wkn new` command")
			}
		} else if len(inputSlice) > 2 && inputSlice[1] == "--db-path" {
			if DoesFileExists(inputSlice[2]) {
				fmt.Println("File is Valid. Starting the REPL")
				startREPL()
				break
			} else {
				log.Fatal("File is Corrupted")
			}
		} else if userInput == "exit" {
			os.Exit(3)
		} else {
			fmt.Println("OOPS! Command not found")
		}
	}
}

func checkAndCreateNewDb() {
	if DoesFileExists(".wkn") {
		/// db file already created
		fmt.Println("Db file already present. You can start the REPL shell using command `wkn` ")
	} else {
		// Create new .wkn file to store data
		_, err := os.Create(".wkn")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("New DB file .wkn created successfully!")
	}
}

// enter REPL shell
// user can use DB commands after entering this db shell
func startREPL() {
	fmt.Println("=====================================")
	fmt.Println("        Welcome to REPL Shell        ")
	fmt.Println("=====================================")
	fmt.Println("  new <array_name> 1,2,3,4   - Create new array")
	fmt.Println("  show <array_name>          - Display an existing array")
	fmt.Println("  add <array_name> 1,2,3     - Add elements to an existing array")
	fmt.Println("  del <array_name>           - to delete the array")
	fmt.Println("-------------------------------------")
	fmt.Println("Note: Data will be loaded from .wkn file if available.")
	// load existing from .wkn file
	existingData, _ := os.ReadFile(".wkn")
	// convert []byte data to map
	json.Unmarshal(existingData, &database)

	for {
		fmt.Print("wkn>")
		scaner := bufio.NewScanner(os.Stdin)
		scaner.Scan()
		replCommand := scaner.Text()
		if len(replCommand) == 0 {
			fmt.Println("OOPS! Command not found")
		} else {
			processUserCommand(replCommand)
		}
	}
}

// proces REPL commands
func processUserCommand(command string) {
	fields := strings.Fields(command)
	fieldsLength := len(fields)

	switch fields[0] {
	case "exit":
		fmt.Println("Bye!")
		os.Exit(3)
	case "new":
		if fieldsLength > 2 {
			parsedArray := CommaSeparatedStringToArray(fields[2])
			CreateNewArray(fields[1], parsedArray)
		} else if fieldsLength > 1 {
			/// array with no elemnts
			CreateNewArray(fields[1], []int{})
		} else {
			fmt.Printf("new is for creating a new array\n Usage:\n\tnew <array_name> 1,2,3,4\n")
		}
	case "show":
		if fieldsLength > 1 {
			GetArray(fields[1])
		} else {
			fmt.Printf("show is for fetching array details\n Usage:\n\tshow <array_name>\n")
		}
	case "del":
		if fieldsLength == 2 {
			DeleteArray(fields[1])
		} else {
			fmt.Printf("del is for deleting array from database\n Usage:\n\tdel <array_name>\n")
		}
	case "add":
		if fieldsLength == 3 {
			parsedArray := CommaSeparatedStringToArray(fields[2])
			AppendElementsToArray(fields[1], parsedArray)
		} else {
			fmt.Println("add <array_name> 1,2,3")
		}

	case "merge":
		///TODO: merge two arrays functionality
	default:
		fmt.Printf("\"%s\" is not a supported operatoin\n", fields[0])
	}
}
