package protocols

func NewTLSProtocol(targetAddress string) *Protocol {
	return &Protocol{
		Name:                    "TLS",
		Target:                  targetAddress,
		MatchBytes:              [][]byte{{'\x16', '\x03', '\x01'}},
		NoComparisonBeforeBytes: 3,
		NoComparisonAfterBytes:  3,
	}
}
