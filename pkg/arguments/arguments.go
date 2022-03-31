package arguments

import (
	"errors"
	"net/url"
	"os"
)

//GetURLs retorna uma lista de URLs validas
//ou um erro caso exista
func GetURLs() ([]string, error, bool) {
	Args, err, wBody := validate()
	return Args, err, wBody
}

func validate() ([]string, error, bool) {
	if len(os.Args) <= 1 {
		return nil, errors.New("URL parameter not informed"), false
	}
	err := isURLsValids(os.Args)
	if err != nil {
		return nil, err, false
	}
	wBody := noBody()
	args := fixParameter()
	return args, nil, wBody
}

func isURLsValids(urls []string) error {
	for _, u := range urls {
		if u != "body" {
			_, err := url.ParseRequestURI(u)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func noBody() bool {
	for _, body := range os.Args {
		if body == "body" {
			return true
		}
	}
	return false
}

func fixParameter() []string {
	var args []string
	for i, s := range os.Args {
		if i != 0 {
			if s != "body" {
				args = append(args, s)
			}
		}
	}
	return args
}
