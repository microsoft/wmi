// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSRedbook_Performance struct
type MSRedbook_Performance struct {
	*cim.WmiInstance

	//
	Active bool

	//
	DataProcessed int64

	//
	InstanceName string

	//
	StreamPausedCount uint32

	//
	TimeReadDelay int64

	//
	TimeReading int64

	//
	TimeStreamDelay int64

	//
	TimeStreaming int64
}

func NewMSRedbook_PerformanceEx1(instance *cim.WmiInstance) (newInstance *MSRedbook_Performance, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSRedbook_Performance{
		WmiInstance: tmp,
	}
	return
}

func NewMSRedbook_PerformanceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSRedbook_Performance, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSRedbook_Performance{
		WmiInstance: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSRedbook_Performance) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSRedbook_Performance) GetPropertyActive() (value bool, err error) {
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

// SetDataProcessed sets the value of DataProcessed for the instance
func (instance *MSRedbook_Performance) SetPropertyDataProcessed(value int64) (err error) {
	return instance.SetProperty("DataProcessed", (value))
}

// GetDataProcessed gets the value of DataProcessed for the instance
func (instance *MSRedbook_Performance) GetPropertyDataProcessed() (value int64, err error) {
	retValue, err := instance.GetProperty("DataProcessed")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSRedbook_Performance) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSRedbook_Performance) GetPropertyInstanceName() (value string, err error) {
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

// SetStreamPausedCount sets the value of StreamPausedCount for the instance
func (instance *MSRedbook_Performance) SetPropertyStreamPausedCount(value uint32) (err error) {
	return instance.SetProperty("StreamPausedCount", (value))
}

// GetStreamPausedCount gets the value of StreamPausedCount for the instance
func (instance *MSRedbook_Performance) GetPropertyStreamPausedCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("StreamPausedCount")
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

// SetTimeReadDelay sets the value of TimeReadDelay for the instance
func (instance *MSRedbook_Performance) SetPropertyTimeReadDelay(value int64) (err error) {
	return instance.SetProperty("TimeReadDelay", (value))
}

// GetTimeReadDelay gets the value of TimeReadDelay for the instance
func (instance *MSRedbook_Performance) GetPropertyTimeReadDelay() (value int64, err error) {
	retValue, err := instance.GetProperty("TimeReadDelay")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}

// SetTimeReading sets the value of TimeReading for the instance
func (instance *MSRedbook_Performance) SetPropertyTimeReading(value int64) (err error) {
	return instance.SetProperty("TimeReading", (value))
}

// GetTimeReading gets the value of TimeReading for the instance
func (instance *MSRedbook_Performance) GetPropertyTimeReading() (value int64, err error) {
	retValue, err := instance.GetProperty("TimeReading")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}

// SetTimeStreamDelay sets the value of TimeStreamDelay for the instance
func (instance *MSRedbook_Performance) SetPropertyTimeStreamDelay(value int64) (err error) {
	return instance.SetProperty("TimeStreamDelay", (value))
}

// GetTimeStreamDelay gets the value of TimeStreamDelay for the instance
func (instance *MSRedbook_Performance) GetPropertyTimeStreamDelay() (value int64, err error) {
	retValue, err := instance.GetProperty("TimeStreamDelay")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}

// SetTimeStreaming sets the value of TimeStreaming for the instance
func (instance *MSRedbook_Performance) SetPropertyTimeStreaming(value int64) (err error) {
	return instance.SetProperty("TimeStreaming", (value))
}

// GetTimeStreaming gets the value of TimeStreaming for the instance
func (instance *MSRedbook_Performance) GetPropertyTimeStreaming() (value int64, err error) {
	retValue, err := instance.GetProperty("TimeStreaming")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}
