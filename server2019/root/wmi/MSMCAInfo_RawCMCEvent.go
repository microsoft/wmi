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

// MSMCAInfo_RawCMCEvent struct
type MSMCAInfo_RawCMCEvent struct {
	*WMIEvent

	//
	Active bool

	//
	Count uint32

	//
	InstanceName string

	//
	Records []MSMCAInfo_Entry
}

func NewMSMCAInfo_RawCMCEventEx1(instance *cim.WmiInstance) (newInstance *MSMCAInfo_RawCMCEvent, err error) {
	tmp, err := NewWMIEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSMCAInfo_RawCMCEvent{
		WMIEvent: tmp,
	}
	return
}

func NewMSMCAInfo_RawCMCEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSMCAInfo_RawCMCEvent, err error) {
	tmp, err := NewWMIEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSMCAInfo_RawCMCEvent{
		WMIEvent: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSMCAInfo_RawCMCEvent) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSMCAInfo_RawCMCEvent) GetPropertyActive() (value bool, err error) {
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

// SetCount sets the value of Count for the instance
func (instance *MSMCAInfo_RawCMCEvent) SetPropertyCount(value uint32) (err error) {
	return instance.SetProperty("Count", (value))
}

// GetCount gets the value of Count for the instance
func (instance *MSMCAInfo_RawCMCEvent) GetPropertyCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("Count")
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
func (instance *MSMCAInfo_RawCMCEvent) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSMCAInfo_RawCMCEvent) GetPropertyInstanceName() (value string, err error) {
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

// SetRecords sets the value of Records for the instance
func (instance *MSMCAInfo_RawCMCEvent) SetPropertyRecords(value []MSMCAInfo_Entry) (err error) {
	return instance.SetProperty("Records", (value))
}

// GetRecords gets the value of Records for the instance
func (instance *MSMCAInfo_RawCMCEvent) GetPropertyRecords() (value []MSMCAInfo_Entry, err error) {
	retValue, err := instance.GetProperty("Records")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(MSMCAInfo_Entry)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " MSMCAInfo_Entry is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, MSMCAInfo_Entry(valuetmp))
	}

	return
}
