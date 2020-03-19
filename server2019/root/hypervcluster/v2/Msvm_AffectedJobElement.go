// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_AffectedJobElement struct
type Msvm_AffectedJobElement struct {
	*CIM_AffectedJobElement
}

func NewMsvm_AffectedJobElementEx1(instance *cim.WmiInstance) (newInstance *Msvm_AffectedJobElement, err error) {
	tmp, err := NewCIM_AffectedJobElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_AffectedJobElement{
		CIM_AffectedJobElement: tmp,
	}
	return
}

func NewMsvm_AffectedJobElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_AffectedJobElement, err error) {
	tmp, err := NewCIM_AffectedJobElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_AffectedJobElement{
		CIM_AffectedJobElement: tmp,
	}
	return
}
