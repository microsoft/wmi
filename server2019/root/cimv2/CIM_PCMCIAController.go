// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_PCMCIAController struct
type CIM_PCMCIAController struct {
	*CIM_Controller

	//
	Manufacturer string
}

func NewCIM_PCMCIAControllerEx1(instance *cim.WmiInstance) (newInstance *CIM_PCMCIAController, err error) {
	tmp, err := NewCIM_ControllerEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_PCMCIAController{
		CIM_Controller: tmp,
	}
	return
}

func NewCIM_PCMCIAControllerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_PCMCIAController, err error) {
	tmp, err := NewCIM_ControllerEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_PCMCIAController{
		CIM_Controller: tmp,
	}
	return
}

// SetManufacturer sets the value of Manufacturer for the instance
func (instance *CIM_PCMCIAController) SetPropertyManufacturer(value string) (err error) {
	return instance.SetProperty("Manufacturer", value)
}

// GetManufacturer gets the value of Manufacturer for the instance
func (instance *CIM_PCMCIAController) GetPropertyManufacturer() (value string, err error) {
	retValue, err := instance.GetProperty("Manufacturer")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
