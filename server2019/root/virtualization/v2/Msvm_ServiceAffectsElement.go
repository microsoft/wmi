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

// Msvm_ServiceAffectsElement struct
type Msvm_ServiceAffectsElement struct {
	*CIM_ServiceAffectsElement
}

func NewMsvm_ServiceAffectsElementEx1(instance *cim.WmiInstance) (newInstance *Msvm_ServiceAffectsElement, err error) {
	tmp, err := NewCIM_ServiceAffectsElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_ServiceAffectsElement{
		CIM_ServiceAffectsElement: tmp,
	}
	return
}

func NewMsvm_ServiceAffectsElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_ServiceAffectsElement, err error) {
	tmp, err := NewCIM_ServiceAffectsElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_ServiceAffectsElement{
		CIM_ServiceAffectsElement: tmp,
	}
	return
}
