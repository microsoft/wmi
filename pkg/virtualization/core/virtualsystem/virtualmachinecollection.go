// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package virtualsystem

import (
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
)

type VirtualMachineCollection []*VirtualMachine

func NewVirtualMachineCollection(instances []*wmi.WmiInstance) (col VirtualMachineCollection, err error) {
	for _, inst := range instances {
		na, err1 := NewVirtualMachine(inst)
		if err1 != nil {
			err = err1
			return
		}
		col = append(col, na)
	}
	return
}

func (vms *VirtualMachineCollection) Close() (err error) {
	for _, value := range *vms {
		value.Close()
	}
	return nil
}

func (vms *VirtualMachineCollection) String() string {
	return ""
}
