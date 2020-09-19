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

// ServerComponent_BITS_Compact_Server struct
type ServerComponent_BITS_Compact_Server struct {
	*MSFT_ServerManagerServerComponentDescriptor
}

func NewServerComponent_BITS_Compact_ServerEx1(instance *cim.WmiInstance) (newInstance *ServerComponent_BITS_Compact_Server, err error) {
	tmp, err := NewMSFT_ServerManagerServerComponentDescriptorEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ServerComponent_BITS_Compact_Server{
		MSFT_ServerManagerServerComponentDescriptor: tmp,
	}
	return
}

func NewServerComponent_BITS_Compact_ServerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ServerComponent_BITS_Compact_Server, err error) {
	tmp, err := NewMSFT_ServerManagerServerComponentDescriptorEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ServerComponent_BITS_Compact_Server{
		MSFT_ServerManagerServerComponentDescriptor: tmp,
	}
	return
}
