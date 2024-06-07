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

// SystemSpecificCollection struct
type SystemSpecificCollection struct { 
	*CIM_SystemSpecificCollection
}

	func NewSystemSpecificCollectionEx1(instance *cim.WmiInstance) (newInstance *SystemSpecificCollection, err error) {tmp, err := NewCIM_SystemSpecificCollectionEx1(instance)
		
	if err != nil { return }
	newInstance = &SystemSpecificCollection {
	CIM_SystemSpecificCollection: tmp,
	}
	return
	}
	

	func NewSystemSpecificCollectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *SystemSpecificCollection, err error) {tmp, err := NewCIM_SystemSpecificCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &SystemSpecificCollection {
	CIM_SystemSpecificCollection: tmp,
	}
	return
	}
	

