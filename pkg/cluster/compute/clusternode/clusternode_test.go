// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package clusternode

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

func TestGetClusterNode(t *testing.T) {
	cn, err := GetLocalClusterNode(whost)
	if err != nil {
		t.Fatal("Failed " + err.Error())
		return
	}
	defer cn.Close()
}

func TestGetClusterNodes(t *testing.T) {
	nc, err := GetClusterNodes(whost)
	if err != nil {
		t.Fatal("Failed " + err.Error())
		return
	}
	defer nc.Close()
	t.Logf("Nodes returned %d\n", len(nc))

	for _, node := range nc {
		nodeName, err := node.GetPropertyName()
		if err != nil {
			t.Fatal("Failed " + err.Error())
			return
		}
		t.Logf("NoodeName : %s\n", nodeName)
	}
}
