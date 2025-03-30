package vm

import "fmt"

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

// Create an instruction with zero arguments
func I0(op Opcode) Instruction {
	return Instruction{O: op}
}

// Create an instruction with one argument
func I1(op Opcode, a1 int16) Instruction {
	return Instruction{O: op, A: [2]int16{a1, 0}}
}

// Create an instruction with two arguments
func I2(op Opcode, a1, a2 int16) Instruction {
	return Instruction{O: op, A: [2]int16{a1, a2}}
}

// ExampleProgram loads a program into VM memory
func (vm *VM) ExampleProgram(instr ...Instruction) {
	for _, inst := range instr {
		vm.B = vm.copy(inst)

		if inst.O == Hlt {
			break // Stop early
		}
	}
}

// copy writes an instruction into memory and returns the updated breakpoint
func (vm *VM) copy(inst Instruction) int16 {
	opc := inst.O
	size := GetInstructionSize(opc)

	// Validate instruction size
	if size == 0 {
		fmt.Println("Invalid instruction size")
		return vm.B // Skip invalid instruction/opcodes
	}

	// Store the opcode/instruction in memory
	vm.M[vm.B] = int8(opc)
	vm.B++

	// Store instruction arguments
	for i := 0; i < int(size)-1 && i < len(inst.A); i++ {
		vm.M[vm.B] = int8(inst.A[i])
		vm.B++
	}

	// Stop copying once we hit `Hlt`
	if opc == Hlt {
		return vm.B
	}

	return vm.B
}

func GetInstructionSize(op Opcode) uint8 {
	// Opcode-to-instruction size mapping
	var instrMap = map[Opcode]uint8{
		Nop:  1,
		Hlt:  1,
		Mov:  3,
		0x09: 3, 0x0A: 3, 0x0B: 3, 0x0C: 3, // MOV variants
		0x0D: 3, 0x0E: 3, 0x0F: 3,
		Ste:  1,
		Stg:  1,
		Stl:  1,
		Sth:  1,
		Cle:  1,
		Clg:  1,
		Cll:  1,
		Clh:  1,
		Push: 3,
		Pop:  3,
	}

	return instrMap[op] // returns the size of an opcode
}
