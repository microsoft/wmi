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

// LookupSids_End struct
type LookupSids_End struct { 
	*LookupSids
}

	func NewLookupSids_EndEx1(instance *cim.WmiInstance) (newInstance *LookupSids_End, err error) {tmp, err := NewLookupSidsEx1(instance)
		
	if err != nil { return }
	newInstance = &LookupSids_End {
	LookupSids: tmp,
	}
	return
	}
	

	func NewLookupSids_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *LookupSids_End, err error) {tmp, err := NewLookupSidsEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &LookupSids_End {
	LookupSids: tmp,
	}
	return
	}
	

