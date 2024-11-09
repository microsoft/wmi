// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package resource

import (
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
)

type ResourceCollection []*Resource

func NewResourceCollection(instances []*wmi.WmiInstance) (col ResourceCollection, err error) {
	for _, inst := range instances {
		na, err1 := NewResource(inst)
		if err1 != nil {
			err = err1
			return
		}
		col = append(col, na)
	}
	return
}

func (instances *ResourceCollection) Close() (err error) {
	for _, value := range *instances {
		value.Close()
	}
	return nil
}

func (instances *ResourceCollection) String() string {
	return ""
}
