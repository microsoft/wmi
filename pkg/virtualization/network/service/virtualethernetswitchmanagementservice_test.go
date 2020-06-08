// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package service

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

func TestGetVirtualEthernetSwitchManagementService(t *testing.T) {
	vess, err := GetVirtualEthernetSwitchManagementService(whost)
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
	if vess == nil {
		t.Fatal("Failed Unable to get Switch Service")
	}
}

func TestGetVirtualSwitches(t *testing.T) {
	vmms, err := GetVirtualEthernetSwitchManagementService(whost)
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}

	vms, err := vmms.GetVirtualSwitches()
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
	vms.Close()
}

func TestVirtualSwitchDelete(t *testing.T) {
	vmms, err := GetVirtualEthernetSwitchManagementService(whost)
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
	vm, err := vmms.FindVirtualSwitchByName("test")
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}

	if vm == nil {
		t.Fatal("Failed to get VM")
	}
	defer vm.Close()
	err = vmms.DeleteVirtualSwitch(vm)
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
}
