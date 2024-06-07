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

// CritSecTrace struct
type CritSecTrace struct { 
	*CritSecTraceProvider
}

	func NewCritSecTraceEx1(instance *cim.WmiInstance) (newInstance *CritSecTrace, err error) {tmp, err := NewCritSecTraceProviderEx1(instance)
		
	if err != nil { return }
	newInstance = &CritSecTrace {
	CritSecTraceProvider: tmp,
	}
	return
	}
	

	func NewCritSecTraceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *CritSecTrace, err error) {tmp, err := NewCritSecTraceProviderEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &CritSecTrace {
	CritSecTraceProvider: tmp,
	}
	return
	}
	

