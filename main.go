package main

import (
	"fmt"
	"os"

	"github.com/formoso/furl/furl"
)

func main() {
	// Este IF verifica o número de argumentos
	if len(os.Args) <= 1 {
		fmt.Fprintf(os.Stderr, "URL parameter not informed\n")
		os.Exit(1)
	}
	for i := 1; i < len(os.Args); i++ {
		url := os.Args[i]
		resp, err := furl.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error fetching URL:%v\n", err)
			os.Exit(1)
		}

		fmt.Printf("%dms %7d %s\n", resp.ElapsedTime, resp.NBytes, resp.Url)
	}
}
