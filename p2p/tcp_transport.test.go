package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	listenAddr := ":3642"
	tr := NewTCPTransport(listenAddr)
	tr.ListenAndAccept()

	assert.Equal(t, listenAddr, tr.listenAddr)
	assert.Nil(t, tr.listener)
}
