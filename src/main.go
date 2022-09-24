// Primer contacto con las gourutines
package main

import (
	"fmt"
	"sync"
	"time"
)

func say(texto string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(texto)
}

func main() {
	var wg sync.WaitGroup

	fmt.Println("Hello")
	wg.Add(1)

	//agregamos concurrencia con "go"
	go say("World", &wg)
	//Espera a que terminen todas las gourutines
	wg.Wait()

	go func(texto string) {
		fmt.Println(texto)
	}("Adios")

	time.Sleep(time.Second * 1)
}
