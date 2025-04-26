/*
* Created on 26 Feb 2024
* @author Sai Sumanth
 */

package main

import (
	"os"
	"strconv"
	"strings"
)

// checks whether a file exists at the given path
// returns true if file exists
func DoesFileExists(filePath string) bool {
	if _, err := os.Stat(filePath); err == nil {
		return true
	}
	return false
}

// converts given comma separated numbers string to []int and returns it
func CommaSeparatedStringToArray(str string) []int {
	res := []int{}
	for _, v := range strings.Split(str, ",") {
		// parse string to integer
		num, _ := strconv.Atoi(v)
		res = append(res, num)
	}
	return res
}
