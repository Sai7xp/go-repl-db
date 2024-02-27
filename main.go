/*
* Created on 26 Feb 2024
* @author Sai Sumanth
 */

package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
// REPL - Read, Eval, Print, Loop

# A REPL shell based db built using Go Lang

It uses a single file to store everything, the goal is to have persistent data. The database can
store integer arrays and perform some operations on them.

Resources Used -
 1. https://www.digitalocean.com/community/tutorials/what-is-repl
    2.
*/
func main() {

	/// start a infinite loop
	fmt.Println("Command: `wkn new` to create new db\nCommand: `wkn` to start the REPL ")

	// label [InfiniteLoop]
InfiniteLoop:
	for {
		scaner := bufio.NewScanner(os.Stdin)
		fmt.Print(">")
		scaner.Scan()
		userInput := scaner.Text()

		/// process user input
		switch userInput {
		case "wkn new":
			// to create new .wkn file
			checkAndCreateNewDb()
		case "wkn":
			// to start the REPL
			if DoesFileExists(".wkn") {
				// 🎊 user can interact with REPL db now
				startREPL()
				break InfiniteLoop
			} else {
				fmt.Println("Create new db using `wkn new` command")
			}

		case "exit":
			os.Exit(3)
		default:
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

func startREPL() {
	fmt.Println("Welcome to REPL Shell")
	for {
		fmt.Print("wkn>")
		scaner := bufio.NewScanner(os.Stdin)
		fmt.Print(">")
		scaner.Scan()
		userInput := scaner.Text()
		handleReplCommand(userInput)

	}
}

func handleReplCommand(command string) {
	switch command {
	case "exit":
		os.Exit(3)
	default:
		fmt.Println("OOPS! Command not found")
	}
}
