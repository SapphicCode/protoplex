package protocols

func NewSTRelayProtocol(targetAddress string) *Protocol {
	return &Protocol{
		Name:   "STRelay",
		Target: targetAddress,
		MatchBytes: [][]byte{
			{'b', 'e', 'p', '-', 'r', 'e', 'l', 'a', 'y'},
			{'b', 'e', 'p', '/'},
		},
	}
}
