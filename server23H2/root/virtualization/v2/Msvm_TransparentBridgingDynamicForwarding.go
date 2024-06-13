// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_TransparentBridgingDynamicForwarding struct
type Msvm_TransparentBridgingDynamicForwarding struct {
	*CIM_TransparentBridgingDynamicForwarding
}

func NewMsvm_TransparentBridgingDynamicForwardingEx1(instance *cim.WmiInstance) (newInstance *Msvm_TransparentBridgingDynamicForwarding, err error) {
	tmp, err := NewCIM_TransparentBridgingDynamicForwardingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_TransparentBridgingDynamicForwarding{
		CIM_TransparentBridgingDynamicForwarding: tmp,
	}
	return
}

func NewMsvm_TransparentBridgingDynamicForwardingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_TransparentBridgingDynamicForwarding, err error) {
	tmp, err := NewCIM_TransparentBridgingDynamicForwardingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_TransparentBridgingDynamicForwarding{
		CIM_TransparentBridgingDynamicForwarding: tmp,
	}
	return
}
