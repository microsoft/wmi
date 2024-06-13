// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.CIMV2.power
//////////////////////////////////////////////
package power

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_ConcreteComponent struct
type CIM_ConcreteComponent struct {
	*CIM_Component
}

func NewCIM_ConcreteComponentEx1(instance *cim.WmiInstance) (newInstance *CIM_ConcreteComponent, err error) {
	tmp, err := NewCIM_ComponentEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_ConcreteComponent{
		CIM_Component: tmp,
	}
	return
}

func NewCIM_ConcreteComponentEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_ConcreteComponent, err error) {
	tmp, err := NewCIM_ComponentEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_ConcreteComponent{
		CIM_Component: tmp,
	}
	return
}
