package resposta

import (
	"fmt"
	"os"

	"github.com/formoso/furl/furl"
)

func RespostaGet(ch *chan furl.Response, urls []string) {
	for i := 1; i < len(urls); i++ {
		resp := <-*ch
		if resp.Err != nil {
			fmt.Fprintf(os.Stderr, "Error fetching URL:%v\n", resp.Err)
			os.Exit(1)
		}
		imprimir(resp)
	}
}

func imprimir(resp furl.Response) {
	fmt.Printf("%dms %7d %s\n", resp.ElapsedTime, resp.NBytes, resp.Url)
}
