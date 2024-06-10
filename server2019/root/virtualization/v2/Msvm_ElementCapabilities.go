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

// Msvm_ElementCapabilities struct
type Msvm_ElementCapabilities struct {
	*CIM_ElementCapabilities
}

func NewMsvm_ElementCapabilitiesEx1(instance *cim.WmiInstance) (newInstance *Msvm_ElementCapabilities, err error) {
	tmp, err := NewCIM_ElementCapabilitiesEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_ElementCapabilities{
		CIM_ElementCapabilities: tmp,
	}
	return
}

func NewMsvm_ElementCapabilitiesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_ElementCapabilities, err error) {
	tmp, err := NewCIM_ElementCapabilitiesEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_ElementCapabilities{
		CIM_ElementCapabilities: tmp,
	}
	return
}
