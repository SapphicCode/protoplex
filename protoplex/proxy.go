package protoplex

import (
	"net"
)

func proxy(from net.Conn, to net.Conn) {
	defer from.Close()

	data := make([]byte, 4096) // 4KiB buffer

	for {
		n, err := from.Read(data)
		if err != nil {
			return
		}
		_, err = to.Write(data[:n])
		if err != nil {
			return
		}
	}
}
