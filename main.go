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
	canal := make(chan furl.Response)
	url := os.Args
	go furl.Get(url, canal)
	for i := 1; i < len(os.Args); i++ {
		resp := <-canal
		if resp.Err != nil {
			fmt.Fprintf(os.Stderr, "Error fetching URL:%v\n", resp.Err)
			os.Exit(1)
		}
		fmt.Printf("%dms %7d %s\n", resp.ElapsedTime, resp.NBytes, resp.Url)
	}

}
