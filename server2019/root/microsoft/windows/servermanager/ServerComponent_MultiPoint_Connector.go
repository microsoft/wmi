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

// ServerComponent_MultiPoint_Connector struct
type ServerComponent_MultiPoint_Connector struct {
	*MSFT_ServerManagerServerComponentDescriptor
}

func NewServerComponent_MultiPoint_ConnectorEx1(instance *cim.WmiInstance) (newInstance *ServerComponent_MultiPoint_Connector, err error) {
	tmp, err := NewMSFT_ServerManagerServerComponentDescriptorEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ServerComponent_MultiPoint_Connector{
		MSFT_ServerManagerServerComponentDescriptor: tmp,
	}
	return
}

func NewServerComponent_MultiPoint_ConnectorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ServerComponent_MultiPoint_Connector, err error) {
	tmp, err := NewMSFT_ServerManagerServerComponentDescriptorEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ServerComponent_MultiPoint_Connector{
		MSFT_ServerManagerServerComponentDescriptor: tmp,
	}
	return
}
