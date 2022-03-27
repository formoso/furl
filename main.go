package main

import (
	"github.com/formoso/furl/furl"
	"github.com/formoso/furl/pkg/argumentos"
	"github.com/formoso/furl/pkg/resposta"
)

func main() {
	// Este IF verifica o número de argumentos
	urls := argumentos.InitArgs()
	ch := furl.Get(urls)
	resposta.RespostaGet(ch, urls)
}
