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

// CIM_ParticipatesInSet struct
type CIM_ParticipatesInSet struct {
	cim.WmiInstance

	//
	Element CIM_PhysicalElement

	//
	Set CIM_ReplacementSet
}

// SetElement sets the value of Element for the instance
func (instance *CIM_ParticipatesInSet) SetPropertyElement(value CIM_PhysicalElement) (err error) {
	return instance.SetProperty("Element", value)
}

// GetElement gets the value of Element for the instance
func (instance *CIM_ParticipatesInSet) GetPropertyElement() (value CIM_PhysicalElement, err error) {
	retValue, err := instance.GetProperty("Element")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_PhysicalElement)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSet sets the value of Set for the instance
func (instance *CIM_ParticipatesInSet) SetPropertySet(value CIM_ReplacementSet) (err error) {
	return instance.SetProperty("Set", value)
}

// GetSet gets the value of Set for the instance
func (instance *CIM_ParticipatesInSet) GetPropertySet() (value CIM_ReplacementSet, err error) {
	retValue, err := instance.GetProperty("Set")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_ReplacementSet)
	if !ok {
		// TODO: Set an error
	}
	return
}
