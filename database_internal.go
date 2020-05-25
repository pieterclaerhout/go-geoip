package geoip

import (
	"errors"
	"net"
)

// isPrivateIP returns true if the IP address is an internal one
func isPrivateIP(ipAddress string) (net.IP, bool, error) {

	ip := net.ParseIP(ipAddress)
	if ip == nil {
		return nil, false, errors.New("Invalid IP")
	}

	_, private24BitBlock, _ := net.ParseCIDR("10.0.0.0/8")
	_, private20BitBlock, _ := net.ParseCIDR("172.16.0.0/12")
	_, private16BitBlock, _ := net.ParseCIDR("192.168.0.0/16")
	private := private24BitBlock.Contains(ip) || private20BitBlock.Contains(ip) || private16BitBlock.Contains(ip)

	return ip, private, nil

}
