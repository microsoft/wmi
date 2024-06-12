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

// Msvm_IDEController struct
type Msvm_IDEController struct {
	*CIM_IDEController
}

func NewMsvm_IDEControllerEx1(instance *cim.WmiInstance) (newInstance *Msvm_IDEController, err error) {
	tmp, err := NewCIM_IDEControllerEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_IDEController{
		CIM_IDEController: tmp,
	}
	return
}

func NewMsvm_IDEControllerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_IDEController, err error) {
	tmp, err := NewCIM_IDEControllerEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_IDEController{
		CIM_IDEController: tmp,
	}
	return
}
