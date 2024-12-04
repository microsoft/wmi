// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package service

import (
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
)

type ClusterServiceCollection []*ClusterService

func NewClusterServiceCollection(instances []*wmi.WmiInstance) (col ClusterServiceCollection, err error) {
	for _, inst := range instances {
		na, err1 := NewClusterService(inst)
		if err1 != nil {
			err = err1
			return
		}
		col = append(col, na)
	}
	return
}

func (instances *ClusterServiceCollection) Close() (err error) {
	for _, value := range *instances {
		value.Close()
	}
	return nil
}

func (instances *ClusterServiceCollection) String() string {
	return ""
}
