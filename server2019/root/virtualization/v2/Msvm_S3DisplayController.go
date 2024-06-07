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

// Msvm_S3DisplayController struct
type Msvm_S3DisplayController struct { 
	*CIM_DisplayController
}

	func NewMsvm_S3DisplayControllerEx1(instance *cim.WmiInstance) (newInstance *Msvm_S3DisplayController, err error) {tmp, err := NewCIM_DisplayControllerEx1(instance)
		
	if err != nil { return }
	newInstance = &Msvm_S3DisplayController {
	CIM_DisplayController: tmp,
	}
	return
	}
	

	func NewMsvm_S3DisplayControllerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Msvm_S3DisplayController, err error) {tmp, err := NewCIM_DisplayControllerEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Msvm_S3DisplayController {
	CIM_DisplayController: tmp,
	}
	return
	}
	

