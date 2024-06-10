// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_RecordAppliesToElement struct
type CIM_RecordAppliesToElement struct {
	*CIM_Dependency
}

func NewCIM_RecordAppliesToElementEx1(instance *cim.WmiInstance) (newInstance *CIM_RecordAppliesToElement, err error) {
	tmp, err := NewCIM_DependencyEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_RecordAppliesToElement{
		CIM_Dependency: tmp,
	}
	return
}

func NewCIM_RecordAppliesToElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_RecordAppliesToElement, err error) {
	tmp, err := NewCIM_DependencyEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_RecordAppliesToElement{
		CIM_Dependency: tmp,
	}
	return
}
