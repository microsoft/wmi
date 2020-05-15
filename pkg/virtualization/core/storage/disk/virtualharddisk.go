// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package disk

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/constant"
	"github.com/microsoft/wmi/pkg/errors"
	"github.com/microsoft/wmi/pkg/virtualization/core/resource/resourceallocation"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	"github.com/microsoft/wmi/server2019/root/virtualization/v2"
)

type VirtualHardDisk struct {
	*v2.Msvm_StorageAllocationSettingData
}

// NewVirtualHardDisk
func NewVirtualHardDisk(instance *wmi.WmiInstance) (*VirtualHardDisk, error) {
	wmivm, err := v2.NewMsvm_StorageAllocationSettingDataEx1(instance)
	if err != nil {
		return nil, err
	}
	return &VirtualHardDisk{wmivm}, nil
}

func (vhd *VirtualHardDisk) GetDrive() (*resourceallocation.ResourceAllocationSettingData, error) {
	parent, err := vhd.GetPropertyParent()
	if err != nil {
		return nil, err
	}

	inst, err := instance.GetWmiInstanceFromPath(vhd.GetWmiHost(), string(constant.Virtualization), parent)
	return resourceallocation.NewResourceAllocationSettingData(inst)
}

func (vhd *VirtualHardDisk) GetPath() (string, error) {
	value, err := vhd.GetProperty("HostResource")
	if err != nil {
		return "", err
	}
	for _, hr := range value.([]interface{}) {
		return hr.(string), nil
	}
	return "", errors.Wrapf(errors.NotFound, "")
}
