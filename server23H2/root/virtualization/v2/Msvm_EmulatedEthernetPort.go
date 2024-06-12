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

// Msvm_EmulatedEthernetPort struct
type Msvm_EmulatedEthernetPort struct {
	*CIM_EthernetPort
}

func NewMsvm_EmulatedEthernetPortEx1(instance *cim.WmiInstance) (newInstance *Msvm_EmulatedEthernetPort, err error) {
	tmp, err := NewCIM_EthernetPortEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_EmulatedEthernetPort{
		CIM_EthernetPort: tmp,
	}
	return
}

func NewMsvm_EmulatedEthernetPortEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_EmulatedEthernetPort, err error) {
	tmp, err := NewCIM_EthernetPortEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_EmulatedEthernetPort{
		CIM_EthernetPort: tmp,
	}
	return
}
