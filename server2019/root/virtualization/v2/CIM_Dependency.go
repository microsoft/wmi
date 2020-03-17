// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_Dependency struct
type CIM_Dependency struct {
	cim.WmiInstance

	// Antecedent represents the independent object in this association.
	Antecedent CIM_ManagedElement

	// Dependent represents the object that is dependent on the Antecedent.
	Dependent CIM_ManagedElement
}

// SetAntecedent sets the value of Antecedent for the instance
func (instance *CIM_Dependency) SetPropertyAntecedent(value CIM_ManagedElement) (err error) {
	return instance.SetProperty("Antecedent", value)
}

// GetAntecedent gets the value of Antecedent for the instance
func (instance *CIM_Dependency) GetPropertyAntecedent() (value CIM_ManagedElement, err error) {
	retValue, err := instance.GetProperty("Antecedent")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_ManagedElement)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDependent sets the value of Dependent for the instance
func (instance *CIM_Dependency) SetPropertyDependent(value CIM_ManagedElement) (err error) {
	return instance.SetProperty("Dependent", value)
}

// GetDependent gets the value of Dependent for the instance
func (instance *CIM_Dependency) GetPropertyDependent() (value CIM_ManagedElement, err error) {
	retValue, err := instance.GetProperty("Dependent")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_ManagedElement)
	if !ok {
		// TODO: Set an error
	}
	return
}
