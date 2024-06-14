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

// Registry_V1_TypeGroup1 struct
type Registry_V1_TypeGroup1 struct {
	*Registry_V1

	//
	ElapsedTime int64

	//
	Index uint32

	//
	KeyHandle uint32

	//
	KeyName string

	//
	Status uint32
}

func NewRegistry_V1_TypeGroup1Ex1(instance *cim.WmiInstance) (newInstance *Registry_V1_TypeGroup1, err error) {
	tmp, err := NewRegistry_V1Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &Registry_V1_TypeGroup1{
		Registry_V1: tmp,
	}
	return
}

func NewRegistry_V1_TypeGroup1Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Registry_V1_TypeGroup1, err error) {
	tmp, err := NewRegistry_V1Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Registry_V1_TypeGroup1{
		Registry_V1: tmp,
	}
	return
}

// SetElapsedTime sets the value of ElapsedTime for the instance
func (instance *Registry_V1_TypeGroup1) SetPropertyElapsedTime(value int64) (err error) {
	return instance.SetProperty("ElapsedTime", (value))
}

// GetElapsedTime gets the value of ElapsedTime for the instance
func (instance *Registry_V1_TypeGroup1) GetPropertyElapsedTime() (value int64, err error) {
	retValue, err := instance.GetProperty("ElapsedTime")
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

// SetIndex sets the value of Index for the instance
func (instance *Registry_V1_TypeGroup1) SetPropertyIndex(value uint32) (err error) {
	return instance.SetProperty("Index", (value))
}

// GetIndex gets the value of Index for the instance
func (instance *Registry_V1_TypeGroup1) GetPropertyIndex() (value uint32, err error) {
	retValue, err := instance.GetProperty("Index")
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

// SetKeyHandle sets the value of KeyHandle for the instance
func (instance *Registry_V1_TypeGroup1) SetPropertyKeyHandle(value uint32) (err error) {
	return instance.SetProperty("KeyHandle", (value))
}

// GetKeyHandle gets the value of KeyHandle for the instance
func (instance *Registry_V1_TypeGroup1) GetPropertyKeyHandle() (value uint32, err error) {
	retValue, err := instance.GetProperty("KeyHandle")
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

// SetKeyName sets the value of KeyName for the instance
func (instance *Registry_V1_TypeGroup1) SetPropertyKeyName(value string) (err error) {
	return instance.SetProperty("KeyName", (value))
}

// GetKeyName gets the value of KeyName for the instance
func (instance *Registry_V1_TypeGroup1) GetPropertyKeyName() (value string, err error) {
	retValue, err := instance.GetProperty("KeyName")
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

// SetStatus sets the value of Status for the instance
func (instance *Registry_V1_TypeGroup1) SetPropertyStatus(value uint32) (err error) {
	return instance.SetProperty("Status", (value))
}

// GetStatus gets the value of Status for the instance
func (instance *Registry_V1_TypeGroup1) GetPropertyStatus() (value uint32, err error) {
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
