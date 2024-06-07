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
 "github.com/microsoft/wmi/pkg/base/instance"
)

// EventTrace struct
type EventTrace struct { 
	*cim.WmiInstance
}

	func NewEventTraceEx1(instance *cim.WmiInstance) (newInstance *EventTrace, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &EventTrace {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewEventTraceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *EventTrace, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &EventTrace {
	WmiInstance: tmp,
	}
	return
	}
	

