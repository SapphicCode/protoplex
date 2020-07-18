# protoplex

*An application protocol multiplexer*

[![Build Status](https://cloud.drone.io/api/badges/Pandentia/protoplex/status.svg?ref=refs/heads/mistress)](https://cloud.drone.io/Pandentia/protoplex)

## What is this?

In a nutshell, this application lets you run multiple kinds of applications
on a single port. This is useful for, for instance, running an OpenVPN server
and a TLS/HTTPS server on port 443, which in turn is useful for evading
firewalls that block all other outbound ports.

## Running

### Native

Assuming you have a properly configured Go setup, get and compile the multiplexer with

```bash
go get github.com/Pandentia/protoplex/cmd/protoplex
```

and then run it with (for example, to run SSH and HTTPS)

```bash
protoplex --ssh your_ssh_host:22 --tls your_webserver:443
```

Protoplex is now running on port `8443` and ready to accept connections.

For more extensive configuration, please see the output of `--help`.

### Docker

[A docker image may be used](https://hub.docker.com/r/pandentia/protoplex)
for ease of use and deployment.

## Goals

The concepts for this multiplexer were as follows:

- Resource usage about on par with `sslh`
- Easily extensible
- Highly dynamic

To this end, protoplex supports multiple matching methods for protocols:

- Bytestring comparison
- Regex matching

These can both be implemented for a protocol, with bytestrings taking
priority (due to efficiency). In addition, protocols support matching limits,
reducing the amount of protocols evaluated for a given handshake.

## Protocol support

Currently supported protocols are:

- SSH
- HTTP
- TLS (/ HTTPS)
- OpenVPN
- SOCKS4 / SOCKS5

Feel free to [file an issue](https://github.com/Pandentia/protoplex/issues/new)
on the GitHub repository if you want a protocol to be supported. Please include
steps to accurately reproduce your client setup.

Alternatively, you may submit a pull request.
