// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package disk

import (
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	"github.com/microsoft/wmi/server2019/root/virtualization/v2"
)

type LogicalDisk struct {
	*v2.Msvm_StorageAllocationSettingData
}

func NewLogicalDisk(instance *wmi.WmiInstance) (*LogicalDisk, error) {
	wmivm, err := v2.NewMsvm_StorageAllocationSettingDataEx1(instance)
	if err != nil {
		return nil, err
	}
	return &LogicalDisk{wmivm}, nil
}