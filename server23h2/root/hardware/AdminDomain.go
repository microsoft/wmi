// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.Hardware
//////////////////////////////////////////////
package hardware

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// AdminDomain struct
type AdminDomain struct {
	*CIM_AdminDomain
}

func NewAdminDomainEx1(instance *cim.WmiInstance) (newInstance *AdminDomain, err error) {
	tmp, err := NewCIM_AdminDomainEx1(instance)

	if err != nil {
		return
	}
	newInstance = &AdminDomain{
		CIM_AdminDomain: tmp,
	}
	return
}

func NewAdminDomainEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *AdminDomain, err error) {
	tmp, err := NewCIM_AdminDomainEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &AdminDomain{
		CIM_AdminDomain: tmp,
	}
	return
}
