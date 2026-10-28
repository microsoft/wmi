// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 10/28/2026 using wmigen
//      Source root.Microsoft.Windows.ServerManager
//////////////////////////////////////////////
package servermanager

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// ServerComponent_WindowsAdminCenterSetup struct
type ServerComponent_WindowsAdminCenterSetup struct {
	*MSFT_ServerManagerServerComponentDescriptor
}

func NewServerComponent_WindowsAdminCenterSetupEx1(instance *cim.WmiInstance) (newInstance *ServerComponent_WindowsAdminCenterSetup, err error) {
	tmp, err := NewMSFT_ServerManagerServerComponentDescriptorEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ServerComponent_WindowsAdminCenterSetup{
		MSFT_ServerManagerServerComponentDescriptor: tmp,
	}
	return
}

func NewServerComponent_WindowsAdminCenterSetupEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ServerComponent_WindowsAdminCenterSetup, err error) {
	tmp, err := NewMSFT_ServerManagerServerComponentDescriptorEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ServerComponent_WindowsAdminCenterSetup{
		MSFT_ServerManagerServerComponentDescriptor: tmp,
	}
	return
}
