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

// Close_Start struct
type Close_Start struct { 
	*Close
}

	func NewClose_StartEx1(instance *cim.WmiInstance) (newInstance *Close_Start, err error) {tmp, err := NewCloseEx1(instance)
		
	if err != nil { return }
	newInstance = &Close_Start {
	Close: tmp,
	}
	return
	}
	

	func NewClose_StartEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Close_Start, err error) {tmp, err := NewCloseEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Close_Start {
	Close: tmp,
	}
	return
	}
	

