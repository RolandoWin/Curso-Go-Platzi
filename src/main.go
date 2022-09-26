// Range, Close y Select en channels
package main

import "fmt"

func message(texto string, c chan string) {
	c <- texto
}

func main() {
	//El canal maneja dos datos a la vez
	c := make(chan string, 2)
	c <- "Mensaje-01"
	c <- "Mensaje-02"

	//len: cantidad de canales utilizados
	//cap: capacidad maxima de canales
	fmt.Println(len(c), cap(c))

	//Range y Close
	close(c)

	for message := range c {
		fmt.Println(message)
	}

	//Select
	email1 := make(chan string)
	email2 := make(chan string)

	go message("mensaje1", email1)
	go message("mensaje2", email2)

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email recibido de email1", m1)
		case m2 := <-email2:
			fmt.Println("Email recibido de email2", m2)
		}
	}
}
