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

// ServerComponent_FS_iSCSITarget_Server struct
type ServerComponent_FS_iSCSITarget_Server struct {
	*MSFT_ServerManagerServerComponentDescriptor
}

func NewServerComponent_FS_iSCSITarget_ServerEx1(instance *cim.WmiInstance) (newInstance *ServerComponent_FS_iSCSITarget_Server, err error) {
	tmp, err := NewMSFT_ServerManagerServerComponentDescriptorEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ServerComponent_FS_iSCSITarget_Server{
		MSFT_ServerManagerServerComponentDescriptor: tmp,
	}
	return
}

func NewServerComponent_FS_iSCSITarget_ServerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ServerComponent_FS_iSCSITarget_Server, err error) {
	tmp, err := NewMSFT_ServerManagerServerComponentDescriptorEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ServerComponent_FS_iSCSITarget_Server{
		MSFT_ServerManagerServerComponentDescriptor: tmp,
	}
	return
}
