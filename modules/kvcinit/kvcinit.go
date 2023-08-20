// kvcinit contains the functions required to initialize a kvc repository.
//
// Author: Kaspar Moberg
package kvcinit

import (
	"fmt"
	"os"
    "errors"
    "path/filepath"
)


// Initializes the current working directory as a kvc repository by 
// generating the necessary directories and files for kvc to function
func Init(verbose bool) error {
    cwd, err := os.Getwd()
    if err != nil {
        fmt.Println("An unexpected error has occured")
        if verbose {
            fmt.Println(err)
        }
        
        return err
    }

    if isKVC, path := isKVCRec(cwd); isKVC {
        fmt.Printf("Directory '%v' is a sub directory of a kvc repository rooted at: '%v'\n", cwd, path)
        return errors.New(fmt.Sprintf("Directory '%v' is a sub directory of a kvc repository rooted at: '%v'\n", cwd, path))
    }

    err = generateKVCDir(cwd)
    if err != nil {
        if errors.Is(err, os.ErrExist){
            fmt.Printf("Directory '%v' is already a kvc repository\n", cwd)
        } else {
            fmt.Println("An unexpected error has occured")   
        }

        if verbose {
            fmt.Println(err)
        }
        return err
    }
    fmt.Printf("Initialized kvc repository at: %v\n", cwd)

    return nil
}


func generateKVCDir(cwd string) error {
    // Generates the main kvc folder. Any directory that has a .kvc folder is 
    // recognized as a kvc repository by kvc
    if err := os.Mkdir(".kvc", os.ModePerm); err != nil {
        return err
    }
    
    // Creates the buffer file used to temporarily store files set to be saved by 
    // calling 'kvc add'
    bufFile, err := os.Create(".kvc/buf.kvc")
    if err != nil {
        return err
    }
    defer bufFile.Close()

    // Creates the store for kvc, which is a folder for storing all saves
    if err := os.Mkdir(".kvc/store", os.ModePerm); err != nil {
        return err
    }

    if err := os.Mkdir(".kvc/store/main", os.ModePerm); err != nil {
        return err
    }

    return nil
}


// Checks if any of the parents of the current folder is a .kvc folder
func isKVCRec(wd string) (bool, string) {
    parent := filepath.Dir(wd)
    if parent == wd {
        return false, ""
    } 

    path := filepath.Join(parent, ".kvc")
    if kvcDir, _ := os.Stat(path); kvcDir != nil {
        return true, parent 
    }

    return isKVCRec(parent)
}
