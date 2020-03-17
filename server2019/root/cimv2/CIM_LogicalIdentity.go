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

// CIM_LogicalIdentity struct
type CIM_LogicalIdentity struct {
	cim.WmiInstance

	//
	SameElement CIM_LogicalElement

	//
	SystemElement CIM_LogicalElement
}

// SetSameElement sets the value of SameElement for the instance
func (instance *CIM_LogicalIdentity) SetPropertySameElement(value CIM_LogicalElement) (err error) {
	return instance.SetProperty("SameElement", value)
}

// GetSameElement gets the value of SameElement for the instance
func (instance *CIM_LogicalIdentity) GetPropertySameElement() (value CIM_LogicalElement, err error) {
	retValue, err := instance.GetProperty("SameElement")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_LogicalElement)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSystemElement sets the value of SystemElement for the instance
func (instance *CIM_LogicalIdentity) SetPropertySystemElement(value CIM_LogicalElement) (err error) {
	return instance.SetProperty("SystemElement", value)
}

// GetSystemElement gets the value of SystemElement for the instance
func (instance *CIM_LogicalIdentity) GetPropertySystemElement() (value CIM_LogicalElement, err error) {
	retValue, err := instance.GetProperty("SystemElement")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_LogicalElement)
	if !ok {
		// TODO: Set an error
	}
	return
}
