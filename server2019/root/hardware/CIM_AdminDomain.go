// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.Hardware
//////////////////////////////////////////////
package hardware

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_AdminDomain struct
type CIM_AdminDomain struct {
	*CIM_System
}

func NewCIM_AdminDomainEx1(instance *cim.WmiInstance) (newInstance *CIM_AdminDomain, err error) {
	tmp, err := NewCIM_SystemEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_AdminDomain{
		CIM_System: tmp,
	}
	return
}

func NewCIM_AdminDomainEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_AdminDomain, err error) {
	tmp, err := NewCIM_SystemEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_AdminDomain{
		CIM_System: tmp,
	}
	return
}
