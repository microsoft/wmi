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

// MSFT_NetUDPEndpoint struct
type MSFT_NetUDPEndpoint struct {
	*MSFT_NetTransportConnection
}

func NewMSFT_NetUDPEndpointEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetUDPEndpoint, err error) {
	tmp, err := NewMSFT_NetTransportConnectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetUDPEndpoint{
		MSFT_NetTransportConnection: tmp,
	}
	return
}

func NewMSFT_NetUDPEndpointEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetUDPEndpoint, err error) {
	tmp, err := NewMSFT_NetTransportConnectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetUDPEndpoint{
		MSFT_NetTransportConnection: tmp,
	}
	return
}
