// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package virtualnetworkadapter

import (
	//"fmt"
	"log"

	"github.com/microsoft/wmi/pkg/virtualization/network/switchport"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
)

type VirtualNetworkAdapter struct {
	*switchport.EthernetPortAllocationSettingData
}

// NewVirtualNetworkAdapter
func NewVirtualNetworkAdapter(instance *wmi.WmiInstance) (*VirtualNetworkAdapter, error) {
	wmivm, err := switchport.NewEthernetPortAllocationSettingData(instance)
	if err != nil {
		return nil, err
	}
	return &VirtualNetworkAdapter{wmivm}, nil
}

func (vna *VirtualNetworkAdapter) GetEthernetPortAllocationSettingData() (epas *switchport.EthernetPortAllocationSettingData, err error) {
	tmp, err := vna.GetRelated("Msvm_EthernetPortAllocationSettingData")
	if err != nil {
		return
	}
	return switchport.NewEthernetPortAllocationSettingData(tmp)
}

func (vna *VirtualNetworkAdapter) CloneEx1() (*VirtualNetworkAdapter, error) {
	tmp, err := vna.Clone()
	if err != nil {
		return nil, err
	}
	return NewVirtualNetworkAdapter(tmp)
}

func (vna *VirtualNetworkAdapter) SetMACAddress(mac string) (err error) {
	return vna.SetPropertyAddress(mac)
}

func (vna *VirtualNetworkAdapter) Rename(newName string) (err error) {
	return vna.SetPropertyElementName(newName)
}

func (vna *VirtualNetworkAdapter) GetIPv4Address() (ip string, err error) {
	defer func() {
		log.Printf("[WMI][VirtualNetworkAdapter] GetIPv4Address [%s]\n", ip)
	}()

	gnac, err := vna.GetGuestNetworkAdapterConfiguration()
	if err != nil {
		return
	}
	defer gnac.Close()
	v4s, _, err := gnac.GetIPAddresses()
	if err != nil {
		return
	}
	if len(v4s) == 0 {
		// Empty IPAddress - Valid
		return
	}

	ip = v4s[0]
	return
}

func (vna *VirtualNetworkAdapter) GetGuestNetworkAdapterInstanceID() (id string, err error) {
	gnac, err := vna.GetGuestNetworkAdapterConfiguration()
	if err != nil {
		return
	}
	defer gnac.Close()
	id, err = gnac.GetPropertyInstanceID()
	return
}
func (vna *VirtualNetworkAdapter) GetGuestNetworkAdapterConfiguration() (*GuestNetworkAdapterConfiguration, error) {
	inst, err := vna.GetRelated("Msvm_GuestNetworkAdapterConfiguration")
	if err != nil {
		return nil, err
	}
	return NewGuestNetworkAdapterConfiguration(inst)
}

// GetVirtualMachine gets the VM that the adapter is attached to
func (vna *VirtualNetworkAdapter) GetVirtualMachine() (instance *wmi.WmiInstance, err error) {
	vmsetting, err := vna.GetRelated("Msvm_VirtualSystemSettingData")
	if err != nil {
		return
	}
	defer vmsetting.Close()
	return vmsetting.GetRelated("Msvm_ComputerSystem")
}
