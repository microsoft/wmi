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

// VssTracingEventClass struct
type VssTracingEventClass struct { 
	*VssTracingProvider
}

	func NewVssTracingEventClassEx1(instance *cim.WmiInstance) (newInstance *VssTracingEventClass, err error) {tmp, err := NewVssTracingProviderEx1(instance)
		
	if err != nil { return }
	newInstance = &VssTracingEventClass {
	VssTracingProvider: tmp,
	}
	return
	}
	

	func NewVssTracingEventClassEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *VssTracingEventClass, err error) {tmp, err := NewVssTracingProviderEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &VssTracingEventClass {
	VssTracingProvider: tmp,
	}
	return
	}
	

