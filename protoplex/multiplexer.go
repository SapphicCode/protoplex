package protoplex

import (
	"./protocols"
	"bytes"
	"fmt"
	"net"
	"time"
)

func RunServer(bind string, p []*protocols.Protocol) {
	fmt.Println("Protocol chain:")
	for _, proto := range p {
		if proto.Target != "" {
			fmt.Printf("- %s @ %s\n", proto.Name, proto.Target)
		}
	}

	listener, err := net.Listen("tcp", bind)
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	fmt.Printf("Listening at %s...\n", listener.Addr())
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error while accepting connection: %s\n", err)
		}
		fmt.Printf("Accepted connection from %s.\n", conn.RemoteAddr())
		go connectionHandler(conn, p)
	}
}

func connectionHandler(conn net.Conn, p []*protocols.Protocol) {
	connectionId := conn.RemoteAddr().String()

	identifyBuffer := make([]byte, 0, 1024) // at max 1KB buffer to identify payload
	readBuffer := make([]byte, 1)
	var protocol *protocols.Protocol

	for protocol == nil {
		// read a byte and add it to our internal buffer
		conn.SetReadDeadline(time.Now().Add(15 * time.Second)) // 15-second timeout to identify
		_, err := conn.Read(readBuffer)
		if err != nil {
			return
		}
		identifyBuffer = append(identifyBuffer, readBuffer[0])

		// determine the protocol
		protocol = determineProtocol(identifyBuffer, p)
		if len(identifyBuffer) == cap(identifyBuffer) { // unsuccessful protocol identify, close and forget
			conn.Close()
			return
		}
	}
	conn.SetReadDeadline(time.Time{})
	fmt.Printf("%s: Recognized protocol %s.\n", connectionId, protocol.Name)

	targetConn, err := net.Dial("tcp", protocol.Target)
	if err != nil {
		conn.Close()
		fmt.Printf("%s: %s rejected our connection.\n", connectionId, protocol.Target)
		return // we were unable to establish the connection with the proxy target
	}
	_, err = targetConn.Write(identifyBuffer) // tell them everything they just told us
	if err != nil {
		conn.Close()
		fmt.Printf("%s: %s cut off our identification payload.\n", connectionId, protocol.Target)
		return // remote rejected us?? okay.
	}
	go proxy(conn, targetConn)
	go proxy(targetConn, conn)
}

func determineProtocol(data []byte, p []*protocols.Protocol) *protocols.Protocol {
	dataLength := len(data)
	for _, protocol := range p {
		// since every protocol is different, let's limit the way we match things
		if (protocol.NoComparisonBeforeBytes != 0 && dataLength < protocol.NoComparisonBeforeBytes) ||
			(protocol.NoComparisonAfterBytes != 0 && dataLength > protocol.NoComparisonAfterBytes) {
			continue // avoids unnecessary comparisons
		}

		for _, byteSlice := range protocol.MatchBytes {
			byteSliceLength := len(byteSlice)
			if dataLength < byteSliceLength {
				continue
			}
			if bytes.Equal(byteSlice, data[0:byteSliceLength]) {
				return protocol
			}
		}

		// let's use regex matching as a last resort
		for _, regex := range protocol.MatchRegexes {
			if regex.Match(data) {
				return protocol
			}
		}
	}
	return nil
}
