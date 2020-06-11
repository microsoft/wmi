// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package switchport

import (
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
)

type EthernetSwitchPortCollection []*EthernetSwitchPort

func NewEthernetSwitchPortCollection(instances []*wmi.WmiInstance) (col EthernetSwitchPortCollection, err error) {
	for _, inst := range instances {
		na, err1 := NewEthernetSwitchPort(inst)
		if err1 != nil {
			err = err1
			return
		}
		col = append(col, na)
	}
	return
}

func (instances *EthernetSwitchPortCollection) Close() (err error) {
	for _, value := range *instances {
		value.Close()
	}
	return nil
}

func (instances *EthernetSwitchPortCollection) String() string {
	return ""
}
