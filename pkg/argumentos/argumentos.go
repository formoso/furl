package argumentos

import (
	"fmt"
	"os"
)

func InitArgs() []string {
	args := validateArgs()
	return args
}

func validateArgs() []string {
	if len(os.Args) <= 1 {
		fmt.Fprintf(os.Stderr, "URL parameter not informed\n")
		os.Exit(1)
	}
	args := removeBodyArgs()
	return args
}

func NoBody() (bool, int) {
	for _, body := range os.Args {
		if body == "body" {
			return true, len(os.Args) - 1
		}
	}
	return false, len(os.Args)
}

func removeBodyArgs() []string {
	var args []string
	for _, s := range os.Args {
		if s != "body" {
			args = append(args, s)
		}
	}
	return args
}
