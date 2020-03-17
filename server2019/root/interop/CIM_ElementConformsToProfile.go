// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Interop
//////////////////////////////////////////////
package interop

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_ElementConformsToProfile struct
type CIM_ElementConformsToProfile struct {
	cim.WmiInstance

	//
	ConformantStandard CIM_RegisteredProfile

	//
	ManagedElement CIM_ManagedElement
}

// SetConformantStandard sets the value of ConformantStandard for the instance
func (instance *CIM_ElementConformsToProfile) SetPropertyConformantStandard(value CIM_RegisteredProfile) (err error) {
	return instance.SetProperty("ConformantStandard", value)
}

// GetConformantStandard gets the value of ConformantStandard for the instance
func (instance *CIM_ElementConformsToProfile) GetPropertyConformantStandard() (value CIM_RegisteredProfile, err error) {
	retValue, err := instance.GetProperty("ConformantStandard")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_RegisteredProfile)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetManagedElement sets the value of ManagedElement for the instance
func (instance *CIM_ElementConformsToProfile) SetPropertyManagedElement(value CIM_ManagedElement) (err error) {
	return instance.SetProperty("ManagedElement", value)
}

// GetManagedElement gets the value of ManagedElement for the instance
func (instance *CIM_ElementConformsToProfile) GetPropertyManagedElement() (value CIM_ManagedElement, err error) {
	retValue, err := instance.GetProperty("ManagedElement")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_ManagedElement)
	if !ok {
		// TODO: Set an error
	}
	return
}
