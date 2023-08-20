package testfunctions

import (
	"io/fs"
	"log"
	"os"
)


// Creates a test environment in the current folder
func EnterTestEnvironment() {
    err := os.Mkdir("test", fs.ModePerm)
    if err != nil {
        log.Fatalf("Error creating test environment: %v", err)
    }

    os.Chdir("test")
}


// Exits and deletes the test environment. Assumes the call is made from the test environment
func ExitTestEnvironment() {
    os.Chdir("..")

    err := os.RemoveAll("test")
    if err != nil {
        log.Fatalf("Error deleting test environment: %v", err)
    }
}

