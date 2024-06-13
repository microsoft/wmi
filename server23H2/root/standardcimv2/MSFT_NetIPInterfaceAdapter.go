// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NetIPInterfaceAdapter struct
type MSFT_NetIPInterfaceAdapter struct {
	*CIM_PortImplementsEndpoint
}

func NewMSFT_NetIPInterfaceAdapterEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetIPInterfaceAdapter, err error) {
	tmp, err := NewCIM_PortImplementsEndpointEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetIPInterfaceAdapter{
		CIM_PortImplementsEndpoint: tmp,
	}
	return
}

func NewMSFT_NetIPInterfaceAdapterEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetIPInterfaceAdapter, err error) {
	tmp, err := NewCIM_PortImplementsEndpointEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetIPInterfaceAdapter{
		CIM_PortImplementsEndpoint: tmp,
	}
	return
}
