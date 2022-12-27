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

func (myPc Pc) String() string {
	return fmt.Sprintf("Tengo un pc marca %s con un disco de %d GB y %d GB de RAM", myPc.Brand, myPc.Disk, myPc.Ram)
}
