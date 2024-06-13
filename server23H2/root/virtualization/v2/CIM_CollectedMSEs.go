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

// CIM_CollectedMSEs struct
type CIM_CollectedMSEs struct {
	*CIM_MemberOfCollection
}

func NewCIM_CollectedMSEsEx1(instance *cim.WmiInstance) (newInstance *CIM_CollectedMSEs, err error) {
	tmp, err := NewCIM_MemberOfCollectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_CollectedMSEs{
		CIM_MemberOfCollection: tmp,
	}
	return
}

func NewCIM_CollectedMSEsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_CollectedMSEs, err error) {
	tmp, err := NewCIM_MemberOfCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_CollectedMSEs{
		CIM_MemberOfCollection: tmp,
	}
	return
}
