// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package processor

import (
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	"github.com/microsoft/wmi/server2019/root/virtualization/v2"
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
