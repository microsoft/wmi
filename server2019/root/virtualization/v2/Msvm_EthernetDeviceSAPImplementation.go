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

// Msvm_EthernetDeviceSAPImplementation struct
type Msvm_EthernetDeviceSAPImplementation struct { 
	*CIM_DeviceSAPImplementation
}

	func NewMsvm_EthernetDeviceSAPImplementationEx1(instance *cim.WmiInstance) (newInstance *Msvm_EthernetDeviceSAPImplementation, err error) {tmp, err := NewCIM_DeviceSAPImplementationEx1(instance)
		
	if err != nil { return }
	newInstance = &Msvm_EthernetDeviceSAPImplementation {
	CIM_DeviceSAPImplementation: tmp,
	}
	return
	}
	

	func NewMsvm_EthernetDeviceSAPImplementationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Msvm_EthernetDeviceSAPImplementation, err error) {tmp, err := NewCIM_DeviceSAPImplementationEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Msvm_EthernetDeviceSAPImplementation {
	CIM_DeviceSAPImplementation: tmp,
	}
	return
	}
	

