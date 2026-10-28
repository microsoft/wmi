// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 10/28/2026 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// LENOVO_GAMEZONE_THERMAL_MODE_EVENT struct
type LENOVO_GAMEZONE_THERMAL_MODE_EVENT struct {
	*WMIEvent

	//
	Active bool

	// Thermal mode change event
	EventId uint32

	//
	InstanceName string
}

func NewLENOVO_GAMEZONE_THERMAL_MODE_EVENTEx1(instance *cim.WmiInstance) (newInstance *LENOVO_GAMEZONE_THERMAL_MODE_EVENT, err error) {
	tmp, err := NewWMIEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &LENOVO_GAMEZONE_THERMAL_MODE_EVENT{
		WMIEvent: tmp,
	}
	return
}

func NewLENOVO_GAMEZONE_THERMAL_MODE_EVENTEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *LENOVO_GAMEZONE_THERMAL_MODE_EVENT, err error) {
	tmp, err := NewWMIEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &LENOVO_GAMEZONE_THERMAL_MODE_EVENT{
		WMIEvent: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *LENOVO_GAMEZONE_THERMAL_MODE_EVENT) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *LENOVO_GAMEZONE_THERMAL_MODE_EVENT) GetPropertyActive() (value bool, err error) {
	retValue, err := instance.GetProperty("Active")
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

// SetEventId sets the value of EventId for the instance
func (instance *LENOVO_GAMEZONE_THERMAL_MODE_EVENT) SetPropertyEventId(value uint32) (err error) {
	return instance.SetProperty("EventId", (value))
}

// GetEventId gets the value of EventId for the instance
func (instance *LENOVO_GAMEZONE_THERMAL_MODE_EVENT) GetPropertyEventId() (value uint32, err error) {
	retValue, err := instance.GetProperty("EventId")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetInstanceName sets the value of InstanceName for the instance
func (instance *LENOVO_GAMEZONE_THERMAL_MODE_EVENT) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *LENOVO_GAMEZONE_THERMAL_MODE_EVENT) GetPropertyInstanceName() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceName")
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
