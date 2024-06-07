// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_FcActiveConnection struct
type Msvm_FcActiveConnection struct { 
	*CIM_ActiveConnection
}

	func NewMsvm_FcActiveConnectionEx1(instance *cim.WmiInstance) (newInstance *Msvm_FcActiveConnection, err error) {tmp, err := NewCIM_ActiveConnectionEx1(instance)
		
	if err != nil { return }
	newInstance = &Msvm_FcActiveConnection {
	CIM_ActiveConnection: tmp,
	}
	return
	}
	

	func NewMsvm_FcActiveConnectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Msvm_FcActiveConnection, err error) {tmp, err := NewCIM_ActiveConnectionEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Msvm_FcActiveConnection {
	CIM_ActiveConnection: tmp,
	}
	return
	}
	

