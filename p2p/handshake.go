package p2p

type HandshakerFunc func(Peer) error

func NOPHandshakerFunc(p Peer) error {
	return nil
}
