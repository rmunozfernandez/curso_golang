package main

import (
	"fmt"
	"go/src/curso_golang_platzi/src/src/pc"
)

func main() {

	a := 10
	b := &a

	fmt.Println(b)
	fmt.Println(*b)

	// El * hace referencia al valor guardado en la dirección de memoria
	// y el & hace referencia a la dirección de memoria.

	*b = 100
	fmt.Println(a)

	var myPC pc.Pc
	myPC.Brand = "MSI"
	myPC.Disk = 400
	myPC.Ram = 64

	fmt.Println(myPC)

	myPC.Ping()
	myPC.DuplicateRAM()
	fmt.Println(myPC)

	myPC.DuplicateRAM()
	fmt.Println(myPC)
}
