// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.virtualization.v2
//
// ////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_LogicalDisk struct
type CIM_LogicalDisk struct {
	*CIM_StorageExtent
}

func NewCIM_LogicalDiskEx1(instance *cim.WmiInstance) (newInstance *CIM_LogicalDisk, err error) {
	tmp, err := NewCIM_StorageExtentEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_LogicalDisk{
		CIM_StorageExtent: tmp,
	}
	return
}

func NewCIM_LogicalDiskEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_LogicalDisk, err error) {
	tmp, err := NewCIM_StorageExtentEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_LogicalDisk{
		CIM_StorageExtent: tmp,
	}
	return
}
