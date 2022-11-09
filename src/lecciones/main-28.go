// Lección 28: Channels: La forma de organizar las goroutines
package main

import "fmt"

//<- el canal solo es de entrada de datos
func say(texto string, c chan<- string) {
	c <- texto
}
func main_28() {
	c := make(chan string, 1)
	fmt.Println("Hello")
	go say("Bye", c)
	fmt.Println(<-c)
}
