// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.RSOP.NSFEBF1FF6_ACFD_4A47_B994_DBBA83EEA64F
//////////////////////////////////////////////
package nsfebf1ff6_acfd_4a47_b994_dbba83eea64f

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_ClassDeletion struct
type CIM_ClassDeletion struct {
	*CIM_ClassIndication
}

func NewCIM_ClassDeletionEx1(instance *cim.WmiInstance) (newInstance *CIM_ClassDeletion, err error) {
	tmp, err := NewCIM_ClassIndicationEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_ClassDeletion{
		CIM_ClassIndication: tmp,
	}
	return
}

func NewCIM_ClassDeletionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_ClassDeletion, err error) {
	tmp, err := NewCIM_ClassIndicationEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_ClassDeletion{
		CIM_ClassIndication: tmp,
	}
	return
}