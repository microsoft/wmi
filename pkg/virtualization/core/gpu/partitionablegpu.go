// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package gpu

import (
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	"github.com/microsoft/wmi/server2019/root/virtualization/v2"
)

type PartitionableGpu struct {
	*v2.Msvm_PartitionableGpu
}

// NewPartitionableGpu
func NewPartitionableGpu(instance *wmi.WmiInstance) (*PartitionableGpu, error) {
	wmivm, err := v2.NewMsvm_PartitionableGpuEx1(instance)
	if err != nil {
		return nil, err
	}
	return &PartitionableGpu{wmivm}, nil
}

func (gpu *PartitionableGpu) CloneEx1() (*PartitionableGpu, error) {
	tmp, err := gpu.Clone()
	if err != nil {
		return nil, err
	}
	return NewPartitionableGpu(tmp)
}
