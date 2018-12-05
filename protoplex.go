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
	http := flag.String("http", "", "The HTTP server address")
	flag.Parse()

	p := []*protocols.Protocol{
		protocols.NewSSHProtocol(*ssh),
		protocols.NewTLSProtocol(*tls),
		protocols.NewHTTPProtocol(*http),
	}

	protoplex.RunServer(*bind, p)
}
