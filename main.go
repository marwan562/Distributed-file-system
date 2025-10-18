package main

import (
	"github.com/Distributed-file-system/p2p"
	"github.com/Distributed-file-system/storage"
)

func OnPeer(p *p2p.TCPPeer) error {
	p.Close()
	return nil
}

func main() {
	// Create a new TCP transport listening on port 3000
	// In a real application, you would implement proper handshake
	// and decoding logic.

	// listenAddr := ":3000"
	// tcpOpts := p2p.TCPTransportOpts{
	// 	HandshakeFunc: p2p.NOPHandshakerFunc,
	// 	Decoder:       &p2p.DefaultDecoder{},
	// 	OnPeer:        OnPeer,
	// }
	// tr := p2p.NewTCPTransport(listenAddr, tcpOpts)
	// if err := tr.ListenAndAccept(); err != nil {
	// 	panic(err)
	// }

	// rpcch := tr.Consume()
	// go func() {
	// 	for rpc := range rpcch {
	// 		println("Received message:", string(rpc.Payload))
	// 	}
	// }()

	optsStorage := storage.StorageOpts{
		PathTransformFunc: storage.DefaultPathTransformFunc,
	}

	stg := storage.NewStorage(optsStorage)
}
