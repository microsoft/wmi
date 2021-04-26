// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package service

import (
	"fmt"

	"github.com/microsoft/wmi/pkg/base/host"
	_ "github.com/microsoft/wmi/pkg/base/session"
	"github.com/microsoft/wmi/pkg/virtualization/network/virtualswitch"
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
	vsms, err := GetVirtualEthernetSwitchManagementService(whost)
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}

	entities, err := vsms.GetVirtualSwitches()
	if err != nil {
		t.Fatal("Failed " + err.Error())
	}
	entities.Close()
}

func TestExternalVirtualSwitchCreateDelete(t *testing.T) {
	vsms, err := GetVirtualEthernetSwitchManagementService(whost)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	setting, err := virtualswitch.GetVirtualEthernetSwitchSettingData(whost, "external")
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	defer setting.Close()
	vswitch, err := vsms.CreateExternalVirtualSwitch("Ethernet", "7dbb1038-caf0-401f-9a2c-4a72d4e8bc97", "7dbb1038-caf0-401f-9a2c-4a72d4e8bc98", setting)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	vswitch.Close()

	vswitch, err = vsms.FindVirtualSwitchByName("external")
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	defer vswitch.Close()
	err = vsms.DeleteVirtualSwitch(vswitch)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
}

func TestVirtualSwitchCreateDeleteMultiple(t *testing.T) {
	vsms, err := GetVirtualEthernetSwitchManagementService(whost)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	count := 5
	for i := 0; i < count; i++ {
		vswitchName := fmt.Sprintf("test-%d", i)
		setting, err := virtualswitch.GetVirtualEthernetSwitchSettingData(whost, vswitchName)
		if err != nil {
			t.Fatalf("Failed [%+v]", err)
		}
		defer setting.Close()
		vswitch, err := vsms.CreatePrivateVirtualSwitch(setting)
		if err != nil {
			t.Fatalf("Failed [%+v]", err)
		}
		vswitch.Close()
	}

	for i := 0; i < count; i++ {
		vswitchName := fmt.Sprintf("test-%d", i)
		vswitch, err := vsms.FindVirtualSwitchByName(vswitchName)
		if err != nil {
			t.Fatalf("Failed [%+v]", err)
		}
		defer vswitch.Close()
		err = vsms.DeleteVirtualSwitch(vswitch)
		if err != nil {
			t.Fatalf("Failed [%+v]", err)
		}
	}

}

func TestVirtualSwitchCreateDelete(t *testing.T) {
	vsms, err := GetVirtualEthernetSwitchManagementService(whost)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	setting, err := virtualswitch.GetVirtualEthernetSwitchSettingData(whost, "external")
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	defer setting.Close()
	vswitch, err := vsms.CreatePrivateVirtualSwitch(setting)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
	defer vswitch.Close()
	err = vsms.DeleteVirtualSwitch(vswitch)
	if err != nil {
		t.Fatalf("Failed [%+v]", err)
	}
}

type testContext struct {
}

func onCallback(ctx interface{}, data string) {
	fmt.Println("Modified :" + data)
}

func TestCreateVirtualSwitchMonitor(t *testing.T) {
	//return
	ctx := &testContext{}
	vnicMonior := virtualswitch.CreateVirtualSwitchMonitor(ctx, onCallback)
	defer vnicMonior.Close()
	vnicMonior.AddEntity("test")
	//for {
	//}
}
