// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.Microsoft.Windows.ManagementTools
//////////////////////////////////////////////
package managementtools

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_MTEventRecord struct
type MSFT_MTEventRecord struct {
	*cim.WmiInstance

	//
	Description string

	//
	Id uint32

	//
	Level uint16

	//
	Log string

	//
	RecordId uint32

	//
	Source string

	//
	Timestamp string
}

func NewMSFT_MTEventRecordEx1(instance *cim.WmiInstance) (newInstance *MSFT_MTEventRecord, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_MTEventRecord{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_MTEventRecordEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_MTEventRecord, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_MTEventRecord{
		WmiInstance: tmp,
	}
	return
}

// SetDescription sets the value of Description for the instance
func (instance *MSFT_MTEventRecord) SetPropertyDescription(value string) (err error) {
	return instance.SetProperty("Description", (value))
}

// GetDescription gets the value of Description for the instance
func (instance *MSFT_MTEventRecord) GetPropertyDescription() (value string, err error) {
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

// SetId sets the value of Id for the instance
func (instance *MSFT_MTEventRecord) SetPropertyId(value uint32) (err error) {
	return instance.SetProperty("Id", (value))
}

// GetId gets the value of Id for the instance
func (instance *MSFT_MTEventRecord) GetPropertyId() (value uint32, err error) {
	retValue, err := instance.GetProperty("Id")
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

// SetLevel sets the value of Level for the instance
func (instance *MSFT_MTEventRecord) SetPropertyLevel(value uint16) (err error) {
	return instance.SetProperty("Level", (value))
}

// GetLevel gets the value of Level for the instance
func (instance *MSFT_MTEventRecord) GetPropertyLevel() (value uint16, err error) {
	retValue, err := instance.GetProperty("Level")
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

// SetLog sets the value of Log for the instance
func (instance *MSFT_MTEventRecord) SetPropertyLog(value string) (err error) {
	return instance.SetProperty("Log", (value))
}

// GetLog gets the value of Log for the instance
func (instance *MSFT_MTEventRecord) GetPropertyLog() (value string, err error) {
	retValue, err := instance.GetProperty("Log")
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

// SetRecordId sets the value of RecordId for the instance
func (instance *MSFT_MTEventRecord) SetPropertyRecordId(value uint32) (err error) {
	return instance.SetProperty("RecordId", (value))
}

// GetRecordId gets the value of RecordId for the instance
func (instance *MSFT_MTEventRecord) GetPropertyRecordId() (value uint32, err error) {
	retValue, err := instance.GetProperty("RecordId")
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

// SetSource sets the value of Source for the instance
func (instance *MSFT_MTEventRecord) SetPropertySource(value string) (err error) {
	return instance.SetProperty("Source", (value))
}

// GetSource gets the value of Source for the instance
func (instance *MSFT_MTEventRecord) GetPropertySource() (value string, err error) {
	retValue, err := instance.GetProperty("Source")
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

// SetTimestamp sets the value of Timestamp for the instance
func (instance *MSFT_MTEventRecord) SetPropertyTimestamp(value string) (err error) {
	return instance.SetProperty("Timestamp", (value))
}

// GetTimestamp gets the value of Timestamp for the instance
func (instance *MSFT_MTEventRecord) GetPropertyTimestamp() (value string, err error) {
	retValue, err := instance.GetProperty("Timestamp")
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
