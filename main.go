package main

import (
	"github.com/Distributed-file-system/p2p"
)

func main() {
	tr := p2p.NewTCPTransport(":3000")

	if err := tr.ListenAndAccept(); err != nil {
		panic(err)
	}

	select {}
}
