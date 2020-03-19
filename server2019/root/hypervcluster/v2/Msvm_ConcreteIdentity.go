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

// Msvm_ConcreteIdentity struct
type Msvm_ConcreteIdentity struct {
	*CIM_ConcreteIdentity
}

func NewMsvm_ConcreteIdentityEx1(instance *cim.WmiInstance) (newInstance *Msvm_ConcreteIdentity, err error) {
	tmp, err := NewCIM_ConcreteIdentityEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_ConcreteIdentity{
		CIM_ConcreteIdentity: tmp,
	}
	return
}

func NewMsvm_ConcreteIdentityEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_ConcreteIdentity, err error) {
	tmp, err := NewCIM_ConcreteIdentityEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_ConcreteIdentity{
		CIM_ConcreteIdentity: tmp,
	}
	return
}
