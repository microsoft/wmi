// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.CIMV2.power
//////////////////////////////////////////////
package power
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_PowerSettingDefinition struct
type Win32_PowerSettingDefinition struct { 
	*CIM_Capabilities
}

	func NewWin32_PowerSettingDefinitionEx1(instance *cim.WmiInstance) (newInstance *Win32_PowerSettingDefinition, err error) {tmp, err := NewCIM_CapabilitiesEx1(instance)
		
	if err != nil { return }
	newInstance = &Win32_PowerSettingDefinition {
	CIM_Capabilities: tmp,
	}
	return
	}
	

	func NewWin32_PowerSettingDefinitionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Win32_PowerSettingDefinition, err error) {tmp, err := NewCIM_CapabilitiesEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Win32_PowerSettingDefinition {
	CIM_Capabilities: tmp,
	}
	return
	}
	

