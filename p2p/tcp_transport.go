package p2p

import (
	"fmt"
	"net"
)

// TCPPeer represents the remote node over TCP established connection.
type TCPPeer struct {
	// conn is underlying connection of this peer
	conn net.Conn
	// if we dial and retrieve the connection, outbound it will be == true
	// if we accept the connection, outbound it will be == false
	outbound bool
}

type TCPTransportOpts struct {
	HandshakeFunc HandshakerFunc
	Decoder       Decoder
	OnPeer        func(*TCPPeer) error
}

type TCPTransport struct {
	TCPTransportOpts
	listenAddr string
	listener   net.Listener
	rpcch      chan RPC
}

// * Peer *
// Close closes the TCP transport.
func (p *TCPPeer) Close() error {
	fmt.Errorf("Closed connection to %s", p.conn.RemoteAddr())
	return p.conn.Close()
}

// NewTCPPeer returns  new TCPPeer instance.
func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

// * Transport *

func NewTCPTransport(listenAddr string, opts TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		listenAddr:       listenAddr,
		TCPTransportOpts: opts,
		rpcch:            make(chan RPC),
	}
}

// Consume returns a channel to consume incoming RPC messages.
func (t *TCPTransport) Consume() <-chan RPC {
	return t.rpcch
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

func (t *TCPTransport) handleConn(conn net.Conn) {

	defer func() {
		fmt.Printf("Disconnected from %s\n", conn.RemoteAddr())
		conn.Close()
	}()

	// Perform handshake
	peer := NewTCPPeer(conn, true)

	// call the OnPeer callback
	// OnPeer its optional
	if t.OnPeer != nil {
		if err := t.OnPeer(peer); err != nil {
			fmt.Printf("OnPeer error: %s", err)
			return
		}
	}

	if err := t.HandshakeFunc(peer); err != nil {
		fmt.Printf("Handshake error: %s", err)
		conn.Close()
		return
	}

	// handle the receive data from peer
	t.handleReceiveData(conn)
}

// handleReceiveData handles incoming data from the connection.
func (t *TCPTransport) handleReceiveData(conn net.Conn) {
	// Read loop
	rpc := RPC{}
	for {
		err := t.Decoder.Decode(conn, &rpc)
		if err != nil {
			return
		}
		// Send the RPC to the channel
		t.rpcch <- rpc
		fmt.Printf("Received message: %s\n", string(rpc.Payload))
	}
}
