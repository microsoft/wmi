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

// Msvm_BindsToLANEndpoint struct
type Msvm_BindsToLANEndpoint struct {
	*CIM_BindsToLANEndpoint
}

func NewMsvm_BindsToLANEndpointEx1(instance *cim.WmiInstance) (newInstance *Msvm_BindsToLANEndpoint, err error) {
	tmp, err := NewCIM_BindsToLANEndpointEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_BindsToLANEndpoint{
		CIM_BindsToLANEndpoint: tmp,
	}
	return
}

func NewMsvm_BindsToLANEndpointEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_BindsToLANEndpoint, err error) {
	tmp, err := NewCIM_BindsToLANEndpointEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_BindsToLANEndpoint{
		CIM_BindsToLANEndpoint: tmp,
	}
	return
}
