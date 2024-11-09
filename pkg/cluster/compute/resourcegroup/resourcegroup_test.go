// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package resourcegroup

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

func TestGetResourceGroup(t *testing.T) {
	cn, err := GetResourceGroup(whost, "Cluster Group")
	if err != nil {
		t.Fatal("TestDisableVirtualMachineDefaultOwner failed: " + err.Error())
		return
	}
	defer cn.Close()
}

func TestGetResourceGroups(t *testing.T) {
	nc, err := GetResourceGroups(whost)
	if err != nil {
		t.Fatal("Failed " + err.Error())
		return
	}
	defer nc.Close()
	t.Logf("Nodes returned %d\n", len(nc))

	for _, resourceGroup := range nc {
		grpSetName, err := resourceGroup.GetPropertyName()
		if err != nil {
			t.Fatal("Failed " + err.Error())
			return
		}
		t.Logf("Resource Group : %s\n", grpSetName)

		state, err := resourceGroup.State()
		if err != nil {
			t.Fatal("Failed " + err.Error())
			return
		}
		t.Logf("Resource States %v \n", state)

		pendingState, err := resourceGroup.IsPendingState()
		if err != nil {
			t.Fatal("Failed " + err.Error())
			return
		}
		t.Logf("IsPending State %v \n", pendingState)
	}
}

func TestGetVirtualMachineResourceGroups(t *testing.T) {
	resourceGroups, err := GetVirtualMachineResourceGroups(whost)
	if err != nil {
		t.Fatal("Failed " + err.Error())
		return
	}
	defer resourceGroups.Close()
	t.Logf("Resource Groups returned %d\n", len(resourceGroups))

	for _, resourceGroup := range resourceGroups {
		grpSetName, err := resourceGroup.GetPropertyName()
		if err != nil {
			t.Fatal("Failed " + err.Error())
			return
		}
		t.Logf("Resource Group : %s\n", grpSetName)

		state, err := resourceGroup.State()
		if err != nil {
			t.Fatal("Failed " + err.Error())
			return
		}
		t.Logf("Resource States %v \n", state)

		pendingState, err := resourceGroup.IsPendingState()
		if err != nil {
			t.Fatal("Failed " + err.Error())
			return
		}
		t.Logf("IsPending State %v \n", pendingState)

		cn, err := GetVirtualMachineResourceGroupViaAssociators(whost, grpSetName)
		if err != nil {
			t.Fatal("Failed " + err.Error())
			return
		}
		t.Logf("Resource Group : %v\n", cn)
		t.Logf("Resource Group VMId : %s\n", cn.VmId)
		defer cn.Close()
	}
}

/*

func TestGetVirtualMachineResourceGroup(t *testing.T) {
	cn, err := GetVirtualMachineResourceGroup(whost, "Virtual Machine")
	if err != nil {
		t.Fatal("Failed " + err.Error())
		return
	}
	defer cn.Close()
}

func TestGetVirtualMachineResourceGroupByVmID(t *testing.T) {
	cn, err := GetVirtualMachineResourceGroupByVmID(whost, "Virtual Machine")
	if err != nil {
		t.Fatal("Failed " + err.Error())
		return
	}
	defer cn.Close()
}

func TestSetVirtualMachineAutoFailback(t *testing.T) {
	cn, err := GetVirtualMachineResourceGroup(whost, "Virtual Machine")
	if err != nil {
		t.Fatal("Failed " + err.Error())
		return
	}
	defer cn.Close()

	err = cn.SetVirtualMachineAutoFailback(true)
	if err != nil {
		t.Fatal("Failed " + err.Error())
		return
	}
}

func TestDisableVirtualMachineDefaultOwner(t *testing.T) {
	cn, err := GetVirtualMachineResourceGroup(whost, "Virtual Machine")
	if err != nil {
		t.Fatal("Failed " + err.Error())
		return
	}
	defer cn.Close()

	err = cn.DisableVirtualMachineDefaultOwner()
	if err != nil {
		t.Fatal("Failed " + err.Error())
		return
	}
}

*/