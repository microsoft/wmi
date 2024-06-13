// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.MSCluster
//////////////////////////////////////////////
package mscluster

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_ComponentCS struct
type CIM_ComponentCS struct {
	*CIM_SystemComponent
}

func NewCIM_ComponentCSEx1(instance *cim.WmiInstance) (newInstance *CIM_ComponentCS, err error) {
	tmp, err := NewCIM_SystemComponentEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_ComponentCS{
		CIM_SystemComponent: tmp,
	}
	return
}

func NewCIM_ComponentCSEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_ComponentCS, err error) {
	tmp, err := NewCIM_SystemComponentEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_ComponentCS{
		CIM_SystemComponent: tmp,
	}
	return
}
