// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package pcie

import (
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
)

type PcieDeviceCollection []*PcieDevice

func NewPcieDeviceCollection(instances []*wmi.WmiInstance) (col PcieDeviceCollection, err error) {
	for _, inst := range instances {
		na, err1 := NewPcieDeviceEx1(inst)
		if err1 != nil {
			err = err1
			return
		}
		col = append(col, na)
	}
	return
}

func (vms *PcieDeviceCollection) Close() (err error) {
	for _, value := range *vms {
		value.Close()
	}
	return nil
}
