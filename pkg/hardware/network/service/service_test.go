// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package service

import (
	wmihost "github.com/microsoft/wmi/pkg/base/host"
	_ "github.com/microsoft/wmi/pkg/base/session"
	"testing"
)

func TestGetNetworkAdapter(t *testing.T) {
	whost := wmihost.NewWmiLocalHost()
	adapter, err := GetNetworkAdapterByName(whost, "Ethernet")
	if err != nil {
		t.Fatalf("[%+v]", err)
	}
	defer adapter.Close()

	ifIndex, err := adapter.GetInterfaceIndex()
	if err != nil {
		t.Fatalf("[%+v]", err)
	}

	t.Logf("Interface Index [%+v]\n", ifIndex)
	adapter1, err := GetNetworkAdapterByInterfaceIndex(whost, ifIndex)
	if err != nil {
		t.Fatalf("[%+v]", err)
	}
	defer adapter1.Close()
}

func TestFindFirstPhysicalAdapterWithDefaultRoute(t *testing.T) {
	whost := wmihost.NewWmiLocalHost()
	adapter, err := FindFirstPhysicalAdapterWithDefaultRoute(whost)
	if err != nil {
		t.Fatalf("[%+v]", err)
	}
	defer adapter.Close()

	ifIndex, err := adapter.GetInterfaceIndex()
	if err != nil {
		t.Fatalf("[%+v]", err)
	}
	adapterName, err := adapter.GetPropertyName()
	if err != nil {
		t.Fatalf("[%+v]", err)
	}

	t.Logf("Interface Index [%+v]  [%+v]\n", ifIndex, adapterName)

}

func TestGetNetworkRoute(t *testing.T) {
	whost := wmihost.NewWmiLocalHost()
	netroutes, err := GetNetworkRouteByDestinationAddress(whost, "0.0.0.0/0")
	if err != nil {
		t.Fatalf("[%+v]", err)
	}

	defer netroutes.Close()

	for _, netroute := range netroutes {
		ifIndex, err := netroute.GetInterfaceIndex()
		if err != nil {
			t.Fatalf("[%+v]", err)
		}
		t.Logf("Interface Index [%+v]\n", ifIndex)
	}

}
