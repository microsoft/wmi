// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package pnp

import (
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
)

type PnpEntityCollection []*PnpEntity

func NewPnpEntityCollection(instances []*wmi.WmiInstance) (col PnpEntityCollection, err error) {
	for _, inst := range instances {
		na, err1 := NewPnpEntity(inst)
		if err1 != nil {
			err = err1
			return
		}
		col = append(col, na)
	}
	return
}

func (vms *PnpEntityCollection) Close() (err error) {
	for _, value := range *vms {
		value.Close()
	}
	return nil
}

func (vms *PnpEntityCollection) String() string {
	return ""
}
