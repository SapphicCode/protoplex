package protocols

import "regexp"

type Protocol struct {
	Name                    string           // the protocol name for auditing
	Target                  string           // the proxy target
	MatchBytes              [][]byte         // the bytestrings by which to match this protocol (prefixes)
	MatchRegexes            []*regexp.Regexp // the regexes by which to match this protocol
	NoComparisonBeforeBytes int              // we know our regexes won't match before this many bytes, set to 0 to ignore
	NoComparisonAfterBytes  int              // we know our regexes won't match after this many bytes, set to 0 to ignore
}
