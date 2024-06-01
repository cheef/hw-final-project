package services

import (
	"errors"
	"github.com/cheef/hw-final-project/internal/domain/models"
	"net"
)

type IPChecker struct {
	source <-chan models.ExceptionList
}

var ErrIPNotValid = errors.New("provided IP-address is invalid")

func NewIPChecker(cidrs <-chan models.ExceptionList) *IPChecker {
	return &IPChecker{source: cidrs}
}

func (c IPChecker) IsInList(ip string) (bool, string, error) {
	if result := net.ParseIP(ip); result == nil {
		return false, "", ErrIPNotValid
	}

	for el := range c.source {
		ips, err := c.getIpsFromCIDR(el.CIDR)

		if err != nil {
			return false, el.Type, nil
		}

		for _, listIp := range ips {
			if listIp == ip {
				return true, el.Type, nil
			}
		}

	}

	return false, "", nil
}

func (c IPChecker) getIpsFromCIDR(cidr string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)

	if err != nil {
		return nil, err
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); c.inc(ip) {
		ips = append(ips, ip.String())
	}

	return ips[1 : len(ips)-1], nil
}

func (c IPChecker) inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
