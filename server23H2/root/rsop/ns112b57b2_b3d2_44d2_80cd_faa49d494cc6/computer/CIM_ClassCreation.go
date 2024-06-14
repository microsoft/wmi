// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.RSOP.NS112B57B2_B3D2_44D2_80CD_FAA49D494CC6.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_ClassCreation struct
type CIM_ClassCreation struct {
	*CIM_ClassIndication
}

func NewCIM_ClassCreationEx1(instance *cim.WmiInstance) (newInstance *CIM_ClassCreation, err error) {
	tmp, err := NewCIM_ClassIndicationEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_ClassCreation{
		CIM_ClassIndication: tmp,
	}
	return
}

func NewCIM_ClassCreationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_ClassCreation, err error) {
	tmp, err := NewCIM_ClassIndicationEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_ClassCreation{
		CIM_ClassIndication: tmp,
	}
	return
}
