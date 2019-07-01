package protocols

import "regexp"

// NewOpenVPNProtocol initializes a Protocol with an OpenVPN signature.
func NewOpenVPNProtocol(targetAddress string) *Protocol {
	return &Protocol{
		Name:   "OpenVPN",
		Target: targetAddress,
		// MatchStartBytes: [][]byte{{0x00, 0x0e, 0x38}},
		MatchRegexes: []*regexp.Regexp{
			regexp.MustCompile(`^\x00[\x0d-\xff]\x38`), // asumming this variant is more common in newer clients
			regexp.MustCompile(`^\x00[\x0d-\xff]$`),
		},
	}
}
