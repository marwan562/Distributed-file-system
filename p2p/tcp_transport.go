package p2p

import (
	"fmt"
	"net"
	"sync"
)

// TCPPeer represents the remote node over TCP established connection.
type TCPPeer struct {
	// conn is underlying connection of this peer
	conn net.Conn
	// if we dial and retrieve the connection, outbound it will be == true
	// if we accept the connection, outbound it will be == false
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

type TCPTransport struct {
	listenAddr string
	listener   net.Listener
	shakeHands HandshakerFunc
	decoder    Decoder

	// mutex its a lock to protect the peer map
	mu    sync.Mutex
	peers map[net.Addr]Peer
}

func NewTCPTransport(listenAddr string) *TCPTransport {
	return &TCPTransport{
		listenAddr: listenAddr,
		shakeHands: NOPHandshakerFunc,
	}
}

// ListenAndAccept starts listening on the specified TCP address and accepts incoming connections.
func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.listener, err = net.Listen("tcp", t.listenAddr)
	if err != nil {
		return err
	}

	go t.startAcceptLoop()
	return nil
}

// startAcceptLoop continuously accepts incoming TCP connections.
func (t *TCPTransport) startAcceptLoop() error {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept error:%s", err)
		}
		go t.handleConn(conn)
	}
}

type Message struct {
}

func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := NewTCPPeer(conn, true)

	if err := t.shakeHands(peer); err != nil {
		fmt.Printf("Handshake error: %s", err)
		conn.Close()
		return
	}

	// Read loop
	msg := &Message{}
	for {
		if err := t.decoder.Decode(conn, msg); err != nil {
			fmt.Printf("Decode error: %s", err)
			continue
		}
	}
}
