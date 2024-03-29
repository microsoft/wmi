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

// MLNX_SoftwareElementCapabilities struct
type MLNX_SoftwareElementCapabilities struct {
	*CIM_ElementCapabilities
}

func NewMLNX_SoftwareElementCapabilitiesEx1(instance *cim.WmiInstance) (newInstance *MLNX_SoftwareElementCapabilities, err error) {
	tmp, err := NewCIM_ElementCapabilitiesEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MLNX_SoftwareElementCapabilities{
		CIM_ElementCapabilities: tmp,
	}
	return
}

func NewMLNX_SoftwareElementCapabilitiesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MLNX_SoftwareElementCapabilities, err error) {
	tmp, err := NewCIM_ElementCapabilitiesEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MLNX_SoftwareElementCapabilities{
		CIM_ElementCapabilities: tmp,
	}
	return
}
