package commands

import (
    "fmt"

	"goman/core"
)

type GoInfo struct {
    Version   string
    Processes string
}

type PossibleFlag struct {
    h bool
    a bool
    g bool
}

func Details(commandArgs []string, version string) error {

	flags := PossibleFlag{h: false, a: false, g: false}

	if len(commandArgs) > 0 {
		switch commandArgs[0] {
		case "-h":
			flags.h = true
		case "-a":
			flags.a = true
		case "-g":
			flags.g = true
		default:
			fmt.Println("Invalid flag please refer to the help menu goman details -h")
			return nil
		}
	}

	if flags.h {
		fmt.Println("")
		fmt.Println("Usage: goman details [flag]")
		fmt.Println("Flags:")
		fmt.Println("  -h    Display help")
		fmt.Println("  -a    Display all details")
		fmt.Println("  -g    Display go details")
		fmt.Println("")
		return nil
	}


	fmt.Println("")
    fmt.Println("       Details of GoMan       ")
	fmt.Println("")
    fmt.Println("  --------------------------  ")
    fmt.Println("  GoMan is a simple tool to manage Go versions on your system.")
    fmt.Println("  Goman Version:", version)
    fmt.Println("  --------------------------  ")
	fmt.Println("")

    
    // Add more details as needed

	if flags.g || flags.a {
		fmt.Println("")
		info := core.GetGoInfo()
		fmt.Println("       Details of Go       ")
		fmt.Println("")
		fmt.Println("  -----------------------  ")
		fmt.Println("  Go version:", info.Version)
		fmt.Println("  Go processes:")
		fmt.Println("")
		fmt.Println(info.Processes)
		fmt.Println("  -----------------------  ")
		fmt.Println("")
	}

    return nil
}