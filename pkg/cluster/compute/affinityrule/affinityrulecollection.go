// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package affinityrule

import (
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
)

type AffinityRuleCollection []*AffinityRule

func NewAffinityRuleCollection(instances []*wmi.WmiInstance) (col AffinityRuleCollection, err error) {
	for _, inst := range instances {
		na, err1 := NewAffinityRule(inst)
		if err1 != nil {
			err = err1
			return
		}
		col = append(col, na)
	}
	return
}

func (instances *AffinityRuleCollection) Close() (err error) {
	for _, value := range *instances {
		value.Close()
	}
	return nil
}

func (instances *AffinityRuleCollection) String() string {
	return ""
}
