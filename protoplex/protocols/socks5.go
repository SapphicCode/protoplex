package protocols

func NewSOCKS5Protocol(targetAddress string) *Protocol {
	return &Protocol{
		Name:            "SOCKS5",
		Target:          targetAddress,
		MatchStartBytes: [][]byte{{0x05}},
	}
}
