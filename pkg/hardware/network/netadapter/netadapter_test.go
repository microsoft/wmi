// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package netadapter

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

	ifdesc, err := adapter.GetPropertyInterfaceDescription()
	if err != nil {
		t.Fatalf("[%+v]", err)
	}

	t.Logf("Interface Description [%+v]\n", ifdesc)
}
