// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package pnp

import (
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
)

type PnpEntity struct {
	*PnpEntityData
}

func NewPnpEntityEx1(instance *wmi.WmiInstance) (*PnpEntity, error) {
	wmivm, err := NewPnpEntityData(instance)
	if err != nil {
		return nil, err
	}
	return &PnpEntity{wmivm}, nil
}

func (pnpEntity *PnpEntity) CloneEx1() (*PnpEntity, error) {
	tmp, err := pnpEntity.Clone()
	if err != nil {
		return nil, err
	}
	return NewPnpEntityEx1(tmp)
}

// GetVirtualMachine gets the VM that the pnp entity is attached to
func (pnpEntity *PnpEntity) GetVirtualMachine() (instance *wmi.WmiInstance, err error) {
	vmsetting, err := pnpEntity.GetRelated("Msvm_VirtualSystemSettingData")
	if err != nil {
		return
	}
	defer vmsetting.Close()
	return vmsetting.GetRelated("Msvm_ComputerSystem")
}
