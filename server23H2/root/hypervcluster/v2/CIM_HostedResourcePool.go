// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_HostedResourcePool struct
type CIM_HostedResourcePool struct {
	*CIM_SystemComponent
}

func NewCIM_HostedResourcePoolEx1(instance *cim.WmiInstance) (newInstance *CIM_HostedResourcePool, err error) {
	tmp, err := NewCIM_SystemComponentEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_HostedResourcePool{
		CIM_SystemComponent: tmp,
	}
	return
}

func NewCIM_HostedResourcePoolEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_HostedResourcePool, err error) {
	tmp, err := NewCIM_SystemComponentEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_HostedResourcePool{
		CIM_SystemComponent: tmp,
	}
	return
}
