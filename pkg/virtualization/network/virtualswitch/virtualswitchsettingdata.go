// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package virtualswitch

import (
	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/constant"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	"github.com/microsoft/wmi/server2019/root/virtualization/v2"
)

type VirtualEthernetSwitchSettingData struct {
	*v2.Msvm_VirtualEthernetSwitchSettingData
}

// NewVirtualEthernetSwitchSettingData
func NewVirtualEthernetSwitchSettingData(instance *wmi.WmiInstance) (*VirtualEthernetSwitchSettingData, error) {
	wmientity, err := v2.NewMsvm_VirtualEthernetSwitchSettingDataEx1(instance)
	if err != nil {
		return nil, err
	}
	return &VirtualEthernetSwitchSettingData{wmientity}, nil
}

func GetDefaultVirtualEthernetSwitchSettingData(whost *host.WmiHost) (*VirtualEthernetSwitchSettingData, error) {
	inst, err := instance.CreateWmiInstance(whost, string(constant.Virtualization), "Msvm_VirtualEthernetSwitchSettingData")
	if err != nil {
		return nil, err
	}
	return NewVirtualEthernetSwitchSettingData(inst)
}

func GetVirtualEthernetSwitchSettingData(whost *host.WmiHost, name string) (*VirtualEthernetSwitchSettingData, error) {
	vmsettings, err := GetDefaultVirtualEthernetSwitchSettingData(whost)
	if err != nil {
		return nil, err
	}
	vmsettings.SetPropertyElementName(name)
	// vmsettings.SetPropertyVirtualSystemSubType
	// vmsettings.SetProperty("VirtualSystemSubType", 1) // 2nd Generation

	return vmsettings, err
}

func (vessd *VirtualEthernetSwitchSettingData) getResourceAllocationSettingData() (col wmi.WmiInstanceCollection, err error) {
	query := query.NewWmiQuery("Cim_ResourceAllocationSettingData")
	col, err = vessd.GetAllRelatedWithQuery(query)
	return
}
