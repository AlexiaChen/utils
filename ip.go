package utils

import (
	"net"
	"strings"
)

// IsValidIP 验证IP是否合法
func IsValidIP(ip string) bool {
	return net.ParseIP(ip) != nil
}

// IsValidCIDRIP CIDRI验证 ip 是否合法
func IsValidCIDRIP(ip string) bool {
	if IsValidIP(ip) {
		return true
	}
	_, _, err := net.ParseCIDR(ip)
	if err != nil {
		return false
	}
	return true
}

// ValidateIpList 验证IP列表是否合法
func ValidateIpList(ipList string) bool {
	if strings.TrimSpace(ipList) == "" {
		return false
	}
	ips := strings.Split(ipList, "\n")
	for _, ip := range ips {
		if strings.TrimSpace(ip) != "" && !IsValidIP(strings.TrimSpace(ip)) {
			return false
		}
	}
	return true
}

// ValidateCIDRIIpList ValidateIpList CIDRI验证IP列表是否合法
func ValidateCIDRIIpList(ipList string) bool {
	if strings.TrimSpace(ipList) == "" {
		return false
	}
	ips := strings.Split(ipList, "\n")
	for _, ip := range ips {
		if strings.TrimSpace(ip) != "" && !IsValidCIDRIP(strings.TrimSpace(ip)) {
			return false
		}
	}
	return true
}

// IPv4InSubnet ipv4是否在给定的网段内
func IPv4InSubnet(ipAddress string, subnetCIDR string) bool {
	ip := net.ParseIP(ipAddress)
	if ip == nil {
		return false
	}

	_, subnet, err := net.ParseCIDR(subnetCIDR)
	if err != nil {
		return false
	}

	return subnet.Contains(ip)
}
