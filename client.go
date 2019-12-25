package gotcp

import (
	"errors"
	"net"
	"time"
)

func send(conn net.Conn, data []byte, retryTimes int) error {
	var err error
	for i := 0; i < retryTimes; i++ {
		_, err = conn.Write(data)
		if err != nil {
			time.Sleep(1 * time.Second)
			continue
		}
		return nil
	}
	return err
}

func Send(conn net.Conn, data []byte, timeout time.Duration) error {
	if conn == nil {
		return errors.New("conn is nil")
	}
	if err := conn.SetWriteDeadline(time.Now().Add(timeout)); err != nil {
		return err
	}
	return send(conn, data, 3)
}
