package protocols

// NewSOCKS4Protocol initializes a Protocol with a SOCKS4 signature.
func NewSOCKS4Protocol(targetAddress string) *Protocol {
	return &Protocol{
		Name:            "SOCKS4",
		Target:          targetAddress,
		MatchStartBytes: [][]byte{{0x04}},
	}
}
