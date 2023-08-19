package main

import (
	"fmt"
	"kvcadd"
	"kvcinit"
	"os"
)


func main() {
    debugFlag := isDebug()

    if len(os.Args) > 1 {
        switch os.Args[1] {
            case "help":
                writeHelp()
            case "init":
                kvcinit.Init(debugFlag)
            case "add":
                if len(os.Args) < 3 {
                    fmt.Printf("Missing argument [file] in command 'add'\nRun 'kvc help' for valid commands\n")
                }
                kvcadd.Add(os.Args[2:], debugFlag)
        }
    } else {
        writeHelp()
    }
}


func writeHelp() {
    fmt.Println("kvc [option]\nOptions:")
    fmt.Println("\tinit: initializes the current working directory as a kvc repository")
    fmt.Println("\tadd [file1, file2, ...]: adds the given files to the add buffer used when calling 'kvc save'")
    fmt.Println("\tsave [description]: saves the files in the add buffer to a new save in the store with the given description")
    fmt.Println("\tload [identifier]: loads a save from the store into a temporary environment, given an identifier")
    fmt.Println("Flags:")
    fmt.Println("\t-debug: enables debug mode giving more detailed error messages if something should occur") 
}


func isDebug() bool {
    for _, flag := range(os.Args[1:]) {
        if flag == "-debug" || flag == "--debug" {
            return true
        }
    }

    return false
}
