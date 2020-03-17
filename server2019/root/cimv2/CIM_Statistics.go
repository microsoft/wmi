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

// CIM_Statistics struct
type CIM_Statistics struct {
	cim.WmiInstance

	//
	Element CIM_ManagedSystemElement

	//
	Stats CIM_StatisticalInformation
}

// SetElement sets the value of Element for the instance
func (instance *CIM_Statistics) SetPropertyElement(value CIM_ManagedSystemElement) (err error) {
	return instance.SetProperty("Element", value)
}

// GetElement gets the value of Element for the instance
func (instance *CIM_Statistics) GetPropertyElement() (value CIM_ManagedSystemElement, err error) {
	retValue, err := instance.GetProperty("Element")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_ManagedSystemElement)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStats sets the value of Stats for the instance
func (instance *CIM_Statistics) SetPropertyStats(value CIM_StatisticalInformation) (err error) {
	return instance.SetProperty("Stats", value)
}

// GetStats gets the value of Stats for the instance
func (instance *CIM_Statistics) GetPropertyStats() (value CIM_StatisticalInformation, err error) {
	retValue, err := instance.GetProperty("Stats")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_StatisticalInformation)
	if !ok {
		// TODO: Set an error
	}
	return
}
