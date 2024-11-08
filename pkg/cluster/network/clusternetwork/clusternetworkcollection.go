// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package clusternetwork

import (
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
)

type ClusterNetworkCollection []*ClusterNetwork

func NewClusterNetworkCollection(instances []*wmi.WmiInstance) (col ClusterNetworkCollection, err error) {
	for _, inst := range instances {
		na, err1 := NewClusterNetwork(inst)
		if err1 != nil {
			err = err1
			return
		}
		col = append(col, na)
	}
	return
}

func (instances *ClusterNetworkCollection) Close() (err error) {
	for _, value := range *instances {
		value.Close()
	}
	return nil
}

func (instances *ClusterNetworkCollection) String() string {
	return ""
}
