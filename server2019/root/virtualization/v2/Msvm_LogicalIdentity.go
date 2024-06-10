// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.virtualization.v2
//
// ////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_LogicalIdentity struct
type Msvm_LogicalIdentity struct {
	*CIM_LogicalIdentity
}

func NewMsvm_LogicalIdentityEx1(instance *cim.WmiInstance) (newInstance *Msvm_LogicalIdentity, err error) {
	tmp, err := NewCIM_LogicalIdentityEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_LogicalIdentity{
		CIM_LogicalIdentity: tmp,
	}
	return
}

func NewMsvm_LogicalIdentityEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_LogicalIdentity, err error) {
	tmp, err := NewCIM_LogicalIdentityEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_LogicalIdentity{
		CIM_LogicalIdentity: tmp,
	}
	return
}
