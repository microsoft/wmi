// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package disk

import (
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
)

type VirtualHardDiskCollection []*VirtualHardDisk

func NewVirtualHardDiskCollection(instances []*wmi.WmiInstance) (col VirtualHardDiskCollection, err error) {
	for _, inst := range instances {
		na, err1 := NewVirtualHardDisk(inst)
		if err1 != nil {
			err = err1
			return
		}
		col = append(col, na)
	}
	return
}

func (vms *VirtualHardDiskCollection) Close() (err error) {
	for _, value := range *vms {
		value.Close()
	}
	return nil
}

func (vms *VirtualHardDiskCollection) String() string {
	return ""
}
