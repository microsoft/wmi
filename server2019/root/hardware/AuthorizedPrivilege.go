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

// AuthorizedPrivilege struct
type AuthorizedPrivilege struct {
	*CIM_AuthorizedPrivilege
}

func NewAuthorizedPrivilegeEx1(instance *cim.WmiInstance) (newInstance *AuthorizedPrivilege, err error) {
	tmp, err := NewCIM_AuthorizedPrivilegeEx1(instance)

	if err != nil {
		return
	}
	newInstance = &AuthorizedPrivilege{
		CIM_AuthorizedPrivilege: tmp,
	}
	return
}

func NewAuthorizedPrivilegeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *AuthorizedPrivilege, err error) {
	tmp, err := NewCIM_AuthorizedPrivilegeEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &AuthorizedPrivilege{
		CIM_AuthorizedPrivilege: tmp,
	}
	return
}
