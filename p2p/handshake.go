package p2p

import "errors"

// ErrInvalidHandshake is returned when a handshake fails.
var ErrInvalidHandshake = errors.New("invalid handshake")

type HandshakerFunc func(Peer) error

func NOPHandshakerFunc(p Peer) error {
	return nil
}
