package vm

// IM (Instruction Mapping) struct
type IM struct {
	O Opcode
	S int8
}

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

func New() *VM {
	return &VM{}
}
