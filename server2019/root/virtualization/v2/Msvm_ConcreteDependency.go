// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_ConcreteDependency struct
type Msvm_ConcreteDependency struct {
	*CIM_ConcreteDependency
}

func NewMsvm_ConcreteDependencyEx1(instance *cim.WmiInstance) (newInstance *Msvm_ConcreteDependency, err error) {
	tmp, err := NewCIM_ConcreteDependencyEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_ConcreteDependency{
		CIM_ConcreteDependency: tmp,
	}
	return
}

func NewMsvm_ConcreteDependencyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_ConcreteDependency, err error) {
	tmp, err := NewCIM_ConcreteDependencyEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_ConcreteDependency{
		CIM_ConcreteDependency: tmp,
	}
	return
}
