package main

import (
	"./protoplex"
	"./protoplex/protocols"
	"flag"
)

func main() {
	bind := flag.String("bind", "0.0.0.0:8443", "The address to bind to")
	ssh := flag.String("ssh", "", "The SSH server address")
	tls := flag.String("tls", "", "The TLS/HTTPS server address")
	openvpn := flag.String("ovpn", "", "The OpenVPN server address")
	http := flag.String("http", "", "The HTTP server address")
	flag.Parse()

	p := make([]*protocols.Protocol, 0, 4)
	if *ssh != "" {
		p = append(p, protocols.NewSSHProtocol(*ssh))
	}
	if *tls != "" {
		p = append(p, protocols.NewTLSProtocol(*tls))
	}
	if *openvpn != "" {
		p = append(p, protocols.NewOpenVPNProtocol(*openvpn))
	}
	if *http != "" {
		p = append(p, protocols.NewHTTPProtocol(*http))
	}

	protoplex.RunServer(*bind, p)
}
