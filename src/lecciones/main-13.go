// Channels: La forma de organizar las goroutines
package main

import "fmt"

func say(texto string, c chan<- string) {
	c <- texto
}

func main_13() {
	c := make(chan string, 1)
	fmt.Println("Hello")

	go say("Bye", c)
	fmt.Println(<-c)
}
