// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MLNX_DriverElementCapabilities struct
type MLNX_DriverElementCapabilities struct {
	*CIM_ElementCapabilities
}

func NewMLNX_DriverElementCapabilitiesEx1(instance *cim.WmiInstance) (newInstance *MLNX_DriverElementCapabilities, err error) {
	tmp, err := NewCIM_ElementCapabilitiesEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MLNX_DriverElementCapabilities{
		CIM_ElementCapabilities: tmp,
	}
	return
}

func NewMLNX_DriverElementCapabilitiesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MLNX_DriverElementCapabilities, err error) {
	tmp, err := NewCIM_ElementCapabilitiesEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MLNX_DriverElementCapabilities{
		CIM_ElementCapabilities: tmp,
	}
	return
}
