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

// Msvm_HostedResourcePool struct
type Msvm_HostedResourcePool struct {
	*CIM_SystemComponent
}

func NewMsvm_HostedResourcePoolEx1(instance *cim.WmiInstance) (newInstance *Msvm_HostedResourcePool, err error) {
	tmp, err := NewCIM_SystemComponentEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_HostedResourcePool{
		CIM_SystemComponent: tmp,
	}
	return
}

func NewMsvm_HostedResourcePoolEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_HostedResourcePool, err error) {
	tmp, err := NewCIM_SystemComponentEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_HostedResourcePool{
		CIM_SystemComponent: tmp,
	}
	return
}
