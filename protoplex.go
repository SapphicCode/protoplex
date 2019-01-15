package main

import (
	"./protoplex"
	"./protoplex/protocols"
	"flag"
	"github.com/juju/loggo"
)

func main() {
	logger := loggo.GetLogger("protoplex")

	bind := flag.String("bind", "0.0.0.0:8443", "The address to bind to")
	verbose := flag.Bool("verbose", false, "Whether to be verbose")

	ssh := flag.String("ssh", "", "The SSH server address")
	tls := flag.String("tls", "", "The TLS/HTTPS server address")
	openvpn := flag.String("ovpn", "", "The OpenVPN server address")
	http := flag.String("http", "", "The HTTP server address")
	socks5 := flag.String("socks5", "", "The SOCKS5 server address")
	socks4 := flag.String("socks4", "", "The SOCKS4 server address")
	stRelay := flag.String("strelay", "", "The Syncthing Relay server address")

	flag.Parse()

	if *verbose {
		logger.SetLogLevel(loggo.DEBUG)
	} else {
		logger.SetLogLevel(loggo.INFO)
	}

	p := make([]*protocols.Protocol, 0, 7)
	// contain-bytes-matched protocols (usually ALPNs) take priority
	// (due to start-bytes-matching overriding some of them)
	if *stRelay != "" {
		p = append(p, protocols.NewSTRelayProtocol(*stRelay))
	}
	// start-bytes-matched protocols are the next most efficient approach
	if *ssh != "" {
		p = append(p, protocols.NewSSHProtocol(*ssh))
	}
	if *tls != "" {
		p = append(p, protocols.NewTLSProtocol(*tls))
	}
	if *openvpn != "" {
		p = append(p, protocols.NewOpenVPNProtocol(*openvpn))
	}
	if *socks5 != "" {
		p = append(p, protocols.NewSOCKS5Protocol(*socks5))
	}
	if *socks4 != "" {
		p = append(p, protocols.NewSOCKS4Protocol(*socks4))
	}
	// regex protocols come at the end of the chain as they'll be expensive anyway if used
	if *http != "" {
		p = append(p, protocols.NewHTTPProtocol(*http))
	}

	protoplex.RunServer(*bind, p)
}
