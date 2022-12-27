package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPC pc) ping() {
	fmt.Println(myPC.brand, "Pong")
}

func (myPC *pc) duplicateRAM() {
	myPC.ram *= 2
}

func main() {

	a := 10
	b := &a

	fmt.Println(b)
	fmt.Println(*b)

	// El * hace referencia al valor guardado en la dirección de memoria
	// y el & hace referencia a la dirección de memoria.

	*b = 100
	fmt.Println(a)

	myPC := pc{ram: 16, disk: 200, brand: "MSI"}
	fmt.Println(myPC)

	myPC.ping()
	myPC.duplicateRAM()
	fmt.Println(myPC)

	myPC.duplicateRAM()
	fmt.Println(myPC)
}
