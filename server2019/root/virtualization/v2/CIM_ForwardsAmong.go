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

// CIM_ForwardsAmong struct
type CIM_ForwardsAmong struct {
	*CIM_ServiceSAPDependency
}

func NewCIM_ForwardsAmongEx1(instance *cim.WmiInstance) (newInstance *CIM_ForwardsAmong, err error) {
	tmp, err := NewCIM_ServiceSAPDependencyEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_ForwardsAmong{
		CIM_ServiceSAPDependency: tmp,
	}
	return
}

func NewCIM_ForwardsAmongEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_ForwardsAmong, err error) {
	tmp, err := NewCIM_ServiceSAPDependencyEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_ForwardsAmong{
		CIM_ServiceSAPDependency: tmp,
	}
	return
}
