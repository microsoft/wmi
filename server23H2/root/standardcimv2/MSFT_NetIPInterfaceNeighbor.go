// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NetIPInterfaceNeighbor struct
type MSFT_NetIPInterfaceNeighbor struct {
	*CIM_BindsToLANEndpoint
}

func NewMSFT_NetIPInterfaceNeighborEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetIPInterfaceNeighbor, err error) {
	tmp, err := NewCIM_BindsToLANEndpointEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetIPInterfaceNeighbor{
		CIM_BindsToLANEndpoint: tmp,
	}
	return
}

func NewMSFT_NetIPInterfaceNeighborEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetIPInterfaceNeighbor, err error) {
	tmp, err := NewCIM_BindsToLANEndpointEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetIPInterfaceNeighbor{
		CIM_BindsToLANEndpoint: tmp,
	}
	return
}
