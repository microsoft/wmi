// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_ElementCapabilities struct
type CIM_ElementCapabilities struct {
	*cim.WmiInstance

	// The Capabilities object associated with the element.
	Capabilities CIM_Capabilities

	// Characteristics provides descriptive information about the Capabilities. when the value 2 "Default" is specified, the associated Capabilities shall represent the default capabilities of the associated Managed Element
	///when the value 2 "Default" is not specified, the Capabilities instance may represent the default capabilities of the Managed Element
	///When the value 3 "Current" is specified, the associated Capabilities shall represent the current capabilities of the associated Managed Element
	///When the value 3 "Current" is not specified, the Capabilities instance may represent the current capabilities of the Managed Element.
	Characteristics []ElementCapabilities_Characteristics

	// The managed element.
	ManagedElement CIM_ManagedElement
}

func NewCIM_ElementCapabilitiesEx1(instance *cim.WmiInstance) (newInstance *CIM_ElementCapabilities, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &CIM_ElementCapabilities{
		WmiInstance: tmp,
	}
	return
}

func NewCIM_ElementCapabilitiesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_ElementCapabilities, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_ElementCapabilities{
		WmiInstance: tmp,
	}
	return
}

// SetCapabilities sets the value of Capabilities for the instance
func (instance *CIM_ElementCapabilities) SetPropertyCapabilities(value CIM_Capabilities) (err error) {
	return instance.SetProperty("Capabilities", (value))
}

// GetCapabilities gets the value of Capabilities for the instance
func (instance *CIM_ElementCapabilities) GetPropertyCapabilities() (value CIM_Capabilities, err error) {
	retValue, err := instance.GetProperty("Capabilities")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(CIM_Capabilities)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " CIM_Capabilities is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = CIM_Capabilities(valuetmp)

	return
}

// SetCharacteristics sets the value of Characteristics for the instance
func (instance *CIM_ElementCapabilities) SetPropertyCharacteristics(value []ElementCapabilities_Characteristics) (err error) {
	return instance.SetProperty("Characteristics", (value))
}

// GetCharacteristics gets the value of Characteristics for the instance
func (instance *CIM_ElementCapabilities) GetPropertyCharacteristics() (value []ElementCapabilities_Characteristics, err error) {
	retValue, err := instance.GetProperty("Characteristics")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(int32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, ElementCapabilities_Characteristics(valuetmp))
	}

	return
}

// SetManagedElement sets the value of ManagedElement for the instance
func (instance *CIM_ElementCapabilities) SetPropertyManagedElement(value CIM_ManagedElement) (err error) {
	return instance.SetProperty("ManagedElement", (value))
}

// GetManagedElement gets the value of ManagedElement for the instance
func (instance *CIM_ElementCapabilities) GetPropertyManagedElement() (value CIM_ManagedElement, err error) {
	retValue, err := instance.GetProperty("ManagedElement")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(CIM_ManagedElement)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " CIM_ManagedElement is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = CIM_ManagedElement(valuetmp)

	return
}
