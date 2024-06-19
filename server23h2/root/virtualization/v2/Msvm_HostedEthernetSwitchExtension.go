// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_HostedEthernetSwitchExtension struct
type Msvm_HostedEthernetSwitchExtension struct {
	*CIM_HostedDependency
}

func NewMsvm_HostedEthernetSwitchExtensionEx1(instance *cim.WmiInstance) (newInstance *Msvm_HostedEthernetSwitchExtension, err error) {
	tmp, err := NewCIM_HostedDependencyEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_HostedEthernetSwitchExtension{
		CIM_HostedDependency: tmp,
	}
	return
}

func NewMsvm_HostedEthernetSwitchExtensionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_HostedEthernetSwitchExtension, err error) {
	tmp, err := NewCIM_HostedDependencyEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_HostedEthernetSwitchExtension{
		CIM_HostedDependency: tmp,
	}
	return
}
