// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WMI
//
// ////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSStorageDriver_ScsiInfoExceptions struct
type MSStorageDriver_ScsiInfoExceptions struct {
	*MSStorageDriver

	//
	Active bool

	//
	Flags uint8

	//
	InstanceName string

	//
	IntervalTimer uint32

	//
	MRIE uint8

	//
	Padding uint8

	//
	PageSavable bool

	//
	ReportCount uint32
}

func NewMSStorageDriver_ScsiInfoExceptionsEx1(instance *cim.WmiInstance) (newInstance *MSStorageDriver_ScsiInfoExceptions, err error) {
	tmp, err := NewMSStorageDriverEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSStorageDriver_ScsiInfoExceptions{
		MSStorageDriver: tmp,
	}
	return
}

func NewMSStorageDriver_ScsiInfoExceptionsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSStorageDriver_ScsiInfoExceptions, err error) {
	tmp, err := NewMSStorageDriverEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSStorageDriver_ScsiInfoExceptions{
		MSStorageDriver: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSStorageDriver_ScsiInfoExceptions) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSStorageDriver_ScsiInfoExceptions) GetPropertyActive() (value bool, err error) {
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

// SetFlags sets the value of Flags for the instance
func (instance *MSStorageDriver_ScsiInfoExceptions) SetPropertyFlags(value uint8) (err error) {
	return instance.SetProperty("Flags", (value))
}

// GetFlags gets the value of Flags for the instance
func (instance *MSStorageDriver_ScsiInfoExceptions) GetPropertyFlags() (value uint8, err error) {
	retValue, err := instance.GetProperty("Flags")
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSStorageDriver_ScsiInfoExceptions) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSStorageDriver_ScsiInfoExceptions) GetPropertyInstanceName() (value string, err error) {
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

// SetIntervalTimer sets the value of IntervalTimer for the instance
func (instance *MSStorageDriver_ScsiInfoExceptions) SetPropertyIntervalTimer(value uint32) (err error) {
	return instance.SetProperty("IntervalTimer", (value))
}

// GetIntervalTimer gets the value of IntervalTimer for the instance
func (instance *MSStorageDriver_ScsiInfoExceptions) GetPropertyIntervalTimer() (value uint32, err error) {
	retValue, err := instance.GetProperty("IntervalTimer")
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

// SetMRIE sets the value of MRIE for the instance
func (instance *MSStorageDriver_ScsiInfoExceptions) SetPropertyMRIE(value uint8) (err error) {
	return instance.SetProperty("MRIE", (value))
}

// GetMRIE gets the value of MRIE for the instance
func (instance *MSStorageDriver_ScsiInfoExceptions) GetPropertyMRIE() (value uint8, err error) {
	retValue, err := instance.GetProperty("MRIE")
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

// SetPadding sets the value of Padding for the instance
func (instance *MSStorageDriver_ScsiInfoExceptions) SetPropertyPadding(value uint8) (err error) {
	return instance.SetProperty("Padding", (value))
}

// GetPadding gets the value of Padding for the instance
func (instance *MSStorageDriver_ScsiInfoExceptions) GetPropertyPadding() (value uint8, err error) {
	retValue, err := instance.GetProperty("Padding")
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

// SetPageSavable sets the value of PageSavable for the instance
func (instance *MSStorageDriver_ScsiInfoExceptions) SetPropertyPageSavable(value bool) (err error) {
	return instance.SetProperty("PageSavable", (value))
}

// GetPageSavable gets the value of PageSavable for the instance
func (instance *MSStorageDriver_ScsiInfoExceptions) GetPropertyPageSavable() (value bool, err error) {
	retValue, err := instance.GetProperty("PageSavable")
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

// SetReportCount sets the value of ReportCount for the instance
func (instance *MSStorageDriver_ScsiInfoExceptions) SetPropertyReportCount(value uint32) (err error) {
	return instance.SetProperty("ReportCount", (value))
}

// GetReportCount gets the value of ReportCount for the instance
func (instance *MSStorageDriver_ScsiInfoExceptions) GetPropertyReportCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("ReportCount")
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
