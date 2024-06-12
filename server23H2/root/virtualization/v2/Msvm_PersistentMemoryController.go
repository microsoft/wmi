// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_PersistentMemoryController struct
type Msvm_PersistentMemoryController struct {
	*CIM_Controller
}

func NewMsvm_PersistentMemoryControllerEx1(instance *cim.WmiInstance) (newInstance *Msvm_PersistentMemoryController, err error) {
	tmp, err := NewCIM_ControllerEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_PersistentMemoryController{
		CIM_Controller: tmp,
	}
	return
}

func NewMsvm_PersistentMemoryControllerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_PersistentMemoryController, err error) {
	tmp, err := NewCIM_ControllerEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_PersistentMemoryController{
		CIM_Controller: tmp,
	}
	return
}
