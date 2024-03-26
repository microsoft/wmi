// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package gpupartition

import (
	"github.com/microsoft/wmi/pkg/errors"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
)

type GpuPartition struct {
	*GpuPartitionSettingData
}

func NewGpuPartition(instance *wmi.WmiInstance) (*GpuPartition, error) {
	wmivm, err := NewGpuPartitionSettingData(instance)
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

func (partition *GpuPartition) GetMinPartitionVRAM() (uint64, error) {
	value, err := partition.GetProperty("MinPartitionVRAM")
	if err != nil {
		return 0, err
	}
	for _, hr := range value.([]interface{}) {
		return hr.(uint64), nil
	}
	return 0, errors.Wrapf(errors.NotFound, "Unable to get minimum partition VRAM for given GPU partition [%s]", partition)
}
