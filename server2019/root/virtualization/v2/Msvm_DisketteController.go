// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_DisketteController struct
type Msvm_DisketteController struct {
	*CIM_Controller
}

func NewMsvm_DisketteControllerEx1(instance *cim.WmiInstance) (newInstance *Msvm_DisketteController, err error) {
	tmp, err := NewCIM_ControllerEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_DisketteController{
		CIM_Controller: tmp,
	}
	return
}

func NewMsvm_DisketteControllerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_DisketteController, err error) {
	tmp, err := NewCIM_ControllerEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_DisketteController{
		CIM_Controller: tmp,
	}
	return
}
