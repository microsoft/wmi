// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// TcpIp_V0 struct
type TcpIp_V0 struct { 
	*MSNT_SystemTrace
}

	func NewTcpIp_V0Ex1(instance *cim.WmiInstance) (newInstance *TcpIp_V0, err error) {tmp, err := NewMSNT_SystemTraceEx1(instance)
		
	if err != nil { return }
	newInstance = &TcpIp_V0 {
	MSNT_SystemTrace: tmp,
	}
	return
	}
	

	func NewTcpIp_V0Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *TcpIp_V0, err error) {tmp, err := NewMSNT_SystemTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &TcpIp_V0 {
	MSNT_SystemTrace: tmp,
	}
	return
	}
	

