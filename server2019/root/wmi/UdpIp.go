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

// UdpIp struct
type UdpIp struct { 
	*MSNT_SystemTrace
}

	func NewUdpIpEx1(instance *cim.WmiInstance) (newInstance *UdpIp, err error) {tmp, err := NewMSNT_SystemTraceEx1(instance)
		
	if err != nil { return }
	newInstance = &UdpIp {
	MSNT_SystemTrace: tmp,
	}
	return
	}
	

	func NewUdpIpEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *UdpIp, err error) {tmp, err := NewMSNT_SystemTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &UdpIp {
	MSNT_SystemTrace: tmp,
	}
	return
	}
	

