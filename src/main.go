// Lección 29: Range, Close y Select en channels
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

	//len: cantidad de canales utilizados (goroutines)
	//cap: capacidad maxima de canales
	fmt.Println("Cantidad Goroutines: ", len(c), "Capacidad Máxima: ", cap(c))

	//Range: Nos ayuda hacer el recorrido de una lista
	//Close: Cierra los canales cuando ya se deja de usar
	close(c)

	for message := range c {
		fmt.Println(message)
	}

	// //Select
	email_01 := make(chan string)
	email_02 := make(chan string)

	go message("Mensaje-01", email_01)
	go message("Mensaje-02", email_02)

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email_01:
			fmt.Println("Email recibido de email01", m1)
		case m2 := <-email_02:
			fmt.Println("Email recibido de email02", m2)
		}
	}
}
