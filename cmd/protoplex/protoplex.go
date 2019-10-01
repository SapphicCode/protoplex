package main

import (
	"fmt"
	"os"

	"github.com/Pandentia/protoplex/protoplex"
	"github.com/Pandentia/protoplex/protoplex/protocols"
	"github.com/rs/zerolog"
	"gopkg.in/alecthomas/kingpin.v2"
)

var version string

func printVersion() {
	if version == "" {
		fmt.Println("Version has not been set.")
		os.Exit(1)
		return
	}
	fmt.Println(version)
	os.Exit(0)
}

func main() {
	app := kingpin.New("protoplex", "A fast and simple protocol multiplexer.")
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	version := app.Flag("version", "Prints the current program version").Short('V').Bool()

	bind := app.Flag("bind", "The address to bind to").Short('b').Default("0.0.0.0:8443").String()
	verbose := app.Flag("verbose", "Enables debug logging").Short('v').Bool()
	pretty := app.Flag("pretty", "Enables pretty logging").Short('p').Bool()

	ssh := app.Flag("ssh", "The SSH server address").String()
	tls := app.Flag("tls", "The TLS/HTTPS server address").String()
	openvpn := app.Flag("ovpn", "The OpenVPN server address").String()
	http := app.Flag("http", "The HTTP server address").String()
	socks5 := app.Flag("socks5", "The SOCKS5 server address").String()
	socks4 := app.Flag("socks4", "The SOCKS4 server address").String()
	// stRelay := flag.String("strelay", "", "The Syncthing Relay server address")

	app.Parse(os.Args[1:])

	if *version {
		printVersion()
	}

	if *pretty {
		logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	if *verbose {
		logger = logger.Level(zerolog.DebugLevel)
	} else {
		logger = logger.Level(zerolog.InfoLevel)
	}

	p := make([]*protocols.Protocol, 0, 7)
	// contain-bytes-matched protocols (usually ALPNs) take priority
	// (due to start-bytes-matching overriding some of them)
	// if *stRelay != "" {
	// 	logger.Warningf("Syncthing Relay support is deprecated.\n")
	// 	p = append(p, protocols.NewSTRelayProtocol(*stRelay))
	// }
	// start-bytes-matched protocols are the next most efficient approach
	if *tls != "" {
		p = append(p, protocols.NewTLSProtocol(*tls))
	}
	if *ssh != "" {
		p = append(p, protocols.NewSSHProtocol(*ssh))
	}
	if *socks5 != "" {
		p = append(p, protocols.NewSOCKS5Protocol(*socks5))
	}
	if *socks4 != "" {
		p = append(p, protocols.NewSOCKS4Protocol(*socks4))
	}
	// regex protocols come at the end of the chain as they'll be expensive anyway if used
	if *openvpn != "" {
		p = append(p, protocols.NewOpenVPNProtocol(*openvpn))
	}
	if *http != "" {
		p = append(p, protocols.NewHTTPProtocol(*http))
	}

	protoplex.RunServer(*bind, p, logger)
}
