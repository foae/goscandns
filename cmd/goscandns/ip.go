package main

import (
	"context"
	"fmt"
	"net"
	"time"
)

func scanIP(ctx context.Context, ip net.IP) ([]net.IPAddr, error) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	res := net.Resolver{
		StrictErrors: false,
		PreferGo:     true,
		Dial: func(ctx context.Context, network, address string) (conn net.Conn, e error) {
			address = net.JoinHostPort(ip.String(), "53")
			d := net.Dialer{Deadline: time.Now().Add(time.Second * 2)}

			return d.DialContext(ctx, "udp4", address)
			//return net.DialTimeout("udp4", fmt.Sprintf("%v:53", ip.String()), time.Second*2)
		},
	}

	ips, err := res.LookupIPAddr(ctx, "www.arbor-observatory.com")
	if err != nil {
		return nil, fmt.Errorf("scanIP: %v", err)
	}

	return ips, nil
}

func hostsFromCIDR(cidr string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}

	return ips[1 : len(ips)-1], nil
}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
