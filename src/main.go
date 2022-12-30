package main

import (
	"fmt"
)

func message(text string, c chan<- string) {
	c <- text
}

func main() {
	c := make(chan string, 2)

	c <- "Mensaje1"
	c <- "Mensaje2"

	/*
		len(v Type) int
		The len built-in function returns the length of v.

		cap(v Type) int
		The cap built-in function returns the capacity of v.
	*/
	fmt.Println(len(c), cap(c))

	close(c)

	//Recorrer un channel
	for message := range c {
		fmt.Println(message)
	}

	//Select
	email1 := make(chan string, 1)
	email2 := make(chan string, 1)

	go message("Mensaje 1", email1)
	go message("Mensaje 2", email2)

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println(m1)
		case m2 := <-email2:
			fmt.Println(m2)
		}
	}
}
