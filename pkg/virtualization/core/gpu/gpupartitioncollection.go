// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package gpu

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

func (partitions *GpuPartitionCollection) Close() (err error) {
	for _, value := range *partitions {
		value.Close()
	}
	return nil
}

func (partitions *GpuPartitionCollection) String() string {
	return ""
}
