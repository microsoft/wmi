// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package resourcegroup

import (
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
)

type ResourceGroupCollection []*ResourceGroup

func NewResourceGroupCollection(instances []*wmi.WmiInstance) (col ResourceGroupCollection, err error) {
	for _, inst := range instances {
		na, err1 := NewResourceGroup(inst)
		if err1 != nil {
			err = err1
			return
		}
		col = append(col, na)
	}
	return
}

func (instances *ResourceGroupCollection) Close() (err error) {
	for _, value := range *instances {
		value.Close()
	}
	return nil
}

func (instances *ResourceGroupCollection) String() string {
	return ""
}
