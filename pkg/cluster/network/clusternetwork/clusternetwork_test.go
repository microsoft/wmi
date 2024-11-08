// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package clusternetwork

import (
	"github.com/microsoft/wmi/pkg/base/host"
	_ "github.com/microsoft/wmi/pkg/base/session"
	"testing"
)

var (
	whost *host.WmiHost
)

func init() {
	whost = host.NewWmiLocalHost()
}

func TestGetClusterNetwork(t *testing.T) {
	cn, err := GetClusterNetwork(whost, "Cluster Network 1")
	if err != nil {
		t.Fatal("Failed " + err.Error())
		return
	}
	defer cn.Close()
}

func TestGetClusterNetworks(t *testing.T) {
	nc, err := GetClusterNetworks(whost)
	if err != nil {
		t.Fatal("Failed " + err.Error())
		return
	}
	defer nc.Close()
	t.Logf("networks returned %d\n", len(nc))

	for _, network := range nc {
		networkName, err := network.GetPropertyName()
		if err != nil {
			t.Fatal("Failed " + err.Error())
			return
		}
		t.Logf("Network Name : %s\n", networkName)
	}
}
