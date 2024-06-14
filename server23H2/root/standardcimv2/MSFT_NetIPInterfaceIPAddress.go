// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NetIPInterfaceIPAddress struct
type MSFT_NetIPInterfaceIPAddress struct {
	*CIM_BindsToLANEndpoint
}

func NewMSFT_NetIPInterfaceIPAddressEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetIPInterfaceIPAddress, err error) {
	tmp, err := NewCIM_BindsToLANEndpointEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetIPInterfaceIPAddress{
		CIM_BindsToLANEndpoint: tmp,
	}
	return
}

func NewMSFT_NetIPInterfaceIPAddressEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetIPInterfaceIPAddress, err error) {
	tmp, err := NewCIM_BindsToLANEndpointEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetIPInterfaceIPAddress{
		CIM_BindsToLANEndpoint: tmp,
	}
	return
}
