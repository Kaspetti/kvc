package testfunctions

import (
	"io/fs"
	"log"
	"os"
)


func CreateTestEnvironment() {
    err := os.Mkdir("test", fs.ModePerm)
    if err != nil {
        log.Fatalf("Error creating test environment: %v", err)
    }
}


func DeleteTestEnvironment() {
    err := os.RemoveAll("test")
    if err != nil {
        log.Fatalf("Error deleting test environment: %v", err)
    }
}

