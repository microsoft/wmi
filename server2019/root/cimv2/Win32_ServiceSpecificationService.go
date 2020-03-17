// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_ServiceSpecificationService struct
type Win32_ServiceSpecificationService struct {
	cim.WmiInstance

	//
	Check Win32_ServiceSpecification

	//
	Element Win32_Service
}

// SetCheck sets the value of Check for the instance
func (instance *Win32_ServiceSpecificationService) SetPropertyCheck(value Win32_ServiceSpecification) (err error) {
	return instance.SetProperty("Check", value)
}

// GetCheck gets the value of Check for the instance
func (instance *Win32_ServiceSpecificationService) GetPropertyCheck() (value Win32_ServiceSpecification, err error) {
	retValue, err := instance.GetProperty("Check")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_ServiceSpecification)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetElement sets the value of Element for the instance
func (instance *Win32_ServiceSpecificationService) SetPropertyElement(value Win32_Service) (err error) {
	return instance.SetProperty("Element", value)
}

// GetElement gets the value of Element for the instance
func (instance *Win32_ServiceSpecificationService) GetPropertyElement() (value Win32_Service, err error) {
	retValue, err := instance.GetProperty("Element")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_Service)
	if !ok {
		// TODO: Set an error
	}
	return
}
