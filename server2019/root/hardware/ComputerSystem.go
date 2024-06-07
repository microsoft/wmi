// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.Hardware
//////////////////////////////////////////////
package hardware
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// ComputerSystem struct
type ComputerSystem struct { 
	*CIM_ComputerSystem
}

	func NewComputerSystemEx1(instance *cim.WmiInstance) (newInstance *ComputerSystem, err error) {tmp, err := NewCIM_ComputerSystemEx1(instance)
		
	if err != nil { return }
	newInstance = &ComputerSystem {
	CIM_ComputerSystem: tmp,
	}
	return
	}
	

	func NewComputerSystemEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ComputerSystem, err error) {tmp, err := NewCIM_ComputerSystemEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ComputerSystem {
	CIM_ComputerSystem: tmp,
	}
	return
	}
	

