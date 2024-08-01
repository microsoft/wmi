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

// ServerComponent_FS_SMB1_CLIENT struct
type ServerComponent_FS_SMB1_CLIENT struct {
	*MSFT_ServerManagerServerComponentDescriptor
}

func NewServerComponent_FS_SMB1_CLIENTEx1(instance *cim.WmiInstance) (newInstance *ServerComponent_FS_SMB1_CLIENT, err error) {
	tmp, err := NewMSFT_ServerManagerServerComponentDescriptorEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ServerComponent_FS_SMB1_CLIENT{
		MSFT_ServerManagerServerComponentDescriptor: tmp,
	}
	return
}

func NewServerComponent_FS_SMB1_CLIENTEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ServerComponent_FS_SMB1_CLIENT, err error) {
	tmp, err := NewMSFT_ServerManagerServerComponentDescriptorEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ServerComponent_FS_SMB1_CLIENT{
		MSFT_ServerManagerServerComponentDescriptor: tmp,
	}
	return
}
