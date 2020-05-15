// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package drive

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/constant"
	"github.com/microsoft/wmi/pkg/virtualization/core/resource/resourceallocation"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
)

type VirtualDrive struct {
	*resourceallocation.ResourceAllocationSettingData
}

// NewVirtualDrive
func NewVirtualDrive(instance *wmi.WmiInstance) (*VirtualDrive, error) {
	wmivm, err := resourceallocation.NewResourceAllocationSettingData(instance)
	if err != nil {
		return nil, err
	}
	return &VirtualDrive{wmivm}, nil
}

func (vdrive *VirtualDrive) GetController() (*resourceallocation.ResourceAllocationSettingData, error) {
	parent, err := vdrive.GetPropertyParent()
	if err != nil {
		return nil, err
	}

	inst, err := instance.GetWmiInstanceFromPath(vdrive.GetWmiHost(), string(constant.Virtualization), parent)
	return resourceallocation.NewResourceAllocationSettingData(inst)
}

func (vdrive *VirtualDrive) GetControllerLocation() (string, error) {
	return vdrive.GetPropertyAddressOnParent()
}

func (vdrive *VirtualDrive) GetControllerNumber() (string, error) {
	controller, err := vdrive.GetController()
	if err != nil {
		return "0", err
	}
	val, err := controller.GetPropertyAddress()
	if err != nil {
		return "0", err
	}
	if len(val) == 0 {
		return "0", nil
	}
	return val, err
}
