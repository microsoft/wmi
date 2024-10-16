// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package disk

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/constant"
	"github.com/microsoft/wmi/pkg/errors"
	"github.com/microsoft/wmi/pkg/virtualization/core/resource/resourceallocation"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	v2 "github.com/microsoft/wmi/server2019/root/virtualization/v2"
)

type LogicalDisk struct {
	*v2.Msvm_StorageAllocationSettingData
}

func NewLogicalDisk(instance *wmi.WmiInstance) (*LogicalDisk, error) {
	wmivm, err := v2.NewMsvm_StorageAllocationSettingDataEx1(instance)
	if err != nil {
		return nil, err
	}
	return &LogicalDisk{wmivm}, nil
}

func (ld *LogicalDisk) GetDrive() (*resourceallocation.ResourceAllocationSettingData, error) {
	parent, err := ld.GetPropertyParent()
	if err != nil {
		return nil, err
	}

	inst, err := instance.GetWmiInstanceFromPath(ld.GetWmiHost(), string(constant.Virtualization), parent)
	if err != nil {
		return nil, err
	}
	return resourceallocation.NewResourceAllocationSettingData(inst)
}

func (ld *LogicalDisk) GetPath() (string, error) {
	value, err := ld.GetProperty("HostResource")
	if err != nil {
		return "", err
	}
	for _, hr := range value.([]interface{}) {
		return hr.(string), nil
	}
	return "", errors.Wrapf(errors.NotFound, "")
}
