package protocols

func NewSOCKS5Protocol(targetAddress string) *Protocol {
	return &Protocol{
		Name:                    "SOCKS5",
		Target:                  targetAddress,
		MatchBytes:              [][]byte{{0x05}},
		NoComparisonBeforeBytes: 1,
	}
}
