// Leccion 27: Primer contacto con las Goroutines
package main

import (
	"fmt"
	"sync"
)

func say(texto string, wg *sync.WaitGroup) {
	defer wg.Done() //Libera a la goroutine del waitgroup
	fmt.Println(texto)
}

func main_27() {
	var wg sync.WaitGroup
	fmt.Println("Hello")
	wg.Add(1)

	go say("world", &wg)

	//Creamos una función anonima
	go func(texto string) {
		fmt.Println(texto)
	}("Adios!!!")

	wg.Wait() //Le decimos al main que espere, hasta que todas las rutinas del waitgroup finalicen
	//time.Sleep(time.Second * 1)
}
