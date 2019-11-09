// +build go1.11

package vsock

import (
	"net"

	"github.com/mdlayher/vsock"
)

func dial(cid, port uint32) (net.Conn, error) {
	return vsock.Dial(cid, port)
}

// Listen listens for a connection.
func Listen(port uint32) (net.Listener, error) {
	return vsock.Listen(port)
}