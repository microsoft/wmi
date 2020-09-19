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

// MLNX_FirmwareElementCapabilities struct
type MLNX_FirmwareElementCapabilities struct {
	*CIM_ElementCapabilities
}

func NewMLNX_FirmwareElementCapabilitiesEx1(instance *cim.WmiInstance) (newInstance *MLNX_FirmwareElementCapabilities, err error) {
	tmp, err := NewCIM_ElementCapabilitiesEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MLNX_FirmwareElementCapabilities{
		CIM_ElementCapabilities: tmp,
	}
	return
}

func NewMLNX_FirmwareElementCapabilitiesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MLNX_FirmwareElementCapabilities, err error) {
	tmp, err := NewCIM_ElementCapabilitiesEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MLNX_FirmwareElementCapabilities{
		CIM_ElementCapabilities: tmp,
	}
	return
}
