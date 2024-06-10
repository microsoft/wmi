// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MLNX_SystemDevice struct
type MLNX_SystemDevice struct {
	*CIM_SystemDevice
}

func NewMLNX_SystemDeviceEx1(instance *cim.WmiInstance) (newInstance *MLNX_SystemDevice, err error) {
	tmp, err := NewCIM_SystemDeviceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MLNX_SystemDevice{
		CIM_SystemDevice: tmp,
	}
	return
}

func NewMLNX_SystemDeviceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MLNX_SystemDevice, err error) {
	tmp, err := NewCIM_SystemDeviceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MLNX_SystemDevice{
		CIM_SystemDevice: tmp,
	}
	return
}
