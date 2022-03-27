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
	return os.Args
}
