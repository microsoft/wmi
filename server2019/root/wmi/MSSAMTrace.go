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

// MSSAMTrace struct
type MSSAMTrace struct { 
	*EventTrace
}

	func NewMSSAMTraceEx1(instance *cim.WmiInstance) (newInstance *MSSAMTrace, err error) {tmp, err := NewEventTraceEx1(instance)
		
	if err != nil { return }
	newInstance = &MSSAMTrace {
	EventTrace: tmp,
	}
	return
	}
	

	func NewMSSAMTraceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSSAMTrace, err error) {tmp, err := NewEventTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSSAMTrace {
	EventTrace: tmp,
	}
	return
	}
	

