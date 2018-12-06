package protocols

func NewTLSProtocol(targetAddress string) *Protocol {
	return &Protocol{
		Name:                    "TLS",
		Target:                  targetAddress,
		MatchBytes:              [][]byte{{0x16, 0x03, 0x01}},
		NoComparisonBeforeBytes: 3,
	}
}
