// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package gpu

import (
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
)

type PartitionableGpuCollection []*PartitionableGpu

func NewPartitionableGpuCollection(instances []*wmi.WmiInstance) (col PartitionableGpuCollection, err error) {
	for _, inst := range instances {
		na, err1 := NewPartitionableGpu(inst)
		if err1 != nil {
			err = err1
			return
		}
		col = append(col, na)
	}
	return
}

func (gpuCollection *PartitionableGpuCollection) Close() (err error) {
	for _, value := range *gpuCollection {
		value.Close()
	}
	return nil
}
