// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package switchextension

import (
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
)

type EthernetSwitchExtensionCollection []*EthernetSwitchExtension

func NewEthernetSwitchExtensionCollection(instances []*wmi.WmiInstance) (col EthernetSwitchExtensionCollection, err error) {
	for _, inst := range instances {
		na, err1 := NewEthernetSwitchExtension(inst)
		if err1 != nil {
			err = err1
			return
		}
		col = append(col, na)
	}
	return
}

func (vms *EthernetSwitchExtensionCollection) Close() (err error) {
	for _, value := range *vms {
		value.Close()
	}
	return nil
}

func (vms *EthernetSwitchExtensionCollection) String() string {
	return ""
}
