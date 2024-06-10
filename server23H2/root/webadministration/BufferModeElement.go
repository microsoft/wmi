// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WebAdministration
//
// ////////////////////////////////////////////
package webadministration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// BufferModeElement struct
type BufferModeElement struct {
	*CollectionElement

	//
	MaxBufferSize int32

	//
	MaxBufferThreads int32

	//
	MaxFlushSize int32

	//
	Name string

	//
	RegularFlushInterval string

	//
	UrgentFlushInterval string

	//
	UrgentFlushThreshold int32
}

func NewBufferModeElementEx1(instance *cim.WmiInstance) (newInstance *BufferModeElement, err error) {
	tmp, err := NewCollectionElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &BufferModeElement{
		CollectionElement: tmp,
	}
	return
}

func NewBufferModeElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *BufferModeElement, err error) {
	tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &BufferModeElement{
		CollectionElement: tmp,
	}
	return
}

// SetMaxBufferSize sets the value of MaxBufferSize for the instance
func (instance *BufferModeElement) SetPropertyMaxBufferSize(value int32) (err error) {
	return instance.SetProperty("MaxBufferSize", (value))
}

// GetMaxBufferSize gets the value of MaxBufferSize for the instance
func (instance *BufferModeElement) GetPropertyMaxBufferSize() (value int32, err error) {
	retValue, err := instance.GetProperty("MaxBufferSize")
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

	value = int32(valuetmp)

	return
}

// SetMaxBufferThreads sets the value of MaxBufferThreads for the instance
func (instance *BufferModeElement) SetPropertyMaxBufferThreads(value int32) (err error) {
	return instance.SetProperty("MaxBufferThreads", (value))
}

// GetMaxBufferThreads gets the value of MaxBufferThreads for the instance
func (instance *BufferModeElement) GetPropertyMaxBufferThreads() (value int32, err error) {
	retValue, err := instance.GetProperty("MaxBufferThreads")
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

	value = int32(valuetmp)

	return
}

// SetMaxFlushSize sets the value of MaxFlushSize for the instance
func (instance *BufferModeElement) SetPropertyMaxFlushSize(value int32) (err error) {
	return instance.SetProperty("MaxFlushSize", (value))
}

// GetMaxFlushSize gets the value of MaxFlushSize for the instance
func (instance *BufferModeElement) GetPropertyMaxFlushSize() (value int32, err error) {
	retValue, err := instance.GetProperty("MaxFlushSize")
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

	value = int32(valuetmp)

	return
}

// SetName sets the value of Name for the instance
func (instance *BufferModeElement) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *BufferModeElement) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
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

// SetRegularFlushInterval sets the value of RegularFlushInterval for the instance
func (instance *BufferModeElement) SetPropertyRegularFlushInterval(value string) (err error) {
	return instance.SetProperty("RegularFlushInterval", (value))
}

// GetRegularFlushInterval gets the value of RegularFlushInterval for the instance
func (instance *BufferModeElement) GetPropertyRegularFlushInterval() (value string, err error) {
	retValue, err := instance.GetProperty("RegularFlushInterval")
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

// SetUrgentFlushInterval sets the value of UrgentFlushInterval for the instance
func (instance *BufferModeElement) SetPropertyUrgentFlushInterval(value string) (err error) {
	return instance.SetProperty("UrgentFlushInterval", (value))
}

// GetUrgentFlushInterval gets the value of UrgentFlushInterval for the instance
func (instance *BufferModeElement) GetPropertyUrgentFlushInterval() (value string, err error) {
	retValue, err := instance.GetProperty("UrgentFlushInterval")
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

// SetUrgentFlushThreshold sets the value of UrgentFlushThreshold for the instance
func (instance *BufferModeElement) SetPropertyUrgentFlushThreshold(value int32) (err error) {
	return instance.SetProperty("UrgentFlushThreshold", (value))
}

// GetUrgentFlushThreshold gets the value of UrgentFlushThreshold for the instance
func (instance *BufferModeElement) GetPropertyUrgentFlushThreshold() (value int32, err error) {
	retValue, err := instance.GetProperty("UrgentFlushThreshold")
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

	value = int32(valuetmp)

	return
}
