package protocols

// NewSSHProtocol initializes a Protocol with a SSH signature.
func NewSSHProtocol(targetAddress string) *Protocol {
	return &Protocol{
		Name:            "SSH",
		Target:          targetAddress,
		MatchStartBytes: [][]byte{{'S', 'S', 'H', '-'}},
	}
}
