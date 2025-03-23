package cidr

import (
	"fmt"
	"net"
)

type Address struct {
	IP   net.IP
	Mask net.IP
}

// Parse parses IP address from CIDR notation
func Parse(input string) (*Address, error) {
	_, ipNet, err := net.ParseCIDR(input)
	if err != nil {
		return nil, fmt.Errorf("failed to parse: %s", input)
	}

	return &Address{
		IP:   ipNet.IP,
		Mask: net.IP(ipNet.Mask),
	}, nil
}
