package kvcinit

import (
	"errors"
	"os"
	"testfunctions"
	"testing"
)


func TestInit(t *testing.T) {
    // Setup test environment
    testfunctions.CreateTestEnvironment()
    os.Chdir("test")

    Init(true)

    // Check if .kvc exists
    kvcDir, err := os.Stat(".kvc")
    if err != nil && errors.Is(err, os.ErrNotExist) {
        t.Fatalf("Init failed to create .kvc folder: %v\n", err)
    }
    // Check if .kvc is a directory. Shouldn't happen, but to make sure
    if !kvcDir.IsDir() {
        t.Fatalf("Init successfully created .kvc but it is not a directory\n")
    }

    // Check if .kvc/store exists
    storeDir, err := os.Stat(".kvc/store") 
    if err != nil {
        t.Fatalf("Init failed to create .kvc/store folder: %v\n", err)
    } 
    // Check if .kvc/store is a directory
    if !storeDir.IsDir() {
        t.Fatalf("Init successully created .kvc but it is not a directory\n")
    }

    // Check if .kvc/buf.kvc exists
    _, err = os.Stat(".kvc/buf.kvc")
    if err != nil {
        t.Fatalf("Init failed to create .kvc/buf.kvc file: %v\n", err)
    }

    os.Chdir("..") 
    testfunctions.DeleteTestEnvironment()
}
