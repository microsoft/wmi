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

// Msvm_MediaPresent struct
type Msvm_MediaPresent struct { 
	*CIM_MediaPresent
}

	func NewMsvm_MediaPresentEx1(instance *cim.WmiInstance) (newInstance *Msvm_MediaPresent, err error) {tmp, err := NewCIM_MediaPresentEx1(instance)
		
	if err != nil { return }
	newInstance = &Msvm_MediaPresent {
	CIM_MediaPresent: tmp,
	}
	return
	}
	

	func NewMsvm_MediaPresentEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Msvm_MediaPresent, err error) {tmp, err := NewCIM_MediaPresentEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Msvm_MediaPresent {
	CIM_MediaPresent: tmp,
	}
	return
	}
	

