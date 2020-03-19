// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_Synthetic3DDisplayController struct
type Msvm_Synthetic3DDisplayController struct {
	*CIM_DisplayController

	//
	AllocatedGPU string
}

func NewMsvm_Synthetic3DDisplayControllerEx1(instance *cim.WmiInstance) (newInstance *Msvm_Synthetic3DDisplayController, err error) {
	tmp, err := NewCIM_DisplayControllerEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_Synthetic3DDisplayController{
		CIM_DisplayController: tmp,
	}
	return
}

func NewMsvm_Synthetic3DDisplayControllerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_Synthetic3DDisplayController, err error) {
	tmp, err := NewCIM_DisplayControllerEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_Synthetic3DDisplayController{
		CIM_DisplayController: tmp,
	}
	return
}

// SetAllocatedGPU sets the value of AllocatedGPU for the instance
func (instance *Msvm_Synthetic3DDisplayController) SetPropertyAllocatedGPU(value string) (err error) {
	return instance.SetProperty("AllocatedGPU", value)
}

// GetAllocatedGPU gets the value of AllocatedGPU for the instance
func (instance *Msvm_Synthetic3DDisplayController) GetPropertyAllocatedGPU() (value string, err error) {
	retValue, err := instance.GetProperty("AllocatedGPU")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
