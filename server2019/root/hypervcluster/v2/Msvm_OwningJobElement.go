// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_OwningJobElement struct
type Msvm_OwningJobElement struct {
	*CIM_OwningJobElement
}

func NewMsvm_OwningJobElementEx1(instance *cim.WmiInstance) (newInstance *Msvm_OwningJobElement, err error) {
	tmp, err := NewCIM_OwningJobElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_OwningJobElement{
		CIM_OwningJobElement: tmp,
	}
	return
}

func NewMsvm_OwningJobElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_OwningJobElement, err error) {
	tmp, err := NewCIM_OwningJobElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_OwningJobElement{
		CIM_OwningJobElement: tmp,
	}
	return
}
