// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_ResourceAllocationFromPool struct
type Msvm_ResourceAllocationFromPool struct {
	*CIM_ResourceAllocationFromPool
}

func NewMsvm_ResourceAllocationFromPoolEx1(instance *cim.WmiInstance) (newInstance *Msvm_ResourceAllocationFromPool, err error) {
	tmp, err := NewCIM_ResourceAllocationFromPoolEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_ResourceAllocationFromPool{
		CIM_ResourceAllocationFromPool: tmp,
	}
	return
}

func NewMsvm_ResourceAllocationFromPoolEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_ResourceAllocationFromPool, err error) {
	tmp, err := NewCIM_ResourceAllocationFromPoolEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_ResourceAllocationFromPool{
		CIM_ResourceAllocationFromPool: tmp,
	}
	return
}
