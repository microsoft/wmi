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

// Msvm_VirtualSystemCapabilities struct
type Msvm_VirtualSystemCapabilities struct {
	*CIM_EnabledLogicalElementCapabilities
}

func NewMsvm_VirtualSystemCapabilitiesEx1(instance *cim.WmiInstance) (newInstance *Msvm_VirtualSystemCapabilities, err error) {
	tmp, err := NewCIM_EnabledLogicalElementCapabilitiesEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualSystemCapabilities{
		CIM_EnabledLogicalElementCapabilities: tmp,
	}
	return
}

func NewMsvm_VirtualSystemCapabilitiesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_VirtualSystemCapabilities, err error) {
	tmp, err := NewCIM_EnabledLogicalElementCapabilitiesEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualSystemCapabilities{
		CIM_EnabledLogicalElementCapabilities: tmp,
	}
	return
}
