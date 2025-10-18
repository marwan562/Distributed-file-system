package main

import (
	"github.com/Distributed-file-system/p2p"
)

func main() {
	// Create a new TCP transport listening on port 3000
	// In a real application, you would implement proper handshake
	// and decoding logic.
	tcpOpts := p2p.TCPTransportOpts{
		HandshakeFunc: p2p.NOPHandshakerFunc,
		Decoder:       &p2p.DefaultDecoder{},
	}
	tr := p2p.NewTCPTransport(":3000", tcpOpts)
	rpcch := tr.Consume()
	if err := tr.ListenAndAccept(); err != nil {
		panic(err)
	}
	for {
		rpc := <-rpcch
		println("RPC message:", string(rpc.Payload))
	}

}
