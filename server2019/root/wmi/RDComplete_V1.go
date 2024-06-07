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

// RDComplete_V1 struct
type RDComplete_V1 struct { 
	*EventTraceEvent_V1
}

	func NewRDComplete_V1Ex1(instance *cim.WmiInstance) (newInstance *RDComplete_V1, err error) {tmp, err := NewEventTraceEvent_V1Ex1(instance)
		
	if err != nil { return }
	newInstance = &RDComplete_V1 {
	EventTraceEvent_V1: tmp,
	}
	return
	}
	

	func NewRDComplete_V1Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *RDComplete_V1, err error) {tmp, err := NewEventTraceEvent_V1Ex6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &RDComplete_V1 {
	EventTraceEvent_V1: tmp,
	}
	return
	}
	

