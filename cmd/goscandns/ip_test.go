package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHostsFromCidr(t *testing.T) {
	{
		cidr := "1.1.1.0/24"
		ips, err := hostsFromCIDR(cidr)
		assert.NoError(t, err)
		assert.Len(t, ips, 254)
	}
	{
		cidr := "1.1.1.100/24"
		ips, err := hostsFromCIDR(cidr)
		assert.NoError(t, err)
		assert.Len(t, ips, 254)
	}
	{
		cidr := "1.1.1.254/24"
		ips, err := hostsFromCIDR(cidr)
		assert.NoError(t, err)
		assert.Len(t, ips, 254)
	}
	{
		cidr := "1.1.1/24"
		_, err := hostsFromCIDR(cidr)
		assert.Error(t, err)
	}
}
