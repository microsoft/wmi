// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package gpu

import (
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
)

type GpuPartitionSettingCollection []*GpuPartitionSettingData

func NewGpuPartitionSettingCollection(instances []*wmi.WmiInstance) (col GpuPartitionSettingCollection, err error) {
	for _, inst := range instances {
		na, err1 := NewGpuPartitionSettingData(inst)
		if err1 != nil {
			err = err1
			return
		}
		col = append(col, na)
	}
	return
}

func (gpuPartitionSettingCollection *GpuPartitionSettingCollection) Close() (err error) {
	for _, value := range *gpuPartitionSettingCollection {
		value.Close()
	}
	return nil
}

func (gpuPartitionSettingCollection *GpuPartitionSettingCollection) String() string {
	return ""
}
