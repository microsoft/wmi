// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// EVENT_DATA_IR struct
type EVENT_DATA_IR struct {
	*WMIEvent

	//
	Active bool

	//
	ctrlId uint32

	//
	eventClassLocale uint32

	//
	eventCode uint32

	//
	InstanceName string

	//
	seqNumber uint32
}

func NewEVENT_DATA_IREx1(instance *cim.WmiInstance) (newInstance *EVENT_DATA_IR, err error) {
	tmp, err := NewWMIEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &EVENT_DATA_IR{
		WMIEvent: tmp,
	}
	return
}

func NewEVENT_DATA_IREx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *EVENT_DATA_IR, err error) {
	tmp, err := NewWMIEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &EVENT_DATA_IR{
		WMIEvent: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *EVENT_DATA_IR) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *EVENT_DATA_IR) GetPropertyActive() (value bool, err error) {
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

// SetctrlId sets the value of ctrlId for the instance
func (instance *EVENT_DATA_IR) SetPropertyctrlId(value uint32) (err error) {
	return instance.SetProperty("ctrlId", (value))
}

// GetctrlId gets the value of ctrlId for the instance
func (instance *EVENT_DATA_IR) GetPropertyctrlId() (value uint32, err error) {
	retValue, err := instance.GetProperty("ctrlId")
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

// SeteventClassLocale sets the value of eventClassLocale for the instance
func (instance *EVENT_DATA_IR) SetPropertyeventClassLocale(value uint32) (err error) {
	return instance.SetProperty("eventClassLocale", (value))
}

// GeteventClassLocale gets the value of eventClassLocale for the instance
func (instance *EVENT_DATA_IR) GetPropertyeventClassLocale() (value uint32, err error) {
	retValue, err := instance.GetProperty("eventClassLocale")
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

// SeteventCode sets the value of eventCode for the instance
func (instance *EVENT_DATA_IR) SetPropertyeventCode(value uint32) (err error) {
	return instance.SetProperty("eventCode", (value))
}

// GeteventCode gets the value of eventCode for the instance
func (instance *EVENT_DATA_IR) GetPropertyeventCode() (value uint32, err error) {
	retValue, err := instance.GetProperty("eventCode")
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
func (instance *EVENT_DATA_IR) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *EVENT_DATA_IR) GetPropertyInstanceName() (value string, err error) {
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

// SetseqNumber sets the value of seqNumber for the instance
func (instance *EVENT_DATA_IR) SetPropertyseqNumber(value uint32) (err error) {
	return instance.SetProperty("seqNumber", (value))
}

// GetseqNumber gets the value of seqNumber for the instance
func (instance *EVENT_DATA_IR) GetPropertyseqNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("seqNumber")
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
