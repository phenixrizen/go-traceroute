package traceroute

import (
	"net"
	"time"
)

// TraceData represents data received by executing traceroute.
type TraceData struct {
	Hops    [][]Hop       `json:"hops,omitempty"`
	Dest    net.IP        `json:"dest,omitempty"`
	Timeout time.Duration `json:"timeout,omitempty"`
	Tries   int           `json:"tries,omitempty"`
	MaxTTL  int           `json:"max_ttl,omitempty"`
	Port    int           `json:"port,omitempty"`
	Proto   string        `json:"proto,omitempty"`
	IPv     string        `json:"i_pv,omitempty"`
}

// Hop represents a path between a source and a destination.
type Hop struct {
	TryNumber int           `json:"try_numberomitempty"`
	TTL       int           `json:"ttl,omitempty"`
	AddrIP    net.IP        `json:"addr_ip,omitempty"`
	AddrDNS   []string      `json:"addr_dns,omitempty"` //net.IPAddr
	Latency   time.Duration `json:"latency,omitempty"`
	Err       error         `json:"err,omitempty"`
}
