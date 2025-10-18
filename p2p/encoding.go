package p2p

import (
	"encoding/gob"
	"net"
)

type Decoder interface {
	Decode(net.Conn, *RPC) error
}

type GOBDecoder struct{}

func (d *GOBDecoder) Decode(r net.Conn, rpc *RPC) error {
	return gob.NewDecoder(r).Decode(rpc)
}

type DefaultDecoder struct{}

func (d *DefaultDecoder) Decode(r net.Conn, rpc *RPC) error {
	buf := make([]byte, 1024)
	n, err := r.Read(buf)
	if err != nil {
		return err
	}

	rpc.From = r.RemoteAddr()
	rpc.Payload = buf[:n]
	return nil
}
