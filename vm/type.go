package vm

// Reg represents an unsigned 16-bit register (equivalent to unsigned short in C)
type Reg uint16

// Registers represents the CPU registers
type Registers struct {
	AX    Reg
	BX    Reg
	CX    Reg
	DX    Reg
	SP    Reg // Stack Pointer
	IP    Reg // Instruction Pointer
	Flags Reg // Flags register
}

// Opcode represents instruction opcodes
type Opcode uint8

// Enum-like constants for Opcodes
const (
	Nop  Opcode = 0x01
	Hlt  Opcode = 0x02
	Mov  Opcode = 0x08 // 0x08 - 0x0f
	Ste  Opcode = 0x10
	Cle  Opcode = 0x11
	Stg  Opcode = 0x12
	Clg  Opcode = 0x13
	Sth  Opcode = 0x14
	Clh  Opcode = 0x15
	Stl  Opcode = 0x16
	Cll  Opcode = 0x17
	Push Opcode = 0x1a
	Pop  Opcode = 0x1b
)

// Args represents instruction arguments (0-2 bytes)
type Args [2]int16

// Instruction represents a VM instruction
type Instruction struct {
	O Opcode
	A Args
}
