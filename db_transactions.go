/*
* Created on 27 Feb 2024
* @author Sai Sumanth
 */
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

// / Map of arrays will be stored in this database, array name will be the key
var database = make(map[string][]int)
var mx sync.Mutex

// writes data to .wkn file
func WriteDataToFile() error {
	mx.Lock()
	defer mx.Unlock()

	// convert map data to []byte data
	data, err := json.Marshal(database)
	if err != nil {
		return err
	}

	/// write byte data to file
	err = os.WriteFile(".wkn", data, 0644)
	if err != nil {
		return err
	}
	return err
}

// trigger when user enters `new array_name` command
func CreateNewArray(arrayName string, arr []int) {
	defer WriteDataToFile()
	mx.Lock()
	defer mx.Unlock()

	database[arrayName] = arr
	arrLength := len(arr)
	if arrLength == 0 {
		fmt.Println("CREATED")
	} else {
		fmt.Printf("CREATED (%d)\n", len(arr))
	}

}

// Deletes an array with given arrName from database
func DeleteArray(arrName string) {
	defer WriteDataToFile()
	mx.Lock()
	defer mx.Unlock()
	if _, doesExists := database[arrName]; doesExists {
		/// remove key from map
		delete(database, arrName)
		fmt.Println("DELETED")
	} else {
		errMessage := fmt.Sprintf("Error: \"%s\" does not exist", arrName)
		fmt.Println(errMessage)
	}
}

// gets triggered when user enters `show array_name` command
func GetArray(arrayName string) bool {
	mx.Lock()
	defer mx.Unlock()

	if arr, doesExists := database[arrayName]; doesExists {
		fmt.Println(arr)
		return true
	} else {
		errMessage := fmt.Sprintf("Error: \"%s\" does not exist", arrayName)
		fmt.Println(errMessage)
		return false
	}
}

func AppendElementsToArray(arrName string, newValues []int) {
	defer WriteDataToFile()
	mx.Lock()

	arr, doesExists := database[arrName]
	if !doesExists {
		fmt.Println("Specified array doesn't exist")
		return
	}
	arr = append(arr, newValues...)
	database[arrName] = arr
	fmt.Println("Values added to an existing array successfully")
	mx.Unlock()
	GetArray(arrName)
}
