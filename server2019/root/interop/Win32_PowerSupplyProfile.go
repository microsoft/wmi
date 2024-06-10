// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.Interop
//////////////////////////////////////////////
package interop

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_PowerSupplyProfile struct
type Win32_PowerSupplyProfile struct {
	*CIM_RegisteredProfile
}

func NewWin32_PowerSupplyProfileEx1(instance *cim.WmiInstance) (newInstance *Win32_PowerSupplyProfile, err error) {
	tmp, err := NewCIM_RegisteredProfileEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PowerSupplyProfile{
		CIM_RegisteredProfile: tmp,
	}
	return
}

func NewWin32_PowerSupplyProfileEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PowerSupplyProfile, err error) {
	tmp, err := NewCIM_RegisteredProfileEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PowerSupplyProfile{
		CIM_RegisteredProfile: tmp,
	}
	return
}
