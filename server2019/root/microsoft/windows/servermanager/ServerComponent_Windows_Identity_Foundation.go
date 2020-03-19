// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.ServerManager
//////////////////////////////////////////////
package servermanager

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// ServerComponent_Windows_Identity_Foundation struct
type ServerComponent_Windows_Identity_Foundation struct {
	*MSFT_ServerManagerServerComponentDescriptor
}

func NewServerComponent_Windows_Identity_FoundationEx1(instance *cim.WmiInstance) (newInstance *ServerComponent_Windows_Identity_Foundation, err error) {
	tmp, err := NewMSFT_ServerManagerServerComponentDescriptorEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ServerComponent_Windows_Identity_Foundation{
		MSFT_ServerManagerServerComponentDescriptor: tmp,
	}
	return
}

func NewServerComponent_Windows_Identity_FoundationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ServerComponent_Windows_Identity_Foundation, err error) {
	tmp, err := NewMSFT_ServerManagerServerComponentDescriptorEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ServerComponent_Windows_Identity_Foundation{
		MSFT_ServerManagerServerComponentDescriptor: tmp,
	}
	return
}
