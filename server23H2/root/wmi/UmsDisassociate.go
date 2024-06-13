// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// UmsDisassociate struct
type UmsDisassociate struct {
	*UmsEvent

	//
	PrimaryThreadId uint32

	//
	ProcessId uint32

	//
	ScheduledThreadId uint32

	//
	Status uint32

	//
	UmsApcControlFlags uint32
}

func NewUmsDisassociateEx1(instance *cim.WmiInstance) (newInstance *UmsDisassociate, err error) {
	tmp, err := NewUmsEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &UmsDisassociate{
		UmsEvent: tmp,
	}
	return
}

func NewUmsDisassociateEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *UmsDisassociate, err error) {
	tmp, err := NewUmsEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &UmsDisassociate{
		UmsEvent: tmp,
	}
	return
}

// SetPrimaryThreadId sets the value of PrimaryThreadId for the instance
func (instance *UmsDisassociate) SetPropertyPrimaryThreadId(value uint32) (err error) {
	return instance.SetProperty("PrimaryThreadId", (value))
}

// GetPrimaryThreadId gets the value of PrimaryThreadId for the instance
func (instance *UmsDisassociate) GetPropertyPrimaryThreadId() (value uint32, err error) {
	retValue, err := instance.GetProperty("PrimaryThreadId")
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

// SetProcessId sets the value of ProcessId for the instance
func (instance *UmsDisassociate) SetPropertyProcessId(value uint32) (err error) {
	return instance.SetProperty("ProcessId", (value))
}

// GetProcessId gets the value of ProcessId for the instance
func (instance *UmsDisassociate) GetPropertyProcessId() (value uint32, err error) {
	retValue, err := instance.GetProperty("ProcessId")
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

// SetScheduledThreadId sets the value of ScheduledThreadId for the instance
func (instance *UmsDisassociate) SetPropertyScheduledThreadId(value uint32) (err error) {
	return instance.SetProperty("ScheduledThreadId", (value))
}

// GetScheduledThreadId gets the value of ScheduledThreadId for the instance
func (instance *UmsDisassociate) GetPropertyScheduledThreadId() (value uint32, err error) {
	retValue, err := instance.GetProperty("ScheduledThreadId")
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

// SetStatus sets the value of Status for the instance
func (instance *UmsDisassociate) SetPropertyStatus(value uint32) (err error) {
	return instance.SetProperty("Status", (value))
}

// GetStatus gets the value of Status for the instance
func (instance *UmsDisassociate) GetPropertyStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("Status")
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

// SetUmsApcControlFlags sets the value of UmsApcControlFlags for the instance
func (instance *UmsDisassociate) SetPropertyUmsApcControlFlags(value uint32) (err error) {
	return instance.SetProperty("UmsApcControlFlags", (value))
}

// GetUmsApcControlFlags gets the value of UmsApcControlFlags for the instance
func (instance *UmsDisassociate) GetPropertyUmsApcControlFlags() (value uint32, err error) {
	retValue, err := instance.GetProperty("UmsApcControlFlags")
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
