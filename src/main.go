package main

import (
	"fmt"
	"sync"
	"time"
)

func say(text string, wg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Println(text)
}

func main() {
	//WaitGroup - Acumular un conjunto de gorutines
	var wg sync.WaitGroup

	//say("Hello")
	fmt.Println("Hello")
	wg.Add(1)

	go say("world", &wg)

	wg.Wait()

	//Función anónima
	go func(text string) {
		fmt.Println(text)
	}("Adios")

	time.Sleep(time.Second * 1) //no es eficiente
}
