// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.Microsoft.Windows.ServerManager
//////////////////////////////////////////////
package servermanager

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// ServerComponent_DirectAccess_VPN struct
type ServerComponent_DirectAccess_VPN struct {
	*MSFT_ServerManagerServerComponentDescriptor
}

func NewServerComponent_DirectAccess_VPNEx1(instance *cim.WmiInstance) (newInstance *ServerComponent_DirectAccess_VPN, err error) {
	tmp, err := NewMSFT_ServerManagerServerComponentDescriptorEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ServerComponent_DirectAccess_VPN{
		MSFT_ServerManagerServerComponentDescriptor: tmp,
	}
	return
}

func NewServerComponent_DirectAccess_VPNEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ServerComponent_DirectAccess_VPN, err error) {
	tmp, err := NewMSFT_ServerManagerServerComponentDescriptorEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ServerComponent_DirectAccess_VPN{
		MSFT_ServerManagerServerComponentDescriptor: tmp,
	}
	return
}
