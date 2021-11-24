package provider

import (
	"context"
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
)

func TestWaiterTCP_Wait(t *testing.T) {
	var waiter WaiterTCP

	conn, _ := net.Listen("tcp", ":8180")

	go func() {
		for {
			_, _ = conn.Accept()
			_ = conn.Close()
			break
		}
	}()

	err := waiter.Wait(context.Background(), "localhost:8180", 10, 1)

	assert.Nil(t, err)
}
