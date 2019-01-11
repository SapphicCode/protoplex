package protocols

func NewSOCKS4Protocol(targetAddress string) *Protocol {
	return &Protocol{
		Name:                    "SOCKS4",
		Target:                  targetAddress,
		MatchBytes:              [][]byte{{0x04}},
		NoComparisonBeforeBytes: 1,
	}
}
