package vm

// Memory represents a fixed-size memory array (0-65536 bytes)
type Memory [65536]int8

// Program represents a program (in this case, a slice of Instructions)
type Program []Instruction

// CPU represents the central processing unit
type CPU struct {
	R Registers
}

// VM represents the virtual machine
type VM struct {
	C CPU
	M Memory
	B int16 // Break line
}

// New initializes a new VM with default values
func NewVirtualmachine() *VM {
	vm := &VM{}
	vm.C.R.SP = 0xFFFF // Stack pointer at the top
	vm.C.R.IP = 0x0000 // Start execution at address 0
	for i := range vm.M {
		vm.M[i] = 0 // Zero out memory
	}
	return vm
}
