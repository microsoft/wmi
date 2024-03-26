// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package gpupartition

import (
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
)

type GpuPartitionCollection []*GpuPartition

func NewGpuPartitionCollection(instances []*wmi.WmiInstance) (col GpuPartitionCollection, err error) {
	for _, inst := range instances {
		na, err1 := NewGpuPartition(inst)
		if err1 != nil {
			err = err1
			return
		}
		col = append(col, na)
	}
	return
}

func (vms *GpuPartitionCollection) Close() (err error) {
	for _, value := range *vms {
		value.Close()
	}
	return nil
}

func (vms *GpuPartitionCollection) String() string {
	return ""
}
