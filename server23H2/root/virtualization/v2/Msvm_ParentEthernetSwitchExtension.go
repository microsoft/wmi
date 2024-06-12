// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_ParentEthernetSwitchExtension struct
type Msvm_ParentEthernetSwitchExtension struct {
	*CIM_Dependency
}

func NewMsvm_ParentEthernetSwitchExtensionEx1(instance *cim.WmiInstance) (newInstance *Msvm_ParentEthernetSwitchExtension, err error) {
	tmp, err := NewCIM_DependencyEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_ParentEthernetSwitchExtension{
		CIM_Dependency: tmp,
	}
	return
}

func NewMsvm_ParentEthernetSwitchExtensionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_ParentEthernetSwitchExtension, err error) {
	tmp, err := NewCIM_DependencyEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_ParentEthernetSwitchExtension{
		CIM_Dependency: tmp,
	}
	return
}
