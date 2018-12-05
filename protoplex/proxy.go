package protoplex

import (
	"net"
)

func proxy(from net.Conn, to net.Conn) {
	defer from.Close()
	var err error

	for {
		data := make([]byte, 1)
		_, err = from.Read(data)
		if err != nil {
			return
		}
		_, err = to.Write(data)
		if err != nil {
			return
		}
	}
}
