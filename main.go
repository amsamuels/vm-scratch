package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"vm-scratch/tcp"
	"vm-scratch/vm"
)

func main() {
	vm := vm.NewVirtualmachine()

	tcpOpts := tcp.TCPTransportOpts{
		ListenAddr:    ":8080",
		HandshakeFunc: tcp.NOPHandshakeFunc,
		Vm:            vm,
		Decoder:       tcp.DefaultDecoder{},
	}
	quitChan := make(chan os.Signal, 1)
	signal.Notify(quitChan, os.Interrupt, syscall.SIGTERM)

	tcpServer := tcp.NewTcp(tcpOpts)
	fmt.Println("starting ")
	go func() {
		tcpServer.Start()
	}()

	<-quitChan
	tcpServer.Shutdown()

	// // Create a simple program
	// vm.ExampleProgram(
	// 	vm.I1(vm.M.Mov, 0x04),
	// 	I0(Ste),
	// 	I1(Push, 0x00),
	// 	I1(0x09, 0x5005),
	// 	I1(Pop, 0x01),
	// 	I0(Hlt),
	// )

	fmt.Printf("VM initialized: %p (Memory size: %d)\n", vm, len(vm.M))
}
