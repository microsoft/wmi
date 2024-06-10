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

// Msvm_SystemComponent struct
type Msvm_SystemComponent struct {
	*CIM_SystemComponent
}

func NewMsvm_SystemComponentEx1(instance *cim.WmiInstance) (newInstance *Msvm_SystemComponent, err error) {
	tmp, err := NewCIM_SystemComponentEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_SystemComponent{
		CIM_SystemComponent: tmp,
	}
	return
}

func NewMsvm_SystemComponentEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_SystemComponent, err error) {
	tmp, err := NewCIM_SystemComponentEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_SystemComponent{
		CIM_SystemComponent: tmp,
	}
	return
}
