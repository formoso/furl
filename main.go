package main

import (
	"fmt"
	"os"

	"github.com/formoso/furl/furl"
	"github.com/formoso/furl/pkg/arguments"
)

func main() {
	urls, err, wBody := arguments.GetURLs()
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error(), "\n")
		os.Exit(1)
	}
	furl.Get(urls, wBody)
}
