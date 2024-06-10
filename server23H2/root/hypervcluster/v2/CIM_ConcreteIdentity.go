// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.HyperVCluster.v2
//
// ////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_ConcreteIdentity struct
type CIM_ConcreteIdentity struct {
	*CIM_LogicalIdentity
}

func NewCIM_ConcreteIdentityEx1(instance *cim.WmiInstance) (newInstance *CIM_ConcreteIdentity, err error) {
	tmp, err := NewCIM_LogicalIdentityEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_ConcreteIdentity{
		CIM_LogicalIdentity: tmp,
	}
	return
}

func NewCIM_ConcreteIdentityEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_ConcreteIdentity, err error) {
	tmp, err := NewCIM_LogicalIdentityEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_ConcreteIdentity{
		CIM_LogicalIdentity: tmp,
	}
	return
}
