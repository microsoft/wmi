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

type DvdDriveCollection []*DvdDrive

func NewDvdDriveCollection(instances []*wmi.WmiInstance) (col DvdDriveCollection, err error) {
	for _, inst := range instances {
		na, err1 := NewDvdDrive(inst)
		if err1 != nil {
			err = err1
			return
		}
		col = append(col, na)
	}
	return
}

func Close(vms *DvdDriveCollection) (err error) {
	for _, value := range *vms {
		value.Close()
	}
	return nil
}

func GetRelatedStorageAllocationSettingData(instance *wmi.WmiInstance) (value *wmi.WmiInstance, err error) {
	return instance.GetRelated("Msvm_StorageAllocationSettingData")
}
