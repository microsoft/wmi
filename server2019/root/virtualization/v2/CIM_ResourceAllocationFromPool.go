// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_ResourceAllocationFromPool struct
type CIM_ResourceAllocationFromPool struct {
	*CIM_Dependency
}

func NewCIM_ResourceAllocationFromPoolEx1(instance *cim.WmiInstance) (newInstance *CIM_ResourceAllocationFromPool, err error) {
	tmp, err := NewCIM_DependencyEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_ResourceAllocationFromPool{
		CIM_Dependency: tmp,
	}
	return
}

func NewCIM_ResourceAllocationFromPoolEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_ResourceAllocationFromPool, err error) {
	tmp, err := NewCIM_DependencyEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_ResourceAllocationFromPool{
		CIM_Dependency: tmp,
	}
	return
}
