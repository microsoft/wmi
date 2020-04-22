// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package virtualharddisk

import (
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	"github.com/microsoft/wmi/server2019/root/virtualization/v2"
)

type VirtualHardDisk struct {
	*v2.Msvm_StorageAllocationSettingData
}

// NewVirtualHardDisk
func NewVirtualHardDisk(instance *wmi.WmiInstance) (*VirtualHardDisk, error) {
	wmivm, err := v2.NewMsvm_StorageAllocationSettingDataEx1(instance)
	if err != nil {
		return nil, err
	}
	return &VirtualHardDisk{wmivm}, nil
}
