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

// Msvm_SwitchPortDynamicForwarding struct
type Msvm_SwitchPortDynamicForwarding struct {
	*CIM_SwitchPortDynamicForwarding
}

func NewMsvm_SwitchPortDynamicForwardingEx1(instance *cim.WmiInstance) (newInstance *Msvm_SwitchPortDynamicForwarding, err error) {
	tmp, err := NewCIM_SwitchPortDynamicForwardingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_SwitchPortDynamicForwarding{
		CIM_SwitchPortDynamicForwarding: tmp,
	}
	return
}

func NewMsvm_SwitchPortDynamicForwardingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_SwitchPortDynamicForwarding, err error) {
	tmp, err := NewCIM_SwitchPortDynamicForwardingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_SwitchPortDynamicForwarding{
		CIM_SwitchPortDynamicForwarding: tmp,
	}
	return
}
