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

// CompCS struct
type CompCS struct { 
	*Thread_V2
}

	func NewCompCSEx1(instance *cim.WmiInstance) (newInstance *CompCS, err error) {tmp, err := NewThread_V2Ex1(instance)
		
	if err != nil { return }
	newInstance = &CompCS {
	Thread_V2: tmp,
	}
	return
	}
	

	func NewCompCSEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *CompCS, err error) {tmp, err := NewThread_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &CompCS {
	Thread_V2: tmp,
	}
	return
	}
	

