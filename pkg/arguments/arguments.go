package arguments

import (
	"errors"
	"net/url"
	"os"
)

//GetURLs retorna uma lista de URLs validas
//ou um erro caso exista
func GetURLs() ([]string, error) {
	Args, err := validate()
	return Args, err
}

func validate() ([]string, error) {
	if len(os.Args) <= 1 {
		return nil, errors.New("URL parameter not informed")
	}
	err := isURLsValids(os.Args)
	if err != nil {
		return nil, err
	}
	args := removeBody()
	return args, nil
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

func NoBody() (bool, int) {
	for _, body := range os.Args {
		if body == "body" {
			return true, len(os.Args) - 1
		}
	}
	return false, len(os.Args)
}

func removeBody() []string {
	var args []string
	for _, s := range os.Args {
		if s != "body" {
			args = append(args, s)
		}
	}
	return args
}
