// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package gpu

import (
	"reflect"
	"strconv"

	"github.com/microsoft/wmi/pkg/errors"
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

func (gpu *PartitionableGpu) GetPartitionCount() (int32, error) {
	value, err := gpu.GetProperty("PartitionCount")
	if err != nil || value == nil {
		return 0, err
	}

	valuetmp, ok := value.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, "%s is invalid. Expected type is int32", reflect.TypeOf(value))
		return 0, err
	}

	return int32(valuetmp), nil
}

func (gpu *PartitionableGpu) GetMinPartitionVRAM() (uint64, error) {
	value, err := gpu.GetProperty("MinPartitionVRAM")
	if err != nil || value == nil {
		return 0, err
	}

	valuetmp, ok := value.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, "%s is invalid. Expected type is string", reflect.TypeOf(value))
		return 0, err
	}

	// base 10, 64 bit
	valueint, err := strconv.ParseUint(valuetmp, 10, 64)
	if err != nil {
		return 0, err
	}

	return valueint, nil
}
