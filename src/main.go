package main

import "fmt"

func main() {

	// For condicional
	fmt.Println("For condicionaal")
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	// For while
	fmt.Println("For while")
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	// For forever
	fmt.Println("For forever")
	counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++
	}
}
