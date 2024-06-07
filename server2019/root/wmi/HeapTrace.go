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

// HeapTrace struct
type HeapTrace struct { 
	*HeapTraceProvider
}

	func NewHeapTraceEx1(instance *cim.WmiInstance) (newInstance *HeapTrace, err error) {tmp, err := NewHeapTraceProviderEx1(instance)
		
	if err != nil { return }
	newInstance = &HeapTrace {
	HeapTraceProvider: tmp,
	}
	return
	}
	

	func NewHeapTraceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *HeapTrace, err error) {tmp, err := NewHeapTraceProviderEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &HeapTrace {
	HeapTraceProvider: tmp,
	}
	return
	}
	

