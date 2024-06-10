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

// Msvm_BootSourceComponent struct
type Msvm_BootSourceComponent struct {
	*CIM_Component
}

func NewMsvm_BootSourceComponentEx1(instance *cim.WmiInstance) (newInstance *Msvm_BootSourceComponent, err error) {
	tmp, err := NewCIM_ComponentEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_BootSourceComponent{
		CIM_Component: tmp,
	}
	return
}

func NewMsvm_BootSourceComponentEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_BootSourceComponent, err error) {
	tmp, err := NewCIM_ComponentEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_BootSourceComponent{
		CIM_Component: tmp,
	}
	return
}
