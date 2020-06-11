// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package virtualswitch

import (
	"fmt"

	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/constant"
	"github.com/microsoft/wmi/pkg/errors"
	"github.com/microsoft/wmi/pkg/virtualization/network/switchextension"
	"github.com/microsoft/wmi/pkg/virtualization/network/switchport"
	na "github.com/microsoft/wmi/pkg/virtualization/network/virtualnetworkadapter"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	v2 "github.com/microsoft/wmi/server2019/root/virtualization/v2"
)

type VirtualSwitch struct {
	*v2.Msvm_VirtualEthernetSwitch
}

// NewVirtualSwitch
func NewVirtualSwitch(instance *wmi.WmiInstance) (*VirtualSwitch, error) {
	wmivm, err := v2.NewMsvm_VirtualEthernetSwitchEx1(instance)
	if err != nil {
		return nil, err
	}
	return &VirtualSwitch{wmivm}, nil
}

// GetVirtualSwitch gets an existing virtual machine
// Make sure to call Close once done using this instance
func GetVirtualSwitch(whost *host.WmiHost, vswitchName string) (vswitch *VirtualSwitch, err error) {
	creds := whost.GetCredential()
	query := query.NewWmiQuery("Msvm_VirtualEthernetSwitch", "ElementName", vswitchName)
	wmivm, err := v2.NewMsvm_VirtualEthernetSwitchEx6(whost.HostName, string(constant.Virtualization), creds.UserName, creds.Password, creds.Domain, query)
	if err != nil {
		return
	}
	vswitch = &VirtualSwitch{wmivm}
	return
}

func (vs *VirtualSwitch) GetVirtualMachineAdapterByName(name string) (vadapter *na.VirtualNetworkAdapter, err error) {
	adapters, err := vs.GetVirtualMachineAdapters()
	if err != nil {
		return
	}
	defer adapters.Close()
	fmt.Printf("VM Adapters [%d]", len(adapters))

	for _, adapter := range adapters {
		curAdapterName, err1 := adapter.GetPropertyElementName()
		if err1 != nil {
			err = err1
			return
		}

		if curAdapterName == name {
			clonedAdapter, err1 := adapter.Clone()
			if err1 != nil {
				err = err1
				return
			}

			return na.NewVirtualNetworkAdapter(clonedAdapter)
		}
	}
	err = errors.Wrapf(errors.NotFound, "VirtualNetworkAdapter [%s]", name)
	return
}

func (vs *VirtualSwitch) GetVirtualMachineAdapters() (col na.VirtualNetworkAdapterCollection, err error) {
	switchPorts, err := vs.getEthernetSwitchPorts()
	if err != nil {
		return nil, err
	}
	defer switchPorts.Close()
	fmt.Printf("SwitchPorts [%d]", len(switchPorts))

	for _, switchPort := range switchPorts {
		pasd, err1 := switchPort.GetEthernetPortAllocationSettingData()
		if err1 != nil {
			err = err1
			return
		}
		defer pasd.Close()
		retValue, err1 := pasd.GetProperty("MappingBehavior")
		if err1 != nil {
			err = err1
			return
		}

		if retValue == nil || retValue.(int32) != int32(v2.ResourceAllocationSettingData_MappingBehavior_Dedicated) {
			nic, err1 := pasd.GetRelated("Msvm_SyntheticEthernetPortSettingData")
			if err1 != nil {
				err = err1
				return
			}
			na, err1 := na.NewVirtualNetworkAdapter(nic)
			if err1 != nil {
				err = err1
				return
			}
			col = append(col, na)
		}
	}
	return
}

func (vs *VirtualSwitch) getEthernetSwitchPorts() (col switchport.EthernetSwitchPortCollection, err error) {
	instances, err := vs.GetAllRelated("Msvm_EthernetSwitchPort")
	//	"Msvm_SystemDevice",
	//	"PartComponent", "GroupComponent")
	if err != nil {
		return nil, err
	}

	col, err = switchport.NewEthernetSwitchPortCollection(instances)
	if err != nil {
		instances.Close()
	}

	return
}

func (vs *VirtualSwitch) GetEthernetSwitchExtensions() (col switchextension.EthernetSwitchExtensionCollection, err error) {
	instances, err := vs.GetAllRelated("Msvm_EthernetSwitchExtension")
	if err != nil {
		return nil, err
	}
	col, err = switchextension.NewEthernetSwitchExtensionCollection(instances)
	if err != nil {
		instances.Close()
	}
	return
}

func (vs *VirtualSwitch) GetEthernetSwitchExtensionByName(name string) (se *switchextension.EthernetSwitchExtension, err error) {
	exts, err := vs.GetEthernetSwitchExtensions()
	if err != nil {
		return
	}
	defer exts.Close()

	for _, ext := range exts {
		extName, err1 := ext.GetPropertyElementName()
		if err1 != nil {
			err = err1
			return
		}

		if extName != name {
			continue
		}

		return ext.CloneEx1()
	}

	err = errors.Wrapf(errors.NotFound, "SwitchExtension [%s]", name)
	return
}
