package main

import "fmt"

func main() {
	//declaración de variables
	helloMessage := "Hello"
	worldMessage := "world"

	//Println
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	//Printf
	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene más de %d cursos\n", nombre, cursos)
	/*
		%s el valor a traer es un string
		%d el valor a traer es un número
		%v no se sabe el tipo de dato que va a traer, la buena practica es especificar el tipo de dato
	*/

	//Sprintf
	message := fmt.Sprintf("%s tiene más de %d cursos\n", nombre, cursos)
	fmt.Print(message)

	//Tipo datos
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T\n", cursos)
}
