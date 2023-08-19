package main

import (
	"fmt"
	"kvcadd"
	"kvcinit"
	"os"
)


func main() {
    verboseFlag := isVerbose()

    if len(os.Args) > 1 {
        switch os.Args[1] {
            case "help":
                writeHelp()
            case "init":
                kvcinit.Init(verboseFlag)
            case "add":
                if len(os.Args) < 3 {
                    fmt.Printf("Missing argument [file] in command 'add'\nRun 'kvc help' for valid commands\n")
                }
                kvcadd.Add(os.Args[2:], verboseFlag)
        }
    } else {
        writeHelp()
    }
}


func writeHelp() {
    fmt.Print("usage: kvc [-v | --verbose] <command> [<args>]\n\n")
    fmt.Print("initialize a kvc repository\n\tkvc init\n\n")
    fmt.Print("add or remove files from the save buffer\n")
    fmt.Print("\tkvc add [<file1> <file2> ... <fileN>]\n")
    fmt.Print("\tkvc rem [<file1> <file2> ... <fileN>]\n\n")
    fmt.Print("save or load snapshots from the store\n")
    fmt.Print("\tkvc save [<description>]\n")
    fmt.Print("\tkvc load [<identifier>]\n\n")
}


func isVerbose() bool {
    for _, flag := range(os.Args[1:]) {
        if flag == "-v" || flag == "--verbose" {
            return true
        }
    }

    return false
}
