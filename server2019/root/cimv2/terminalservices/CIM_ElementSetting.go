// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.CIMV2.TerminalServices
//////////////////////////////////////////////
package terminalservices
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_ElementSetting struct
type CIM_ElementSetting struct { 
	*CIM_ManagedSystemElement
}

	func NewCIM_ElementSettingEx1(instance *cim.WmiInstance) (newInstance *CIM_ElementSetting, err error) {tmp, err := NewCIM_ManagedSystemElementEx1(instance)
		
	if err != nil { return }
	newInstance = &CIM_ElementSetting {
	CIM_ManagedSystemElement: tmp,
	}
	return
	}
	

	func NewCIM_ElementSettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *CIM_ElementSetting, err error) {tmp, err := NewCIM_ManagedSystemElementEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &CIM_ElementSetting {
	CIM_ManagedSystemElement: tmp,
	}
	return
	}
	

