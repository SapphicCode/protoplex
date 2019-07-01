package protocols

// NewSTRelayProtocol initializes a Protocol with a Syncthing Relay signature.
// This signature is deprecated.
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
