// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package groupset

import (
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
)

type GroupSetCollection []*GroupSet

func NewGroupSetCollection(instances []*wmi.WmiInstance) (col GroupSetCollection, err error) {
	for _, inst := range instances {
		na, err1 := NewGroupSet(inst)
		if err1 != nil {
			err = err1
			return
		}
		col = append(col, na)
	}
	return
}

func (instances *GroupSetCollection) Close() (err error) {
	for _, value := range *instances {
		value.Close()
	}
	return nil
}

func (instances *GroupSetCollection) String() string {
	return ""
}
