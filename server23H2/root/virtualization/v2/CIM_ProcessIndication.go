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

// CIM_ProcessIndication struct
type CIM_ProcessIndication struct {
	*CIM_Indication
}

func NewCIM_ProcessIndicationEx1(instance *cim.WmiInstance) (newInstance *CIM_ProcessIndication, err error) {
	tmp, err := NewCIM_IndicationEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_ProcessIndication{
		CIM_Indication: tmp,
	}
	return
}

func NewCIM_ProcessIndicationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_ProcessIndication, err error) {
	tmp, err := NewCIM_IndicationEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_ProcessIndication{
		CIM_Indication: tmp,
	}
	return
}
