// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.Hardware
//////////////////////////////////////////////
package hardware

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// PrivilegeGroup struct
type PrivilegeGroup struct {
	*CIM_Group
}

func NewPrivilegeGroupEx1(instance *cim.WmiInstance) (newInstance *PrivilegeGroup, err error) {
	tmp, err := NewCIM_GroupEx1(instance)

	if err != nil {
		return
	}
	newInstance = &PrivilegeGroup{
		CIM_Group: tmp,
	}
	return
}

func NewPrivilegeGroupEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PrivilegeGroup, err error) {
	tmp, err := NewCIM_GroupEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PrivilegeGroup{
		CIM_Group: tmp,
	}
	return
}
