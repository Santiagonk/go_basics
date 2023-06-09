package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	inicio := time.Now()
	canal := make(chan string)
	servidores := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}

	// for _, servidor := range servidores {
	// 	go revisarServidor(servidor, canal)
	// }

	// for i := 0; i < len(servidores); i++ {
	// 	fmt.Println(<-canal)
	// }

	i := 0

	for {
		if i > 2 {
			break
		}
		for _, servidor := range servidores {
			go revisarServidor(servidor, canal)
		}
		time.Sleep(1 * time.Second)
		fmt.Println(<-canal)
		i++
	}

	tiempoPaso := time.Since(inicio)
	fmt.Printf("Tiempo de ejecución %s\n", tiempoPaso)
}

func revisarServidor(servidor string, canal chan string) {
	_, err := http.Get(servidor)
	if err != nil {
		// fmt.Println(servidor, "no esta disponible")
		canal <- servidor + " no se encuentra disponible"
	} else {
		// fmt.Printf("servidor funcionando %s\n", servidor)
		canal <- servidor + " se encuentra disponible"
	}
}
