package mypackage

import "fmt"

// CarPublic con acceso público
type CarPublic struct {
	Brand string
	Year  int
}

// PrintMessage
func PrintMessage(texto string) {
	fmt.Println(texto)
}
