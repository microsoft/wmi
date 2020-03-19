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

// Msvm_ReplicationProvider struct
type Msvm_ReplicationProvider struct {
	*CIM_ManagedSystemElement
}

func NewMsvm_ReplicationProviderEx1(instance *cim.WmiInstance) (newInstance *Msvm_ReplicationProvider, err error) {
	tmp, err := NewCIM_ManagedSystemElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_ReplicationProvider{
		CIM_ManagedSystemElement: tmp,
	}
	return
}

func NewMsvm_ReplicationProviderEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_ReplicationProvider, err error) {
	tmp, err := NewCIM_ManagedSystemElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_ReplicationProvider{
		CIM_ManagedSystemElement: tmp,
	}
	return
}
