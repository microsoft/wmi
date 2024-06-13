// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Hardware
//////////////////////////////////////////////
package hardware

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_AuthorizedPrivilege struct
type CIM_AuthorizedPrivilege struct {
	*CIM_Privilege
}

func NewCIM_AuthorizedPrivilegeEx1(instance *cim.WmiInstance) (newInstance *CIM_AuthorizedPrivilege, err error) {
	tmp, err := NewCIM_PrivilegeEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_AuthorizedPrivilege{
		CIM_Privilege: tmp,
	}
	return
}

func NewCIM_AuthorizedPrivilegeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_AuthorizedPrivilege, err error) {
	tmp, err := NewCIM_PrivilegeEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_AuthorizedPrivilege{
		CIM_Privilege: tmp,
	}
	return
}
