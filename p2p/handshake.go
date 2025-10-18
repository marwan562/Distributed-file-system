package p2p

import "errors"

// ErrInvalidHandshake is returned when a handshake fails.
var ErrInvalidHandshake = errors.New("invalid handshake")

type HandshakerFunc func(*TCPPeer) error

func NOPHandshakerFunc(p *TCPPeer) error {
	return nil
}
