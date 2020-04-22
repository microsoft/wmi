// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package virtualnetworkadapter

import (
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
)

type VirtualNetworkAdapterCollection []*VirtualNetworkAdapter

func NewVirtualNetworkAdapterCollection(instances []*wmi.WmiInstance) (col VirtualNetworkAdapterCollection, err error) {
	for _, inst := range instances {
		na, err1 := NewVirtualNetworkAdapter(inst)
		if err1 != nil {
			err = err1
			return
		}
		col = append(col, na)
	}
	return
}

func (vms *VirtualNetworkAdapterCollection) Close() (err error) {
	for _, value := range *vms {
		value.Close()
	}
	return nil
}

func (vms *VirtualNetworkAdapterCollection) String() string {
	return ""
}
