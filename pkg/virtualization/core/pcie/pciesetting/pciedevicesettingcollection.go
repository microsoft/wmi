// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package pciesetting

import (
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
)

type PcieDeviceSettingCollection []*PcieDeviceSetting

func NewPcieDeviceSettingCollection(instances []*wmi.WmiInstance) (col PcieDeviceSettingCollection, err error) {
	for _, inst := range instances {
		na, err1 := NewPcieDeviceSetting(inst)
		if err1 != nil {
			err = err1
			return
		}
		col = append(col, na)
	}
	return
}

func (vms *PcieDeviceSettingCollection) Close() (err error) {
	for _, value := range *vms {
		value.Close()
	}
	return nil
}
