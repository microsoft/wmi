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

// Msvm_DisketteDrive struct
type Msvm_DisketteDrive struct { 
	*CIM_DisketteDrive
}

	func NewMsvm_DisketteDriveEx1(instance *cim.WmiInstance) (newInstance *Msvm_DisketteDrive, err error) {tmp, err := NewCIM_DisketteDriveEx1(instance)
		
	if err != nil { return }
	newInstance = &Msvm_DisketteDrive {
	CIM_DisketteDrive: tmp,
	}
	return
	}
	

	func NewMsvm_DisketteDriveEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Msvm_DisketteDrive, err error) {tmp, err := NewCIM_DisketteDriveEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Msvm_DisketteDrive {
	CIM_DisketteDrive: tmp,
	}
	return
	}
	

