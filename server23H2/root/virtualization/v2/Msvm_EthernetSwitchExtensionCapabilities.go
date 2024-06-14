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

// Msvm_EthernetSwitchExtensionCapabilities struct
type Msvm_EthernetSwitchExtensionCapabilities struct {
	*CIM_ElementCapabilities
}

func NewMsvm_EthernetSwitchExtensionCapabilitiesEx1(instance *cim.WmiInstance) (newInstance *Msvm_EthernetSwitchExtensionCapabilities, err error) {
	tmp, err := NewCIM_ElementCapabilitiesEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_EthernetSwitchExtensionCapabilities{
		CIM_ElementCapabilities: tmp,
	}
	return
}

func NewMsvm_EthernetSwitchExtensionCapabilitiesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_EthernetSwitchExtensionCapabilities, err error) {
	tmp, err := NewCIM_ElementCapabilitiesEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_EthernetSwitchExtensionCapabilities{
		CIM_ElementCapabilities: tmp,
	}
	return
}
