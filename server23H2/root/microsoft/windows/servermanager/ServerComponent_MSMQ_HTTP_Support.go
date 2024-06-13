// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.Windows.ServerManager
//////////////////////////////////////////////
package servermanager

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// ServerComponent_MSMQ_HTTP_Support struct
type ServerComponent_MSMQ_HTTP_Support struct {
	*MSFT_ServerManagerServerComponentDescriptor
}

func NewServerComponent_MSMQ_HTTP_SupportEx1(instance *cim.WmiInstance) (newInstance *ServerComponent_MSMQ_HTTP_Support, err error) {
	tmp, err := NewMSFT_ServerManagerServerComponentDescriptorEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ServerComponent_MSMQ_HTTP_Support{
		MSFT_ServerManagerServerComponentDescriptor: tmp,
	}
	return
}

func NewServerComponent_MSMQ_HTTP_SupportEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ServerComponent_MSMQ_HTTP_Support, err error) {
	tmp, err := NewMSFT_ServerManagerServerComponentDescriptorEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ServerComponent_MSMQ_HTTP_Support{
		MSFT_ServerManagerServerComponentDescriptor: tmp,
	}
	return
}
