package tcp

import (
	"encoding/gob"
	"fmt"
	"io"
	"strings"
)

type Decoder interface {
	Decode(io.Reader, *RPC) error
}

type GOBDecoder struct{}

func (dec GOBDecoder) Decode(r io.Reader, msg *RPC) error {
	return gob.NewDecoder(r).Decode(msg)
}

type DefaultDecoder struct{}

func (dec DefaultDecoder) Decode(r io.Reader, msg *RPC) error {
	buf := make([]byte, 2048)
	n, err := r.Read(buf)
	if err != nil {
		return err
	}

	// Example: parse "MOV 0x04"
	input := strings.TrimSpace(string(buf[:n]))
	parts := strings.Fields(input)

	if len(parts) < 1 {
		return fmt.Errorf("invalid input")
	}

	msg.Method = strings.ToUpper(parts[0])
	msg.Args = []int16{}

	for _, arg := range parts[1:] {
		var val int16
		_, err := fmt.Sscanf(arg, "%x", &val)
		if err != nil {
			return fmt.Errorf("invalid argument: %s", arg)
		}
		msg.Args = append(msg.Args, val)
	}

	return nil
}
