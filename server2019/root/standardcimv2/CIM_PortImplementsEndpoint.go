// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_PortImplementsEndpoint struct
type CIM_PortImplementsEndpoint struct {
	*CIM_DeviceSAPImplementation
}

func NewCIM_PortImplementsEndpointEx1(instance *cim.WmiInstance) (newInstance *CIM_PortImplementsEndpoint, err error) {
	tmp, err := NewCIM_DeviceSAPImplementationEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_PortImplementsEndpoint{
		CIM_DeviceSAPImplementation: tmp,
	}
	return
}

func NewCIM_PortImplementsEndpointEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_PortImplementsEndpoint, err error) {
	tmp, err := NewCIM_DeviceSAPImplementationEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_PortImplementsEndpoint{
		CIM_DeviceSAPImplementation: tmp,
	}
	return
}
