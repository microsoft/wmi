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

// PowerEvents struct
type PowerEvents struct { 
	*MSNT_SystemTrace
}

	func NewPowerEventsEx1(instance *cim.WmiInstance) (newInstance *PowerEvents, err error) {tmp, err := NewMSNT_SystemTraceEx1(instance)
		
	if err != nil { return }
	newInstance = &PowerEvents {
	MSNT_SystemTrace: tmp,
	}
	return
	}
	

	func NewPowerEventsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *PowerEvents, err error) {tmp, err := NewMSNT_SystemTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &PowerEvents {
	MSNT_SystemTrace: tmp,
	}
	return
	}
	

