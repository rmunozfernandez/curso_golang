package pc

import "fmt"

type Pc struct {
	Ram   int
	Disk  int
	Brand string
}

func (myPC Pc) Ping() {
	fmt.Println("Marca", myPC.Brand)
}

func (myPC *Pc) DuplicateRAM() {
	myPC.Ram *= 2
}
