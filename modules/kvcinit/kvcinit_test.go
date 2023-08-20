package kvcinit

import (
	"os"
	"testfunctions"
	"testing"
)


func TestInit(t *testing.T) { // Setup test environment
    testfunctions.EnterTestEnvironment()
    // Make sure to delete the test environment even if the tests fail
    defer testfunctions.ExitTestEnvironment()

    err := Init(true)
    if err != nil {
        t.Fatalf("Init failed with error: %v", err)
    }

    // Check if .kvc exists
    if kvcDir, err := os.Stat(".kvc"); err != nil {
        t.Fatalf("Init failed to create .kvc folder: %v\n", err)
    } else if !kvcDir.IsDir() {
        t.Fatalf("Init successfully created .kvc but it is not a directory\n")
    }

    // Check if .kvc/store exists
    if storeDir, err := os.Stat(".kvc/store"); err != nil {
        t.Fatalf("Init failed to create .kvc/store folder: %v\n", err)
    } else if !storeDir.IsDir() {
        t.Fatalf("Init successfully created .kvc/store but it is not a directory\n")
    } 

    // Check if main branch has been created
    if mainDir, err := os.Stat(".kvc/store/main"); err != nil {
        t.Fatalf("Init failed to create ./kvc/store/main folder: %v\n", err)
    } else if !mainDir.IsDir() {
        t.Fatalf("Init successfully created .kvc/store/main but it is not a directory\n")
    }

    // Check if .kvc/buf.kvc exists
    _, err = os.Stat(".kvc/buf.kvc")
    if err != nil {
        t.Fatalf("Init failed to create .kvc/buf.kvc file: %v\n", err)
    }
}


// Tests if the Init function returns an error if the folder 
// Init is called from already has a .kvc repository
func TestInitExists(t *testing.T) {
    testfunctions.EnterTestEnvironment()
    defer testfunctions.ExitTestEnvironment()

    if err := Init(true); err != nil {
        t.Fatalf("Init failed on first Init call with error: %v\n", err)
    } 
    
    if err := Init(true); err == nil {
        t.Fatal("Init allowed to be called twice without error\n")
    }
}


// Tests if the Init function returns an error if any parent folder 
// of the folder where Init is called from already has a .kvc repository
func TestInitExistsRecursive(t *testing.T) {
    testfunctions.EnterTestEnvironment()
    defer testfunctions.ExitTestEnvironment()

    if err := Init(true); err != nil {
        t.Fatalf("Error while calling Init: %v", err)
    }

    testfunctions.EnterTestEnvironment()
    defer testfunctions.ExitTestEnvironment()

    if err := Init(true); err == nil {
        t.Fatalf("Init ran successfully even though parent was a kvc repository")
    }
}


