package protocols

func NewSOCKS4Protocol(targetAddress string) *Protocol {
	return &Protocol{
		Name:            "SOCKS4",
		Target:          targetAddress,
		MatchStartBytes: [][]byte{{0x04}},
	}
}
