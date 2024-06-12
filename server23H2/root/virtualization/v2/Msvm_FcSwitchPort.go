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

// Msvm_FcSwitchPort struct
type Msvm_FcSwitchPort struct {
	*CIM_FCPort
}

func NewMsvm_FcSwitchPortEx1(instance *cim.WmiInstance) (newInstance *Msvm_FcSwitchPort, err error) {
	tmp, err := NewCIM_FCPortEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_FcSwitchPort{
		CIM_FCPort: tmp,
	}
	return
}

func NewMsvm_FcSwitchPortEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_FcSwitchPort, err error) {
	tmp, err := NewCIM_FCPortEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_FcSwitchPort{
		CIM_FCPort: tmp,
	}
	return
}
