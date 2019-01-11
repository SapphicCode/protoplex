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
	socks5 := flag.String("socks5", "", "The SOCKS5 server address")
	socks4 := flag.String("socks4", "", "The SOCKS4 server address")

	flag.Parse()

	p := make([]*protocols.Protocol, 0, 6)
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
	if *socks5 != "" {
		p = append(p, protocols.NewSOCKS5Protocol(*socks5))
	}
	if *socks4 != "" {
		p = append(p, protocols.NewSOCKS4Protocol(*socks4))
	}

	protoplex.RunServer(*bind, p)
}
