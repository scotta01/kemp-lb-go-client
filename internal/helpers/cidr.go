package helpers

import (
	"fmt"
	"net"
)

func CheckIPInRange(CIDR string, ipAddress string) bool {
	_, ipNet, err := net.ParseCIDR(CIDR)
	if err != nil {
		fmt.Println("Error parsing CIDR:", err)
		return false
	}

	ip := net.ParseIP(ipAddress)
	if ip == nil {
		fmt.Println("Error parsing IP address")
		return false
	}

	if ipNet.Contains(ip) {
		return true
	} else {
		return false
	}
}
