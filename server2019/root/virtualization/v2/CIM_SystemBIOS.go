// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_SystemBIOS struct
type CIM_SystemBIOS struct {
	*CIM_SystemComponent
}

func NewCIM_SystemBIOSEx1(instance *cim.WmiInstance) (newInstance *CIM_SystemBIOS, err error) {
	tmp, err := NewCIM_SystemComponentEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_SystemBIOS{
		CIM_SystemComponent: tmp,
	}
	return
}

func NewCIM_SystemBIOSEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_SystemBIOS, err error) {
	tmp, err := NewCIM_SystemComponentEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_SystemBIOS{
		CIM_SystemComponent: tmp,
	}
	return
}
