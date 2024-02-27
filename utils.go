/*
* Created on 26 Feb 2024
* @author Sai Sumanth
 */

package main

import (
	"os"
)

// checks whether a file exists at the given path
// returns true if file exists
func DoesFileExists(filePath string) bool {
	if _, err := os.Stat(filePath); err == nil {
		return true
	}
	return false
}
