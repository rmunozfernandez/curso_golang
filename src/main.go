package main

import "fmt"

// El channel va a recibir el dato cuando tiene la flecha al lado derecho <-
func say(text string, c chan<- string) {
	//el channel recibe el dato
	c <- text
}

func main() {
	c := make(chan string, 1) //creación del channel

	fmt.Println("Hello")

	//se crea la gorutine
	go say("world", c)

	//obtenemos el dato de salida del channel
	fmt.Println(<-c)
}
