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

// HeapTraceProvider struct
type HeapTraceProvider struct { 
	*EventTrace
}

	func NewHeapTraceProviderEx1(instance *cim.WmiInstance) (newInstance *HeapTraceProvider, err error) {tmp, err := NewEventTraceEx1(instance)
		
	if err != nil { return }
	newInstance = &HeapTraceProvider {
	EventTrace: tmp,
	}
	return
	}
	

	func NewHeapTraceProviderEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *HeapTraceProvider, err error) {tmp, err := NewEventTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &HeapTraceProvider {
	EventTrace: tmp,
	}
	return
	}
	

