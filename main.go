/*
**
** main
** Provides an IPv4 DNS gateway to make NAT64 easier
**
** Distributed under the COOL License.
**
** Copyright (c) 2024 IPv6.rs <https://ipv6.rs>
** All Rights Reserved
**
*/

package main

import (
  "fmt"
  "log"
  "net"
  "strings"
  "github.com/miekg/dns"
)

const sld = "visibleip.com."

func main() {
  dns.HandleFunc(".", handleDNSRequest)
  server := &dns.Server{Addr: ":53", Net: "udp"}
  log.Println("Starting DNS server on port 53")
  err := server.ListenAndServe()
  if err != nil {
    log.Fatalf("Failed to start server: %s\n", err.Error())
  }
  defer server.Shutdown()
}

func handleDNSRequest(w dns.ResponseWriter, r *dns.Msg) {
  msg := dns.Msg{}
  msg.SetReply(r)
  msg.Authoritative = true

  for _, q := range r.Question {
    if q.Qtype == dns.TypeA && strings.HasSuffix(q.Name, sld) {
      rr := createARecord(q.Name)
      if rr != nil {
        msg.Answer = append(msg.Answer, rr)
      }
    }
  }

  w.WriteMsg(&msg)
}

func createARecord(queryName string) *dns.A {
  octets := strings.Split(queryName, ".")

  ip, err := getIPFromQuery(octets)
  if err != nil {
    return nil
  }

  return &dns.A{
    Hdr: dns.RR_Header{
      Name:   queryName,
      Rrtype: dns.TypeA,
      Class:  dns.ClassINET,
      Ttl:    0,
    },
    A: ip,
  }
}

func getIPFromQuery(octets []string) (net.IP, error) {
  if len(octets) < 5 {
    return nil, fmt.Errorf("invalid query format")
  }
  ipStr := octets[0] + "." + octets[1] + "." + octets[2] + "." + octets[3]
  ip := net.ParseIP(ipStr)
  if ip == nil {
    return nil, fmt.Errorf("failed to parse IP: %s", ipStr)
  }
  return ip, nil
}

