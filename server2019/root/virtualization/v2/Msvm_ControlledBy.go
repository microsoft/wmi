// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.virtualization.v2
//
// ////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_ControlledBy struct
type Msvm_ControlledBy struct {
	*CIM_ControlledBy
}

func NewMsvm_ControlledByEx1(instance *cim.WmiInstance) (newInstance *Msvm_ControlledBy, err error) {
	tmp, err := NewCIM_ControlledByEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_ControlledBy{
		CIM_ControlledBy: tmp,
	}
	return
}

func NewMsvm_ControlledByEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_ControlledBy, err error) {
	tmp, err := NewCIM_ControlledByEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_ControlledBy{
		CIM_ControlledBy: tmp,
	}
	return
}
