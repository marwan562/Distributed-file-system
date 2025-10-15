package p2p

// Peer represents a node in the P2P network.
type Peer interface {
}

// Transport defines the methods for a transport mechanism.
// and it's anythink to handle different transport implementations
// from TCP, UDP, WebRTC ..ect
// from ("p2p/tcp_transport.go")
type Transport interface {
	ListenAndAccept() error
}
