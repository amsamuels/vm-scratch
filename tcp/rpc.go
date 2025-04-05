package tcp

import "net"

/*
1. RPC represents a remote procedure call exchanged between nodes.
2. It contains the method name, its arguments, and the sender's address.
3. Method - The operation to execute (e.g., "MOV", "PUSH", "HLT")
4. Args - Optional arguments for the operation
5. From - Address of the sender; populated internally
*/

type RPC struct {
	Method string
	Args   []int16
	From   net.Addr `json:"-"`
}
