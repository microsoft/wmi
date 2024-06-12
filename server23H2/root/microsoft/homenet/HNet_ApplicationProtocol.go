// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.Microsoft.HomeNet
//////////////////////////////////////////////
package homenet

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// HNet_ApplicationProtocol struct
type HNet_ApplicationProtocol struct {
	*cim.WmiInstance

	//
	BuiltIn bool

	//
	Enabled bool

	//
	Id string

	//
	Name string

	//
	OutgoingIPProtocol uint8

	//
	OutgoingPort uint16

	//
	ResponseArray []HNet_ResponseRange

	//
	ResponseCount uint16
}

func NewHNet_ApplicationProtocolEx1(instance *cim.WmiInstance) (newInstance *HNet_ApplicationProtocol, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &HNet_ApplicationProtocol{
		WmiInstance: tmp,
	}
	return
}

func NewHNet_ApplicationProtocolEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *HNet_ApplicationProtocol, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &HNet_ApplicationProtocol{
		WmiInstance: tmp,
	}
	return
}

// SetBuiltIn sets the value of BuiltIn for the instance
func (instance *HNet_ApplicationProtocol) SetPropertyBuiltIn(value bool) (err error) {
	return instance.SetProperty("BuiltIn", (value))
}

// GetBuiltIn gets the value of BuiltIn for the instance
func (instance *HNet_ApplicationProtocol) GetPropertyBuiltIn() (value bool, err error) {
	retValue, err := instance.GetProperty("BuiltIn")
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

// SetEnabled sets the value of Enabled for the instance
func (instance *HNet_ApplicationProtocol) SetPropertyEnabled(value bool) (err error) {
	return instance.SetProperty("Enabled", (value))
}

// GetEnabled gets the value of Enabled for the instance
func (instance *HNet_ApplicationProtocol) GetPropertyEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("Enabled")
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

// SetId sets the value of Id for the instance
func (instance *HNet_ApplicationProtocol) SetPropertyId(value string) (err error) {
	return instance.SetProperty("Id", (value))
}

// GetId gets the value of Id for the instance
func (instance *HNet_ApplicationProtocol) GetPropertyId() (value string, err error) {
	retValue, err := instance.GetProperty("Id")
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

// SetName sets the value of Name for the instance
func (instance *HNet_ApplicationProtocol) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *HNet_ApplicationProtocol) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
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

// SetOutgoingIPProtocol sets the value of OutgoingIPProtocol for the instance
func (instance *HNet_ApplicationProtocol) SetPropertyOutgoingIPProtocol(value uint8) (err error) {
	return instance.SetProperty("OutgoingIPProtocol", (value))
}

// GetOutgoingIPProtocol gets the value of OutgoingIPProtocol for the instance
func (instance *HNet_ApplicationProtocol) GetPropertyOutgoingIPProtocol() (value uint8, err error) {
	retValue, err := instance.GetProperty("OutgoingIPProtocol")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetOutgoingPort sets the value of OutgoingPort for the instance
func (instance *HNet_ApplicationProtocol) SetPropertyOutgoingPort(value uint16) (err error) {
	return instance.SetProperty("OutgoingPort", (value))
}

// GetOutgoingPort gets the value of OutgoingPort for the instance
func (instance *HNet_ApplicationProtocol) GetPropertyOutgoingPort() (value uint16, err error) {
	retValue, err := instance.GetProperty("OutgoingPort")
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

// SetResponseArray sets the value of ResponseArray for the instance
func (instance *HNet_ApplicationProtocol) SetPropertyResponseArray(value []HNet_ResponseRange) (err error) {
	return instance.SetProperty("ResponseArray", (value))
}

// GetResponseArray gets the value of ResponseArray for the instance
func (instance *HNet_ApplicationProtocol) GetPropertyResponseArray() (value []HNet_ResponseRange, err error) {
	retValue, err := instance.GetProperty("ResponseArray")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(HNet_ResponseRange)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " HNet_ResponseRange is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, HNet_ResponseRange(valuetmp))
	}

	return
}

// SetResponseCount sets the value of ResponseCount for the instance
func (instance *HNet_ApplicationProtocol) SetPropertyResponseCount(value uint16) (err error) {
	return instance.SetProperty("ResponseCount", (value))
}

// GetResponseCount gets the value of ResponseCount for the instance
func (instance *HNet_ApplicationProtocol) GetPropertyResponseCount() (value uint16, err error) {
	retValue, err := instance.GetProperty("ResponseCount")
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
