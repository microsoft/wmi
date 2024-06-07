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

// SetInfoPolicy_Start struct
type SetInfoPolicy_Start struct { 
	*SetInfoPolicy
}

	func NewSetInfoPolicy_StartEx1(instance *cim.WmiInstance) (newInstance *SetInfoPolicy_Start, err error) {tmp, err := NewSetInfoPolicyEx1(instance)
		
	if err != nil { return }
	newInstance = &SetInfoPolicy_Start {
	SetInfoPolicy: tmp,
	}
	return
	}
	

	func NewSetInfoPolicy_StartEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *SetInfoPolicy_Start, err error) {tmp, err := NewSetInfoPolicyEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &SetInfoPolicy_Start {
	SetInfoPolicy: tmp,
	}
	return
	}
	

