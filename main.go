package main

import (
	"github.com/formoso/furl/furl"
	"github.com/formoso/furl/pkg/argumentos"
	"github.com/formoso/furl/pkg/resposta"
)

func main() {
	urls := argumentos.InitArgs()
	ch := furl.Get(urls)
	resposta.RespostaGet(ch)
}
