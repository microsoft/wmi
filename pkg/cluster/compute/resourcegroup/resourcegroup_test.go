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
		t.Fatalf("Failed %+v\n ", err)
		return
	}
	defer nc.Close()
	t.Logf("Nodes returned %d\n", len(nc))

	for _, resourceGroup := range nc {
		grpSetName, err := resourceGroup.GetPropertyName()
		if err != nil {
			t.Fatalf("Failed %+v\n ", err)
			return
		}
		t.Logf("Resource Group : %s\n", grpSetName)

		state, err := resourceGroup.State()
		if err != nil {
			t.Fatalf("Failed %+v\n ", err)
			return
		}
		t.Logf("Resource States %v \n", state)

		pendingState, err := resourceGroup.IsPendingState()
		if err != nil {
			t.Fatalf("Failed %+v\n ", err)
			return
		}
		t.Logf("IsPending State %v \n", pendingState)
	}
}

func TestGetVirtualMachineResourceGroups(t *testing.T) {
	resourceGroups, err := GetVirtualMachineResourceGroups(whost)
	if err != nil {
		t.Fatalf("Failed %+v\n ", err)
		return
	}
	defer resourceGroups.Close()
	t.Logf("Resource Groups returned %d\n", len(resourceGroups))

	for _, resourceGroup := range resourceGroups {
		grpSetName, err := resourceGroup.GetPropertyName()
		if err != nil {
			t.Fatalf("Failed %+v\n ", err)
			return
		}
		t.Logf("Resource Group : %s\n", grpSetName)

		vmId, err := GetVirtualMachineID(whost, grpSetName)
		if err != nil {
			t.Fatalf("Failed %+v\n ", err)
			return
		}
		t.Logf("VMId : %s\n", vmId)

		state, err := resourceGroup.State()
		if err != nil {
			t.Fatalf("Failed %+v\n ", err)
			return
		}
		t.Logf("Resource States %v \n", state)

		pendingState, err := resourceGroup.IsPendingState()
		if err != nil {
			t.Fatalf("Failed %+v\n ", err)
			return
		}
		t.Logf("IsPending State %v \n", pendingState)

		pw, err := resourceGroup.GetPreferredOwnersEx1()
		if err != nil {
			t.Fatalf("Failed %+v\n ", err)
			return
		}
		t.Logf("Preferred Owners %v \n", pw)
		err = resourceGroup.SetVirtualMachinePriority(1000)
		if err != nil {
			t.Fatalf("Failed %+v\n ", err)
			return
		}
		curPriority, err := resourceGroup.GetVirtualMachinePriority()
		if err != nil {
			t.Fatalf("Failed %+v\n ", err)
			return
		}
		t.Logf("Priority Set : %v \n", curPriority)
		if curPriority != 1000 {
			t.Fatalf("Failed %+v\n ", "Unable to set priority")
			return
		}
	}
}

/*

func TestGetVirtualMachineResourceGroupByVmID(t *testing.T) {
	cn, err := GetVirtualMachineResourceGroupByVmID(whost, "Virtual Machine")
	if err != nil {
		t.Fatalf("Failed %+v\n ", err)
		return
	}
	defer cn.Close()
}

func TestSetVirtualMachineAutoFailback(t *testing.T) {
	cn, err := GetVirtualMachineResourceGroup(whost, "Virtual Machine")
	if err != nil {
		t.Fatalf("Failed %+v\n ", err)
		return
	}
	defer cn.Close()

	err = cn.SetVirtualMachineAutoFailback(true)
	if err != nil {
		t.Fatalf("Failed %+v\n ", err)
		return
	}
}

func TestDisableVirtualMachineDefaultOwner(t *testing.T) {
	cn, err := GetVirtualMachineResourceGroup(whost, "Virtual Machine")
	if err != nil {
		t.Fatalf("Failed %+v\n ", err)
		return
	}
	defer cn.Close()

	err = cn.DisableVirtualMachineDefaultOwner()
	if err != nil {
		t.Fatalf("Failed %+v\n ", err)
		return
	}
}

*/
