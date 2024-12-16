package main

import (
    "flag"
    "fmt"
    "os"

    "goman/commands"
    "goman/utils"
)

const version = "0.0.1"

func main() {
    helpFlag := flag.Bool("h", false, "Display help")
    versionFlag := flag.Bool("v", false, "Display version")
	contactFlag := flag.Bool("c", false, "Contact us")

    flag.Parse()

    // Handle flags
    if *helpFlag {
        utils.PrintUsage()
        os.Exit(0)
    }
    if *versionFlag {
        fmt.Println("Goman version", version)
        os.Exit(0)
    }

	if *contactFlag {
		utils.ContactDeveloper()
		os.Exit(0)
	}

    // Check for commands
    args := flag.Args()
    if len(args) < 1 {
        utils.PrintUsage()
        os.Exit(1)
    }

    command := args[0]
    commandArgs := args[1:]

    // Execute command
    var err error
    switch command {
    case "install":
        err = commands.Install(commandArgs)
    case "use":
        err = commands.Use(commandArgs)
    case "list":
        err = commands.List()
    case "remove":
        err = commands.Remove(commandArgs)
    case "purge":
        err = commands.Purge()
    case "setup":
        err = commands.Setup()
	case "details":
		err = commands.Details(commandArgs, version)
    default:
        utils.PrintUsage()
        os.Exit(1)
    }

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        os.Exit(1)
    }
}