package project

import (
	"os"
)

func ProjectHandler(args []string) error {
	if len(args) != 1 {
		Help()
		os.Exit(1)
	}
	Init(args)
	return nil
}