// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package clusternode

import (
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
)

type ClusterNodeCollection []*ClusterNode

func NewClusterNodeCollection(instances []*wmi.WmiInstance) (col ClusterNodeCollection, err error) {
	for _, inst := range instances {
		na, err1 := NewClusterNode(inst)
		if err1 != nil {
			err = err1
			return
		}
		col = append(col, na)
	}
	return
}

func (instances *ClusterNodeCollection) Close() (err error) {
	for _, value := range *instances {
		value.Close()
	}
	return nil
}

func (instances *ClusterNodeCollection) String() string {
	return ""
}
