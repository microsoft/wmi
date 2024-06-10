// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_ConcreteComponent struct
type Msvm_ConcreteComponent struct {
	*CIM_ConcreteComponent
}

func NewMsvm_ConcreteComponentEx1(instance *cim.WmiInstance) (newInstance *Msvm_ConcreteComponent, err error) {
	tmp, err := NewCIM_ConcreteComponentEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_ConcreteComponent{
		CIM_ConcreteComponent: tmp,
	}
	return
}

func NewMsvm_ConcreteComponentEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_ConcreteComponent, err error) {
	tmp, err := NewCIM_ConcreteComponentEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_ConcreteComponent{
		CIM_ConcreteComponent: tmp,
	}
	return
}
