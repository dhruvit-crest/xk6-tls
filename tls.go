package udp

import (
	"crypto/tls"
	"net"
	"time"

	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/tls", new(TLS))
}

const maxPackageSize = 63 * 1024 // 63KiB

var newLine = []byte("\n")

type TLS struct{}

func (u *TLS) Connect(address string) (net.Conn, error) {
	conn, err := tls.Dial("tcp", address, nil)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (u *TLS) Writeln(conn net.Conn, data []byte) error {
	return u.Write(conn, append(data, newLine...))
}

func (u *TLS) Write(conn net.Conn, data []byte) error {
	for i := 0; i < len(data); i += maxPackageSize {
		end := i + maxPackageSize
		if end > len(data) {
			end = len(data)
		}

		chunk := data[i:end]

		_, err := conn.Write(chunk)
		if err != nil {
			return err
		}
	}

	return nil
}

func (u *TLS) Read(conn net.Conn, size int, timeout_opt ...int) ([]byte, error) {
	timeout_ms := 0
	if len(timeout_opt) > 0 {
		timeout_ms = timeout_opt[0]
	}

	err := conn.SetReadDeadline(time.Now().Add(time.Millisecond * time.Duration(timeout_ms)))
	if err != nil {
		return nil, err
	}

	buf := make([]byte, size)

	_, err = conn.Read(buf)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (u *TLS) Close(conn net.Conn) error {
	err := conn.Close()
	if err != nil {
		return err
	}

	return nil
}
