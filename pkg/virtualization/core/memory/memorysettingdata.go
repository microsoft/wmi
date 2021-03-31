// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package memory

import (
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	"github.com/microsoft/wmi/server2019/root/virtualization/v2"
	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/constant"
)

type MemorySettingData struct {
	*v2.Msvm_MemorySettingData
}

// NewMemorySettingData
func NewMemorySettingData(instance *wmi.WmiInstance) (*MemorySettingData, error) {
	wmivm, err := v2.NewMsvm_MemorySettingDataEx1(instance)
	if err != nil {
		return nil, err
	}
	return &MemorySettingData{wmivm}, nil
}

func (msd *MemorySettingData) SetSizeMB(size uint64) (err error) {
	return msd.SetPropertyVirtualQuantity(size)
}

func GetDefaultMemorySettingData(whost *host.WmiHost) (*MemorySettingData, error) {
	inst, err := instance.CreateWmiInstance(whost, string(constant.Virtualization), "Msvm_MemorySettingData")
	if err != nil {
		return nil, err
	}
	return NewMemorySettingData(inst)
}