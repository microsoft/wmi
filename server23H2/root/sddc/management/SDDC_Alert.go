// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.SDDC.Management
//////////////////////////////////////////////
package management

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// SDDC_Alert struct
type SDDC_Alert struct {
	*cim.WmiInstance

	//
	Actions string

	//
	Description string

	//
	FaultingObjectDescription string

	//
	FaultingObjectLocation string

	//
	Id string

	//
	Severity uint16

	//
	Time string

	//
	Title string

	//
	Type string
}

func NewSDDC_AlertEx1(instance *cim.WmiInstance) (newInstance *SDDC_Alert, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &SDDC_Alert{
		WmiInstance: tmp,
	}
	return
}

func NewSDDC_AlertEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SDDC_Alert, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SDDC_Alert{
		WmiInstance: tmp,
	}
	return
}

// SetActions sets the value of Actions for the instance
func (instance *SDDC_Alert) SetPropertyActions(value string) (err error) {
	return instance.SetProperty("Actions", (value))
}

// GetActions gets the value of Actions for the instance
func (instance *SDDC_Alert) GetPropertyActions() (value string, err error) {
	retValue, err := instance.GetProperty("Actions")
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

// SetDescription sets the value of Description for the instance
func (instance *SDDC_Alert) SetPropertyDescription(value string) (err error) {
	return instance.SetProperty("Description", (value))
}

// GetDescription gets the value of Description for the instance
func (instance *SDDC_Alert) GetPropertyDescription() (value string, err error) {
	retValue, err := instance.GetProperty("Description")
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

// SetFaultingObjectDescription sets the value of FaultingObjectDescription for the instance
func (instance *SDDC_Alert) SetPropertyFaultingObjectDescription(value string) (err error) {
	return instance.SetProperty("FaultingObjectDescription", (value))
}

// GetFaultingObjectDescription gets the value of FaultingObjectDescription for the instance
func (instance *SDDC_Alert) GetPropertyFaultingObjectDescription() (value string, err error) {
	retValue, err := instance.GetProperty("FaultingObjectDescription")
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

// SetFaultingObjectLocation sets the value of FaultingObjectLocation for the instance
func (instance *SDDC_Alert) SetPropertyFaultingObjectLocation(value string) (err error) {
	return instance.SetProperty("FaultingObjectLocation", (value))
}

// GetFaultingObjectLocation gets the value of FaultingObjectLocation for the instance
func (instance *SDDC_Alert) GetPropertyFaultingObjectLocation() (value string, err error) {
	retValue, err := instance.GetProperty("FaultingObjectLocation")
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

// SetId sets the value of Id for the instance
func (instance *SDDC_Alert) SetPropertyId(value string) (err error) {
	return instance.SetProperty("Id", (value))
}

// GetId gets the value of Id for the instance
func (instance *SDDC_Alert) GetPropertyId() (value string, err error) {
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

// SetSeverity sets the value of Severity for the instance
func (instance *SDDC_Alert) SetPropertySeverity(value uint16) (err error) {
	return instance.SetProperty("Severity", (value))
}

// GetSeverity gets the value of Severity for the instance
func (instance *SDDC_Alert) GetPropertySeverity() (value uint16, err error) {
	retValue, err := instance.GetProperty("Severity")
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

// SetTime sets the value of Time for the instance
func (instance *SDDC_Alert) SetPropertyTime(value string) (err error) {
	return instance.SetProperty("Time", (value))
}

// GetTime gets the value of Time for the instance
func (instance *SDDC_Alert) GetPropertyTime() (value string, err error) {
	retValue, err := instance.GetProperty("Time")
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

// SetTitle sets the value of Title for the instance
func (instance *SDDC_Alert) SetPropertyTitle(value string) (err error) {
	return instance.SetProperty("Title", (value))
}

// GetTitle gets the value of Title for the instance
func (instance *SDDC_Alert) GetPropertyTitle() (value string, err error) {
	retValue, err := instance.GetProperty("Title")
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

// SetType sets the value of Type for the instance
func (instance *SDDC_Alert) SetPropertyType(value string) (err error) {
	return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *SDDC_Alert) GetPropertyType() (value string, err error) {
	retValue, err := instance.GetProperty("Type")
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
