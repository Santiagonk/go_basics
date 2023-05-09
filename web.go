package main

import (
	"fmt"
	"io"
	"net/http"
)

type escritorioWeb struct{}

func (escritorioWeb) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	return len(p), nil
}

func main() {
	respuesta, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println(err)
	}
	e := escritorioWeb{}
	io.Copy(e, respuesta.Body)
}
