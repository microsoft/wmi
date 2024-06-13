// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_StorageAlert struct
type Msvm_StorageAlert struct {
	*CIM_AlertIndication
}

func NewMsvm_StorageAlertEx1(instance *cim.WmiInstance) (newInstance *Msvm_StorageAlert, err error) {
	tmp, err := NewCIM_AlertIndicationEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_StorageAlert{
		CIM_AlertIndication: tmp,
	}
	return
}

func NewMsvm_StorageAlertEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_StorageAlert, err error) {
	tmp, err := NewCIM_AlertIndicationEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_StorageAlert{
		CIM_AlertIndication: tmp,
	}
	return
}
