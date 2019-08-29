package protocols

// NewSTRelayProtocol initializes a Protocol with a Syncthing Relay signature.
//
// Deprecated: This signature does not function properly and unless further developed, will not establish a working
// connection.
func NewSTRelayProtocol(targetAddress string) *Protocol {
	return &Protocol{
		Name:   "STRelay",
		Target: targetAddress,
		MatchBytes: [][]byte{
			{'b', 'e', 'p', '-', 'r', 'e', 'l', 'a', 'y'},
			{'b', 'e', 'p', '/'},
		},
	}
}
