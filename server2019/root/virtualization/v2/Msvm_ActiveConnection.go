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

// Msvm_ActiveConnection struct
type Msvm_ActiveConnection struct {
	*CIM_ActiveConnection
}

func NewMsvm_ActiveConnectionEx1(instance *cim.WmiInstance) (newInstance *Msvm_ActiveConnection, err error) {
	tmp, err := NewCIM_ActiveConnectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_ActiveConnection{
		CIM_ActiveConnection: tmp,
	}
	return
}

func NewMsvm_ActiveConnectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_ActiveConnection, err error) {
	tmp, err := NewCIM_ActiveConnectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_ActiveConnection{
		CIM_ActiveConnection: tmp,
	}
	return
}
