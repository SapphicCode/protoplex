package protocols

func NewSSHProtocol(targetAddress string) *Protocol {
	return &Protocol{
		Name:                    "SSH",
		Target:                  targetAddress,
		MatchBytes:              [][]byte{{'S', 'S', 'H', '-'}},
		NoComparisonBeforeBytes: 4,
		NoComparisonAfterBytes:  4,
	}
}
