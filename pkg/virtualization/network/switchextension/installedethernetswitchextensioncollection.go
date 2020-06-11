// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package switchextension

import (
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
)

type InstalledEthernetSwitchExtensionCollection []*InstalledEthernetSwitchExtension

func NewInstalledEthernetSwitchExtensionCollection(instances []*wmi.WmiInstance) (col InstalledEthernetSwitchExtensionCollection, err error) {
	for _, inst := range instances {
		na, err1 := NewInstalledEthernetSwitchExtension(inst)
		if err1 != nil {
			err = err1
			return
		}
		col = append(col, na)
	}
	return
}

func (vms *InstalledEthernetSwitchExtensionCollection) Close() (err error) {
	for _, value := range *vms {
		value.Close()
	}
	return nil
}

func (vms *InstalledEthernetSwitchExtensionCollection) String() string {
	return ""
}
