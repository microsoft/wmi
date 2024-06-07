// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.Microsoft.Windows.ServerManager
//////////////////////////////////////////////
package servermanager
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// ServerComponent_RDC struct
type ServerComponent_RDC struct { 
	*MSFT_ServerManagerServerComponentDescriptor
}

	func NewServerComponent_RDCEx1(instance *cim.WmiInstance) (newInstance *ServerComponent_RDC, err error) {tmp, err := NewMSFT_ServerManagerServerComponentDescriptorEx1(instance)
		
	if err != nil { return }
	newInstance = &ServerComponent_RDC {
	MSFT_ServerManagerServerComponentDescriptor: tmp,
	}
	return
	}
	

	func NewServerComponent_RDCEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ServerComponent_RDC, err error) {tmp, err := NewMSFT_ServerManagerServerComponentDescriptorEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ServerComponent_RDC {
	MSFT_ServerManagerServerComponentDescriptor: tmp,
	}
	return
	}
	

