// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package drive

import (
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	"github.com/microsoft/wmi/server2019/root/virtualization/v2"
)

type DvdDrive struct {
	*v2.Msvm_StorageAllocationSettingData
}

func NewDvdDrive(instance *wmi.WmiInstance) (*DvdDrive, error) {
	wmivm, err := v2.NewMsvm_StorageAllocationSettingDataEx1(instance)
	if err != nil {
		return nil, err
	}
	return &DvdDrive{wmivm}, nil
}