// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package service

import (
	"fmt"

	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/constant"
	"github.com/microsoft/wmi/pkg/errors"

	"github.com/microsoft/wmi/pkg/hardware/network/netadapter"
	"github.com/microsoft/wmi/pkg/hardware/network/netroute"
	"github.com/microsoft/wmi/server2019/root/standardcimv2"
)

// GetNetworkRouteByDestinationAddress
func GetNetworkRouteByDestinationAddress(whost *host.WmiHost, destinationPrefix string) (netroutes netroute.NetworkRouteCollection, err error) {
	query := query.NewWmiQuery("MSFT_NetRoute", "DestinationPrefix", destinationPrefix)
	instances, err := instance.GetWmiInstancesFromHost(whost, string(constant.StadardCimV2), query)
	if err != nil {
		return nil, err
	}

	return netroute.NewNetworkRouteCollection(instances)

}

// GetNetworkAdapterByName
func GetNetworkAdapterByName(whost *host.WmiHost, name string) (adapter *netadapter.NetworkAdapter, err error) {
	creds := whost.GetCredential()
	querytmp := query.NewWmiQuery("MSFT_NetAdapter", "Name", name)
	tmp, err := standardcimv2.NewMSFT_NetAdapterEx6(whost.HostName, string(constant.StadardCimV2), creds.UserName, creds.Password, creds.Domain, querytmp)
	if err != nil {
		return
	}
	return &netadapter.NetworkAdapter{tmp}, nil
}

// GetNetworkAdapterByName
func GetNetworkAdapterByInterfaceIndex(whost *host.WmiHost, ifindex int32) (adapter *netadapter.NetworkAdapter, err error) {
	creds := whost.GetCredential()
	querytmp := query.NewWmiQuery("MSFT_NetAdapter", "InterfaceIndex", fmt.Sprintf("%d", ifindex))
	tmp, err := standardcimv2.NewMSFT_NetAdapterEx6(whost.HostName, string(constant.StadardCimV2), creds.UserName, creds.Password, creds.Domain, querytmp)
	if err != nil {
		return
	}
	return &netadapter.NetworkAdapter{tmp}, nil
}

// FindDefaultExternalAdapter - finds the first physical adapter which has a default gateway
func FindFirstPhysicalAdapterWithDefaultRoute(whost *host.WmiHost) (adapter *netadapter.NetworkAdapter, err error) {
	routes, err := GetNetworkRouteByDestinationAddress(whost, "0.0.0.0/0")
	if err != nil {
		return
	}
	defer routes.Close()

	for _, route := range routes {
		ifindex, err1 := route.GetInterfaceIndex()
		if err1 != nil {
			err = err1
			return
		}
		adapter1, err1 := GetNetworkAdapterByInterfaceIndex(whost, ifindex)
		if err1 != nil {
			err = err1
			return
		}
		defer adapter1.Close()

		isvirtual, err1 := adapter1.GetPropertyVirtual()
		if err1 != nil {
			err = err1
			return
		}

		if isvirtual {
			continue
		}

		adapter, err = adapter1.CloneEx1()
		return
	}

	err = errors.Wrapf(errors.NotFound, "No suitable adapter was found with default route")
	return
}
