// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package gpu

import (
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	"github.com/microsoft/wmi/server2019/root/virtualization/v2"
)

type GpuPartition struct {
	*v2.Msvm_GpuPartition
}

// NewGpuPartition
func NewGpuPartition(instance *wmi.WmiInstance) (*GpuPartition, error) {
	wmivm, err := v2.NewMsvm_GpuPartitionEx1(instance)
	if err != nil {
		return nil, err
	}
	return &GpuPartition{wmivm}, nil
}

func (partition *GpuPartition) CloneEx1() (*GpuPartition, error) {
	tmp, err := partition.Clone()
	if err != nil {
		return nil, err
	}
	return NewGpuPartition(tmp)
}
