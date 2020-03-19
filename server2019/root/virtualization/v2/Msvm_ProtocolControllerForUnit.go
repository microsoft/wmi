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

// Msvm_ProtocolControllerForUnit struct
type Msvm_ProtocolControllerForUnit struct {
	*CIM_ProtocolControllerForUnit
}

func NewMsvm_ProtocolControllerForUnitEx1(instance *cim.WmiInstance) (newInstance *Msvm_ProtocolControllerForUnit, err error) {
	tmp, err := NewCIM_ProtocolControllerForUnitEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_ProtocolControllerForUnit{
		CIM_ProtocolControllerForUnit: tmp,
	}
	return
}

func NewMsvm_ProtocolControllerForUnitEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_ProtocolControllerForUnit, err error) {
	tmp, err := NewCIM_ProtocolControllerForUnitEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_ProtocolControllerForUnit{
		CIM_ProtocolControllerForUnit: tmp,
	}
	return
}
