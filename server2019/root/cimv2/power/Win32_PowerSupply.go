// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.CIMV2.power
//////////////////////////////////////////////
package power

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_PowerSupply struct
type Win32_PowerSupply struct {
	*CIM_PowerSupply
}

func NewWin32_PowerSupplyEx1(instance *cim.WmiInstance) (newInstance *Win32_PowerSupply, err error) {
	tmp, err := NewCIM_PowerSupplyEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PowerSupply{
		CIM_PowerSupply: tmp,
	}
	return
}

func NewWin32_PowerSupplyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PowerSupply, err error) {
	tmp, err := NewCIM_PowerSupplyEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PowerSupply{
		CIM_PowerSupply: tmp,
	}
	return
}
