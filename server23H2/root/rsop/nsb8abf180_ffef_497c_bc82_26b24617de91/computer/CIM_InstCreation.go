// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.RSOP.NSB8ABF180_FFEF_497C_BC82_26B24617DE91.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_InstCreation struct
type CIM_InstCreation struct {
	*CIM_InstIndication
}

func NewCIM_InstCreationEx1(instance *cim.WmiInstance) (newInstance *CIM_InstCreation, err error) {
	tmp, err := NewCIM_InstIndicationEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_InstCreation{
		CIM_InstIndication: tmp,
	}
	return
}

func NewCIM_InstCreationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_InstCreation, err error) {
	tmp, err := NewCIM_InstIndicationEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_InstCreation{
		CIM_InstIndication: tmp,
	}
	return
}
