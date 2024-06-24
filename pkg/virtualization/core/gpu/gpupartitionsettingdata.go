// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package gpu

import (
	"reflect"
	"strconv"

	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/constant"
	"github.com/microsoft/wmi/pkg/errors"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	v2 "github.com/microsoft/wmi/server2019/root/virtualization/v2"
)

type GpuPartitionSettingData struct {
	*v2.Msvm_GpuPartitionSettingData
}

// NewGpuPartitionSettingData
func NewGpuPartitionSettingData(instance *wmi.WmiInstance) (*GpuPartitionSettingData, error) {
	wmivm, err := v2.NewMsvm_GpuPartitionSettingDataEx1(instance)
	if err != nil {
		return nil, err
	}
	return &GpuPartitionSettingData{wmivm}, nil
}

func GetDefaultGpuPartitionSettingData(whost *host.WmiHost) (*GpuPartitionSettingData, error) {
	creds := whost.GetCredential()
	wmivm, err := v2.NewMsvm_GpuPartitionSettingDataEx6(whost.HostName, string(constant.Virtualization), creds.UserName, creds.Password, creds.Domain, query.NewWmiQuery("Msvm_GpuPartitionSettingData"))
	if err != nil {
		return nil, err
	}
	return &GpuPartitionSettingData{wmivm}, nil
}

func (partitionSettingData *GpuPartitionSettingData) CloneEx1() (*GpuPartitionSettingData, error) {
	tmp, err := partitionSettingData.Clone()
	if err != nil {
		return nil, err
	}
	return NewGpuPartitionSettingData(tmp)
}

func (partitionSettingData *GpuPartitionSettingData) GetMinPartitionVRAM() (uint64, error) {
	value, err := partitionSettingData.GetProperty("MinPartitionVRAM")
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
