// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.Microsoft.Windows.ServerManager
//////////////////////////////////////////////
package servermanager

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// ServerComponent_MSMQ_Directory struct
type ServerComponent_MSMQ_Directory struct {
	*MSFT_ServerManagerServerComponentDescriptor
}

func NewServerComponent_MSMQ_DirectoryEx1(instance *cim.WmiInstance) (newInstance *ServerComponent_MSMQ_Directory, err error) {
	tmp, err := NewMSFT_ServerManagerServerComponentDescriptorEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ServerComponent_MSMQ_Directory{
		MSFT_ServerManagerServerComponentDescriptor: tmp,
	}
	return
}

func NewServerComponent_MSMQ_DirectoryEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ServerComponent_MSMQ_Directory, err error) {
	tmp, err := NewMSFT_ServerManagerServerComponentDescriptorEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ServerComponent_MSMQ_Directory{
		MSFT_ServerManagerServerComponentDescriptor: tmp,
	}
	return
}
