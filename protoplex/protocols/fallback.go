package protocols

// FallbackProtocol initializes a Protocol with a SOCKS4 signature.
func FallbackProtocol(targetAddress string) *Protocol {
	return &Protocol{
		Name:       "Fallback",
		Target:     targetAddress,
		MatchBytes: [][]byte{{}},
	}
}
