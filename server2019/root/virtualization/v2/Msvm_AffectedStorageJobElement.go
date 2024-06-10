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

// Msvm_AffectedStorageJobElement struct
type Msvm_AffectedStorageJobElement struct {
	*CIM_AffectedJobElement
}

func NewMsvm_AffectedStorageJobElementEx1(instance *cim.WmiInstance) (newInstance *Msvm_AffectedStorageJobElement, err error) {
	tmp, err := NewCIM_AffectedJobElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_AffectedStorageJobElement{
		CIM_AffectedJobElement: tmp,
	}
	return
}

func NewMsvm_AffectedStorageJobElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_AffectedStorageJobElement, err error) {
	tmp, err := NewCIM_AffectedJobElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_AffectedStorageJobElement{
		CIM_AffectedJobElement: tmp,
	}
	return
}
