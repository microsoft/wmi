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

// ServerComponent_Web_Lgcy_Scripting struct
type ServerComponent_Web_Lgcy_Scripting struct {
	*MSFT_ServerManagerServerComponentDescriptor
}

func NewServerComponent_Web_Lgcy_ScriptingEx1(instance *cim.WmiInstance) (newInstance *ServerComponent_Web_Lgcy_Scripting, err error) {
	tmp, err := NewMSFT_ServerManagerServerComponentDescriptorEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ServerComponent_Web_Lgcy_Scripting{
		MSFT_ServerManagerServerComponentDescriptor: tmp,
	}
	return
}

func NewServerComponent_Web_Lgcy_ScriptingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ServerComponent_Web_Lgcy_Scripting, err error) {
	tmp, err := NewMSFT_ServerManagerServerComponentDescriptorEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ServerComponent_Web_Lgcy_Scripting{
		MSFT_ServerManagerServerComponentDescriptor: tmp,
	}
	return
}
