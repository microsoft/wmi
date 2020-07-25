// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package memory

import (
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	"github.com/microsoft/wmi/server2019/root/virtualization/v2"
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
