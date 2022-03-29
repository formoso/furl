package resposta

import (
	"fmt"
	"os"

	"github.com/formoso/furl/furl"
	"github.com/formoso/furl/pkg/argumentos"
)

func RespostaGet(ch *chan furl.Response) {
	body, leng := argumentos.NoBody()
	for i := 1; i < leng; i++ {
		resp := <-*ch
		if resp.Err != nil {
			fmt.Fprintf(os.Stderr, "Error fetching URL:%v\n", resp.Err)
			os.Exit(1)
		}
		if body {
			imprimirBody(resp)
		} else {
			imprimirNoBody(resp)
		}
	}
}

func imprimirNoBody(resp furl.Response) {
	fmt.Printf("%dms %7d %s\n", resp.ElapsedTime, resp.NBytes, resp.Url)
}

func imprimirBody(resp furl.Response) {
	fmt.Printf("%dms %7d %s \n Body:%s\n", resp.ElapsedTime, resp.NBytes, resp.Url, resp.Body)
}
