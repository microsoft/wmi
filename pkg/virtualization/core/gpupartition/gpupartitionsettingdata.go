// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package gpupartition

import (
	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/constant"
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
