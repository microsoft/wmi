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

// Msvm_SystemBIOS struct
type Msvm_SystemBIOS struct {
	*CIM_SystemBIOS
}

func NewMsvm_SystemBIOSEx1(instance *cim.WmiInstance) (newInstance *Msvm_SystemBIOS, err error) {
	tmp, err := NewCIM_SystemBIOSEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_SystemBIOS{
		CIM_SystemBIOS: tmp,
	}
	return
}

func NewMsvm_SystemBIOSEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_SystemBIOS, err error) {
	tmp, err := NewCIM_SystemBIOSEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_SystemBIOS{
		CIM_SystemBIOS: tmp,
	}
	return
}
