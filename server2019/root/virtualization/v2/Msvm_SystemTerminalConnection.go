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

// Msvm_SystemTerminalConnection struct
type Msvm_SystemTerminalConnection struct {
	*CIM_HostedDependency
}

func NewMsvm_SystemTerminalConnectionEx1(instance *cim.WmiInstance) (newInstance *Msvm_SystemTerminalConnection, err error) {
	tmp, err := NewCIM_HostedDependencyEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_SystemTerminalConnection{
		CIM_HostedDependency: tmp,
	}
	return
}

func NewMsvm_SystemTerminalConnectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_SystemTerminalConnection, err error) {
	tmp, err := NewCIM_HostedDependencyEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_SystemTerminalConnection{
		CIM_HostedDependency: tmp,
	}
	return
}
