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

// MSiSCSI_Eventlog struct
type MSiSCSI_Eventlog struct {
	*__ExtrinsicEvent

	//
	Active bool

	// Additional data to include in eventlog message, typically iSCSI Header
	AdditionalData []uint8

	//
	InstanceName string

	// If zero then this event is not logged to system eventlog
	LogToEventlog uint32

	// Size of Additional Data
	Size uint32

	// Type of eventlog message
	Type Eventlog_Type
}

func NewMSiSCSI_EventlogEx1(instance *cim.WmiInstance) (newInstance *MSiSCSI_Eventlog, err error) {
	tmp, err := New__ExtrinsicEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSiSCSI_Eventlog{
		__ExtrinsicEvent: tmp,
	}
	return
}

func NewMSiSCSI_EventlogEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSiSCSI_Eventlog, err error) {
	tmp, err := New__ExtrinsicEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSiSCSI_Eventlog{
		__ExtrinsicEvent: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSiSCSI_Eventlog) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSiSCSI_Eventlog) GetPropertyActive() (value bool, err error) {
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

// SetAdditionalData sets the value of AdditionalData for the instance
func (instance *MSiSCSI_Eventlog) SetPropertyAdditionalData(value []uint8) (err error) {
	return instance.SetProperty("AdditionalData", (value))
}

// GetAdditionalData gets the value of AdditionalData for the instance
func (instance *MSiSCSI_Eventlog) GetPropertyAdditionalData() (value []uint8, err error) {
	retValue, err := instance.GetProperty("AdditionalData")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint8)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint8(valuetmp))
	}

	return
}

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSiSCSI_Eventlog) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSiSCSI_Eventlog) GetPropertyInstanceName() (value string, err error) {
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

// SetLogToEventlog sets the value of LogToEventlog for the instance
func (instance *MSiSCSI_Eventlog) SetPropertyLogToEventlog(value uint32) (err error) {
	return instance.SetProperty("LogToEventlog", (value))
}

// GetLogToEventlog gets the value of LogToEventlog for the instance
func (instance *MSiSCSI_Eventlog) GetPropertyLogToEventlog() (value uint32, err error) {
	retValue, err := instance.GetProperty("LogToEventlog")
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

// SetSize sets the value of Size for the instance
func (instance *MSiSCSI_Eventlog) SetPropertySize(value uint32) (err error) {
	return instance.SetProperty("Size", (value))
}

// GetSize gets the value of Size for the instance
func (instance *MSiSCSI_Eventlog) GetPropertySize() (value uint32, err error) {
	retValue, err := instance.GetProperty("Size")
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

// SetType sets the value of Type for the instance
func (instance *MSiSCSI_Eventlog) SetPropertyType(value Eventlog_Type) (err error) {
	return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *MSiSCSI_Eventlog) GetPropertyType() (value Eventlog_Type, err error) {
	retValue, err := instance.GetProperty("Type")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = Eventlog_Type(valuetmp)

	return
}
