// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package drive

import (
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
)

type DvdDrive struct {
	*VirtualDrive
}

func NewDvdDrive(instance *wmi.WmiInstance) (*DvdDrive, error) {
	wmivm, err := NewVirtualDrive(instance)
	if err != nil {
		return nil, err
	}
	return &DvdDrive{wmivm}, nil
}

func GetRelatedStorageAllocationSettingData(instance *wmi.WmiInstance) (value *wmi.WmiInstance, err error) {
	return instance.GetRelated("Msvm_StorageAllocationSettingData")
}
