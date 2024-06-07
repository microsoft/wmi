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

// Msvm_TimeSyncComponent struct
type Msvm_TimeSyncComponent struct { 
	*CIM_LogicalDevice
}

	func NewMsvm_TimeSyncComponentEx1(instance *cim.WmiInstance) (newInstance *Msvm_TimeSyncComponent, err error) {tmp, err := NewCIM_LogicalDeviceEx1(instance)
		
	if err != nil { return }
	newInstance = &Msvm_TimeSyncComponent {
	CIM_LogicalDevice: tmp,
	}
	return
	}
	

	func NewMsvm_TimeSyncComponentEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Msvm_TimeSyncComponent, err error) {tmp, err := NewCIM_LogicalDeviceEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Msvm_TimeSyncComponent {
	CIM_LogicalDevice: tmp,
	}
	return
	}
	

