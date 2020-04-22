// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package virtualsystem

import (
	"fmt"
	"net"

	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	vhd "github.com/microsoft/wmi/pkg/virtualization/core/storage/virtualharddisk"
	na "github.com/microsoft/wmi/pkg/virtualization/network/virtualnetworkadapter"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	"github.com/microsoft/wmi/server2019/root/virtualization/v2"
)

type VirtualMachineSetting struct {
	*v2.Msvm_VirtualSystemSettingData
}

// NewVirtualMachineSetting
func NewVirtualMachineSetting(instance *wmi.WmiInstance) (*VirtualMachineSetting, error) {
	wmivm, err := v2.NewMsvm_VirtualSystemSettingDataEx1(instance)
	if err != nil {
		return nil, err
	}
	return &VirtualMachineSetting{wmivm}, nil
}

func (vm *VirtualMachineSetting) GetSyntheticVirtualNetworkAdapters() (col na.VirtualNetworkAdapterCollection, err error) {
	psds, err := vm.GetAllRelated("Msvm_SyntheticEthernetPortSettingData")
	if err != nil {
		return nil, err
	}

	col, err = na.NewVirtualNetworkAdapterCollection(psds)
	if err != nil {
		psds.Close()
	}
	return
}

func (vm *VirtualMachineSetting) GetEmulatedVirtualNetworkAdapters() (col na.VirtualNetworkAdapterCollection, err error) {
	psds, err := vm.GetAllRelated("Msvm_EmulatedEthernetPortSettingData")
	if err != nil {
		return nil, err
	}

	col, err = na.NewVirtualNetworkAdapterCollection(psds)
	if err != nil {
		psds.Close()
	}
	return
}

func (vm *VirtualMachineSetting) GetVirtualNetworkAdapters() (col na.VirtualNetworkAdapterCollection, err error) {
	psds, err := vm.GetSyntheticVirtualNetworkAdapters()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			psds.Close()
		}
	}()
	psdse, err := vm.GetEmulatedVirtualNetworkAdapters()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			psdse.Close()
		}
	}()

	col = append(col, psds...)
	col = append(col, psdse...)
	return
}

// GetVirtualNetworkAdapterByMACAddress
func (vm *VirtualMachineSetting) GetVirtualNetworkAdapterByMACAddress(macAddress string) (vna *na.VirtualNetworkAdapter, err error) {
	hw, err := net.ParseMAC(macAddress)
	if err != nil {
		err = errors.Wrapf(errors.InvalidInput, "MacAddress")
		return
	}
	nas, err := vm.GetVirtualNetworkAdapters()
	if err != nil {
		return
	}

	defer nas.Close()

	for _, networkAdapter := range nas {
		physicalAddress, err1 := networkAdapter.GetPropertyAddress()
		if err1 != nil {
			err = err1
			return
		}
		hwcurrent, err1 := net.ParseMAC(physicalAddress)
		if err1 != nil {
			err = err1
			return
		}
		if hw.String() == hwcurrent.String() {
			// Found the match
			// Assumption is only one adapter would match, as MAC is unique
			inst, err1 := networkAdapter.Clone()
			if err1 != nil {
				err = err1
				return
			}
			vna, err = na.NewVirtualNetworkAdapter(inst)
			return
		}
	}
	err = errors.Wrapf(errors.NotFound, "Virtual Network Adapter with macaddress [%s]", macAddress)
	return
}

// GetVirtualNetworkAdapterByName
func (vm *VirtualMachineSetting) GetVirtualNetworkAdapterByName(adapterName string) (vna *na.VirtualNetworkAdapter, err error) {
	nas, err := vm.GetVirtualNetworkAdapters()
	if err != nil {
		return
	}

	defer nas.Close()

	for _, networkAdapter := range nas {
		name, err1 := networkAdapter.GetPropertyElementName()
		if err1 != nil {
			err = err1
			return
		}
		if name == adapterName {
			// Found the match
			// Assumption is only one adapter would match, - FIXME - Duplicates
			inst, err1 := networkAdapter.Clone()
			if err1 != nil {
				err = err1
				return
			}
			vna, err = na.NewVirtualNetworkAdapter(inst)
			return
		}
	}
	err = errors.Wrapf(errors.NotFound, "Virtual Network Adapter with name [%s]", adapterName)
	return
}

func (vm *VirtualMachineSetting) GetVirtualHardDisks() (col vhd.VirtualHardDiskCollection, err error) {
	rasdcollection, err := vm.getResourceAllocationSettingData(v2.ResourceAllocationSettingData_ResourceType_Disk_Drive)
	if err != nil {
		return nil, err
	}

	col, err = vhd.NewVirtualHardDiskCollection(rasdcollection)
	if err != nil {
		rasdcollection.Close()
	}
	return
}

func (vm *VirtualMachineSetting) getResourceAllocationSettingData(rtype v2.ResourceAllocationSettingData_ResourceType) (col wmi.WmiInstanceCollection, err error) {
	resourceType := fmt.Sprintf("%d", int32(rtype))
	query := query.NewWmiQuery("Cim_ResourceAllocationSettingData", "ResourceType", resourceType)
	col, err = vm.GetAllRelatedWithQuery(query)
	return
}
