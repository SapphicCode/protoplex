package protoplex

import (
	"./protocols"
	"bytes"
	"github.com/juju/loggo"
	"net"
	"os"
	"time"
)

func RunServer(bind string, p []*protocols.Protocol) {
	logger := loggo.GetLogger("protoplex.listener")

	if len(p) == 0 {
		logger.Warningf("No protocols defined.\n")
	} else {
		logger.Infof("Protocol chain:\n")
		for _, proto := range p {
			logger.Infof("- %s @ %s\n", proto.Name, proto.Target)
		}
	}

	listener, err := net.Listen("tcp", bind)
	if err != nil {
		logger.Criticalf("Unable to create listener: %s\n", err)
		os.Exit(1)
	}
	defer listener.Close()
	logger.Infof("Listening at %s...\n", listener.Addr())
	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.Debugf("Error while accepting connection: %s\n", err)
		}
		logger.Debugf("%s: Connection accepted.\n", conn.RemoteAddr())
		go connectionHandler(conn, p)
	}
}

func connectionHandler(conn net.Conn, p []*protocols.Protocol) {
	defer conn.Close() // the connection must close after this goroutine exits
	connectionId := conn.RemoteAddr().String()
	logger := loggo.GetLogger("protoplex.connection")

	identifyBuffer := make([]byte, 1024) // at max 1KB buffer to identify payload

	// read the handshake into our buffer
	_ = conn.SetReadDeadline(time.Now().Add(15 * time.Second)) // 15-second timeout to identify
	n, err := conn.Read(identifyBuffer)
	if err != nil {
		logger.Debugf("%s: Identify read error (%s). Connection closed.\n", connectionId, err)
		return
	}
	_ = conn.SetReadDeadline(time.Time{}) // reset our timeout

	// determine the protocol
	protocol := determineProtocol(identifyBuffer[:n], p)
	if protocol == nil { // unsuccessful protocol identify, close and forget
		logger.Debugf("%s: Protocol unrecognized. Connection closed.\n", connectionId)
		return
	}
	logger.Debugf("%s: Recognized protocol %s.\n", connectionId, protocol.Name)

	// establish our connection with the target
	targetConn, err := net.Dial("tcp", protocol.Target)
	if err != nil {
		logger.Debugf("%s: %s error (%s). Connection closed.\n", connectionId, protocol.Target, err)
		return // we were unable to establish the connection with the proxy target
	}
	defer targetConn.Close()
	_, err = targetConn.Write(identifyBuffer[:n]) // tell them everything they just told us
	if err != nil {
		logger.Debugf("%s: %s error (%s). Connection closed.\n", connectionId, protocol.Target, err)
		return // remote rejected us?? okay.
	}

	// run the proxy readers
	closed := make(chan bool, 2)
	go proxy(conn, targetConn, closed)
	go proxy(targetConn, conn, closed)

	// wait for any connection to close
	<-closed
	logger.Debugf("%s: Connection closed.\n", connectionId)
}

func determineProtocol(data []byte, p []*protocols.Protocol) *protocols.Protocol {
	dataLength := len(data)
	for _, protocol := range p {
		// since every protocol is different, let's limit the way we match things
		if (protocol.NoComparisonBeforeBytes != 0 && dataLength < protocol.NoComparisonBeforeBytes) ||
			(protocol.NoComparisonAfterBytes != 0 && dataLength > protocol.NoComparisonAfterBytes) {
			continue // avoids unnecessary comparisons
		}

		// compare against bytestrings first for efficiency
		// first "contains" (due to ALPNs we can't match against TLS start bytes first)
		for _, byteSlice := range protocol.MatchBytes {
			byteSliceLength := len(byteSlice)
			if dataLength < byteSliceLength {
				continue
			}
			if bytes.Contains(data, byteSlice) {
				return protocol
			}
		}
		// then against prefixes
		for _, byteSlice := range protocol.MatchStartBytes {
			byteSliceLength := len(byteSlice)
			if dataLength < byteSliceLength {
				continue
			}
			if bytes.Equal(byteSlice, data[:byteSliceLength]) {
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
