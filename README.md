# legacydns for golang

Provides a legacy DNS server for users who are on NAT64.

## Features

- Responds with the A record corresponding a name, like `1.1.1.1.visibleip.com` will return `1.1.1.1`

## Example

- Try it out, type `host IPv4_ADDRESS.visibleip.com` in your terminal.

## Use Case

- This makes it easier to work with your grandpa's legacy IPv4 networks.

## Usage

- Edit the DNS (visibleip.com)
- Type `go build main.go`
- Type `./main`

You'll need to setup a Nameserver and DNS prior, of course.

## License

Distributed under the COOL License.

Copyright (c) 2024 IPv6.rs <https://ipv6.rs>
All Rights Reserved
