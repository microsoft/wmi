// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.Microsoft.Windows.HardwareManagement
//////////////////////////////////////////////
package hardwaremanagement

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_RecordForLog struct
type CIM_RecordForLog struct {
	*CIM_ManagedElement

	//
	Locale string

	//
	PerceivedSeverity uint16

	//
	RecordData string

	//
	RecordFormat string
}

func NewCIM_RecordForLogEx1(instance *cim.WmiInstance) (newInstance *CIM_RecordForLog, err error) {
	tmp, err := NewCIM_ManagedElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_RecordForLog{
		CIM_ManagedElement: tmp,
	}
	return
}

func NewCIM_RecordForLogEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_RecordForLog, err error) {
	tmp, err := NewCIM_ManagedElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_RecordForLog{
		CIM_ManagedElement: tmp,
	}
	return
}

// SetLocale sets the value of Locale for the instance
func (instance *CIM_RecordForLog) SetPropertyLocale(value string) (err error) {
	return instance.SetProperty("Locale", (value))
}

// GetLocale gets the value of Locale for the instance
func (instance *CIM_RecordForLog) GetPropertyLocale() (value string, err error) {
	retValue, err := instance.GetProperty("Locale")
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

// SetPerceivedSeverity sets the value of PerceivedSeverity for the instance
func (instance *CIM_RecordForLog) SetPropertyPerceivedSeverity(value uint16) (err error) {
	return instance.SetProperty("PerceivedSeverity", (value))
}

// GetPerceivedSeverity gets the value of PerceivedSeverity for the instance
func (instance *CIM_RecordForLog) GetPropertyPerceivedSeverity() (value uint16, err error) {
	retValue, err := instance.GetProperty("PerceivedSeverity")
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

// SetRecordData sets the value of RecordData for the instance
func (instance *CIM_RecordForLog) SetPropertyRecordData(value string) (err error) {
	return instance.SetProperty("RecordData", (value))
}

// GetRecordData gets the value of RecordData for the instance
func (instance *CIM_RecordForLog) GetPropertyRecordData() (value string, err error) {
	retValue, err := instance.GetProperty("RecordData")
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

// SetRecordFormat sets the value of RecordFormat for the instance
func (instance *CIM_RecordForLog) SetPropertyRecordFormat(value string) (err error) {
	return instance.SetProperty("RecordFormat", (value))
}

// GetRecordFormat gets the value of RecordFormat for the instance
func (instance *CIM_RecordForLog) GetPropertyRecordFormat() (value string, err error) {
	retValue, err := instance.GetProperty("RecordFormat")
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
