// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package resource

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

/*
	func TestGetResource(t *testing.T) {
		cn, err := GetResource(whost, "*")
		if err != nil {
			t.Fatal("Failed " + err.Error())
			return
		}
		defer cn.Close()
	}
*/

func TestGetResources(t *testing.T) {
	nc, err := GetResources(whost)
	if err != nil {
		t.Fatal("Failed " + err.Error())
		return
	}
	defer nc.Close()
	t.Logf("Nodes returned %d\n", len(nc))

	for _, resource := range nc {
		grpSetName, err := resource.GetPropertyName()
		if err != nil {
			t.Fatal("Failed " + err.Error())
			return
		}
		t.Logf("Resource Name : %s\n", grpSetName)

		state, err := resource.State()
		if err != nil {
			t.Fatal("Failed " + err.Error())
			return
		}
		t.Logf("Resource States %v \n", state)

		owners, err := resource.GetPossibleOwnersEx1()
		if err != nil {
			t.Fatal("Failed " + err.Error())
			return
		}
		t.Logf("Resource Owners %v \n", owners)
	}
}
