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

// CIM_SwitchesAmong struct
type CIM_SwitchesAmong struct { 
	*CIM_ForwardsAmong
}

	func NewCIM_SwitchesAmongEx1(instance *cim.WmiInstance) (newInstance *CIM_SwitchesAmong, err error) {tmp, err := NewCIM_ForwardsAmongEx1(instance)
		
	if err != nil { return }
	newInstance = &CIM_SwitchesAmong {
	CIM_ForwardsAmong: tmp,
	}
	return
	}
	

	func NewCIM_SwitchesAmongEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *CIM_SwitchesAmong, err error) {tmp, err := NewCIM_ForwardsAmongEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &CIM_SwitchesAmong {
	CIM_ForwardsAmong: tmp,
	}
	return
	}
	

