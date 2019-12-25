package gotcp

import (
	"context"
	"net"
	"time"
)

type ReaderInterface interface {
	HeaderSize() uint8
	NewHeader([]byte) (HeaderInterface, error)
	Timeout() time.Duration
	HandleError(error)
}

type HeaderInterface interface {
	GetSize() uint16
	GetCmd() uint16
}

type TransportInterface interface {
	Clear(conn net.Conn)
	Handle(ctx context.Context, conn net.Conn, cmd uint16, buff []byte)
}

type Package struct {
	conn net.Conn
	cmd  uint16
	buff []byte
}

func newPackage(conn net.Conn, cmd uint16, buff []byte) *Package {
	return &Package{
		conn: conn,
		cmd:  cmd,
		buff: buff,
	}
}
