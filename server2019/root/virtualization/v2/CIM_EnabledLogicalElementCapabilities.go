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
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_EnabledLogicalElementCapabilities struct
type CIM_EnabledLogicalElementCapabilities struct {
	*CIM_Capabilities

	// Boolean indicating whether the ElementName can be modified.
	ElementNameEditSupported bool

	// This string expresses the restrictions on ElementName.The mask is expressed as a regular expression.See DMTF standard ABNF with the Management Profile Specification Usage Guide, appendix C for the regular expression syntax permitted.
	///Since the ElementNameMask can describe the maximum length of the ElementName,any length defined in the regexp is in addition to the restriction defined in MaxElementNameLen (causing the smaller value to be the maximum length) The ElementName value satisfies the restriction, if and only if it matches the regular expression
	ElementNameMask string

	// Maximum supported ElementName length.
	MaxElementNameLen uint16

	// RequestedStatesSupported indicates the possible states that can be requested when using the method RequestStateChange on the EnabledLogicalElement.
	RequestedStatesSupported []EnabledLogicalElementCapabilities_RequestedStatesSupported
}

func NewCIM_EnabledLogicalElementCapabilitiesEx1(instance *cim.WmiInstance) (newInstance *CIM_EnabledLogicalElementCapabilities, err error) {
	tmp, err := NewCIM_CapabilitiesEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_EnabledLogicalElementCapabilities{
		CIM_Capabilities: tmp,
	}
	return
}

func NewCIM_EnabledLogicalElementCapabilitiesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_EnabledLogicalElementCapabilities, err error) {
	tmp, err := NewCIM_CapabilitiesEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_EnabledLogicalElementCapabilities{
		CIM_Capabilities: tmp,
	}
	return
}

// SetElementNameEditSupported sets the value of ElementNameEditSupported for the instance
func (instance *CIM_EnabledLogicalElementCapabilities) SetPropertyElementNameEditSupported(value bool) (err error) {
	return instance.SetProperty("ElementNameEditSupported", (value))
}

// GetElementNameEditSupported gets the value of ElementNameEditSupported for the instance
func (instance *CIM_EnabledLogicalElementCapabilities) GetPropertyElementNameEditSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("ElementNameEditSupported")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetElementNameMask sets the value of ElementNameMask for the instance
func (instance *CIM_EnabledLogicalElementCapabilities) SetPropertyElementNameMask(value string) (err error) {
	return instance.SetProperty("ElementNameMask", (value))
}

// GetElementNameMask gets the value of ElementNameMask for the instance
func (instance *CIM_EnabledLogicalElementCapabilities) GetPropertyElementNameMask() (value string, err error) {
	retValue, err := instance.GetProperty("ElementNameMask")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetMaxElementNameLen sets the value of MaxElementNameLen for the instance
func (instance *CIM_EnabledLogicalElementCapabilities) SetPropertyMaxElementNameLen(value uint16) (err error) {
	return instance.SetProperty("MaxElementNameLen", (value))
}

// GetMaxElementNameLen gets the value of MaxElementNameLen for the instance
func (instance *CIM_EnabledLogicalElementCapabilities) GetPropertyMaxElementNameLen() (value uint16, err error) {
	retValue, err := instance.GetProperty("MaxElementNameLen")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetRequestedStatesSupported sets the value of RequestedStatesSupported for the instance
func (instance *CIM_EnabledLogicalElementCapabilities) SetPropertyRequestedStatesSupported(value []EnabledLogicalElementCapabilities_RequestedStatesSupported) (err error) {
	return instance.SetProperty("RequestedStatesSupported", (value))
}

// GetRequestedStatesSupported gets the value of RequestedStatesSupported for the instance
func (instance *CIM_EnabledLogicalElementCapabilities) GetPropertyRequestedStatesSupported() (value []EnabledLogicalElementCapabilities_RequestedStatesSupported, err error) {
	retValue, err := instance.GetProperty("RequestedStatesSupported")
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
		value = append(value, EnabledLogicalElementCapabilities_RequestedStatesSupported(valuetmp))
	}

	return
}
