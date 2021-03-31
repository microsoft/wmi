// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package processor

import (
	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/constant"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	v2 "github.com/microsoft/wmi/server2019/root/virtualization/v2"
)

type ProcessorSettingData struct {
	*v2.Msvm_ProcessorSettingData
}

// NewProcessorSettingData
func NewProcessorSettingData(instance *wmi.WmiInstance) (*ProcessorSettingData, error) {
	wmivm, err := v2.NewMsvm_ProcessorSettingDataEx1(instance)
	if err != nil {
		return nil, err
	}
	return &ProcessorSettingData{wmivm}, nil
}

func (msd *ProcessorSettingData) SetCPUCount(count uint64) (err error) {
	return msd.SetPropertyVirtualQuantity(count)
}

func GetDefaultProcessorSettingData(whost *host.WmiHost) (*ProcessorSettingData, error) {
	inst, err := instance.CreateWmiInstance(whost, string(constant.Virtualization), "Msvm_ProcessorSettingData")
	if err != nil {
		return nil, err
	}
	return NewProcessorSettingData(inst)
}
