// kvcadd contains functions for adding files to the add buffer in kvc
//
// Author: Kaspar Moberg
package kvcadd

import (
	"fmt"
	"io/fs"
	"os"
)

// Adds the given files to the add buffer. Ignores flags and has a possibility to enable
// debug mode for better error messages
func Add(files []string, debug bool) {
    var filePaths []string

    for _, file := range(files) {
        // Ignore flags
        if file[0] == '-' {
            continue
        }

        fileInfo, err := os.Stat(file)
        if err != nil {
            fmt.Printf("File '%v' does not exist", fileInfo)
        }

        if fileInfo.IsDir() {
            getFilesInDirectory(fileInfo)
        } else {
            filePaths = append(filePaths, file)
        }
    }

    getKVCDir()
}


func getFilesInDirectory(dir fs.FileInfo) []string {
    return make([]string, 0)
}


func getKVCDir() {
    // cwd, _ := os.Getwd()

}
