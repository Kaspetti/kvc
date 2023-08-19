// kvcinit contains the functions required to initialize a kvc repository.
//
// Author: Kaspar Moberg
package kvcinit

import (
	"fmt"
	"os"
    "errors"
)


// Initializes the current working directory as a kvc repository by 
// generating the necessary directories and files for kvc to function
func Init(debug bool) {
    cwd, err := os.Getwd()
    if err != nil {
        fmt.Println("An unexpected error has occured")
        if debug {
            fmt.Println(err)
        }
    }

    err = generateKVCDir(cwd)
    if err != nil {
        if errors.Is(err, os.ErrExist){
            fmt.Printf("Directory '%v' is already a kvc repository\n", cwd)
        } else {
            fmt.Println("An unexpected error has occured")   
        }

        if debug {
            fmt.Println(err)
        }
        return
    }
    fmt.Printf("Initialized kvc repository at: %v\n", cwd)
}


func generateKVCDir(cwd string) error {
    // Generates the main kvc folder. Any directory that has a .kvc folder is 
    // recognized as a kvc repository by kvc
    err := os.Mkdir(".kvc", os.ModePerm)
    if err != nil {
        return err
    }
    
    // Creates the buffer file used to temporarily store files set to be saved by 
    // calling 'kvc add'
    _, err = os.Create(".kvc/buf.kvc")
    if err != nil {
        return err
    }

    // Creates the store for kvc, which is a folder for storing all saves
    return os.Mkdir(".kvc/store", os.ModePerm)
}
