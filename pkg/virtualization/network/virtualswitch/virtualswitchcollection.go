// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package virtualswitch

import (
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
)

type VirtualSwitchCollection []*VirtualSwitch

func NewVirtualSwitchCollection(instances []*wmi.WmiInstance) (col VirtualSwitchCollection, err error) {
	for _, inst := range instances {
		na, err1 := NewVirtualSwitch(inst)
		if err1 != nil {
			err = err1
			return
		}
		col = append(col, na)
	}
	return
}

func (instances *VirtualSwitchCollection) Close() (err error) {
	for _, value := range *instances {
		value.Close()
	}
	return nil
}

func (instances *VirtualSwitchCollection) String() string {
	return ""
}
