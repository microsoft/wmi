// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.RSOP.NS9606B916_B14A_4EB6_84F0_165F56E1DD34.Computer
//
// ////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_InstDeletion struct
type CIM_InstDeletion struct {
	*CIM_InstIndication
}

func NewCIM_InstDeletionEx1(instance *cim.WmiInstance) (newInstance *CIM_InstDeletion, err error) {
	tmp, err := NewCIM_InstIndicationEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_InstDeletion{
		CIM_InstIndication: tmp,
	}
	return
}

func NewCIM_InstDeletionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_InstDeletion, err error) {
	tmp, err := NewCIM_InstIndicationEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_InstDeletion{
		CIM_InstIndication: tmp,
	}
	return
}
