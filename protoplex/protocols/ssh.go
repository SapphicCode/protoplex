package protocols

func NewSSHProtocol(targetAddress string) *Protocol {
	return &Protocol{
		Name:            "SSH",
		Target:          targetAddress,
		MatchStartBytes: [][]byte{{'S', 'S', 'H', '-'}},
	}
}
