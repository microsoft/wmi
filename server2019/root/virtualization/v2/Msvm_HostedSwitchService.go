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

// Msvm_HostedSwitchService struct
type Msvm_HostedSwitchService struct {
	*CIM_HostedService
}

func NewMsvm_HostedSwitchServiceEx1(instance *cim.WmiInstance) (newInstance *Msvm_HostedSwitchService, err error) {
	tmp, err := NewCIM_HostedServiceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_HostedSwitchService{
		CIM_HostedService: tmp,
	}
	return
}

func NewMsvm_HostedSwitchServiceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_HostedSwitchService, err error) {
	tmp, err := NewCIM_HostedServiceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_HostedSwitchService{
		CIM_HostedService: tmp,
	}
	return
}
