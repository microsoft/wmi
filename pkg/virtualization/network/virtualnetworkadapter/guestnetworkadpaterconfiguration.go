// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package virtualnetworkadapter

import (
	"log"
	"net"
	"strings"

	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	v2 "github.com/microsoft/wmi/server2019/root/virtualization/v2"
)

type GuestNetworkAdapterConfiguration struct {
	*v2.Msvm_GuestNetworkAdapterConfiguration
}

// NewGuestNetworkAdapterConfiguration
func NewGuestNetworkAdapterConfiguration(instance *wmi.WmiInstance) (*GuestNetworkAdapterConfiguration, error) {
	wmivm, err := v2.NewMsvm_GuestNetworkAdapterConfigurationEx1(instance)
	if err != nil {
		return nil, err
	}
	return &GuestNetworkAdapterConfiguration{wmivm}, nil
}

func (gnac *GuestNetworkAdapterConfiguration) GetIPAddresses() (ipv4s, ipv6s []string, err error) {
	defer func() {
		log.Printf("[WMI] [GuestNetworkAdapterConfiguration] IPv4[%+v] IPv6[%+v] Error[%+v]\n", ipv4s, ipv6s, err)
	}()
	retValue, err := gnac.GetProperty("IPAddresses")
	if err != nil {
		return
	}
	//addresses := strings.Split(retValue.([]string), " ")
	for _, add := range retValue.([]interface{}) {
		address := add.(string)
		// Validate the IP
		if net.ParseIP(address) == nil {
			continue
		}
		if strings.Contains(address, ":") {
			ipv6s = append(ipv6s, address)
		} else if strings.Contains(address, ".") {
			ipv4s = append(ipv4s, address)
		}
	}

	return
}
