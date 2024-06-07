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

// WSAT_TraceEvent struct
type WSAT_TraceEvent struct { 
	*WSAT_TraceProvider
}

	func NewWSAT_TraceEventEx1(instance *cim.WmiInstance) (newInstance *WSAT_TraceEvent, err error) {tmp, err := NewWSAT_TraceProviderEx1(instance)
		
	if err != nil { return }
	newInstance = &WSAT_TraceEvent {
	WSAT_TraceProvider: tmp,
	}
	return
	}
	

	func NewWSAT_TraceEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *WSAT_TraceEvent, err error) {tmp, err := NewWSAT_TraceProviderEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &WSAT_TraceEvent {
	WSAT_TraceProvider: tmp,
	}
	return
	}
	

