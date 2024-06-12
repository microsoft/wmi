// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Registry_HiveInitialize struct
type Registry_HiveInitialize struct {
	*Registry

	//
	FileName string

	//
	Hive uint32

	//
	OperationType uint32

	//
	PoolTag uint32

	//
	Size uint32
}

func NewRegistry_HiveInitializeEx1(instance *cim.WmiInstance) (newInstance *Registry_HiveInitialize, err error) {
	tmp, err := NewRegistryEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Registry_HiveInitialize{
		Registry: tmp,
	}
	return
}

func NewRegistry_HiveInitializeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Registry_HiveInitialize, err error) {
	tmp, err := NewRegistryEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Registry_HiveInitialize{
		Registry: tmp,
	}
	return
}

// SetFileName sets the value of FileName for the instance
func (instance *Registry_HiveInitialize) SetPropertyFileName(value string) (err error) {
	return instance.SetProperty("FileName", (value))
}

// GetFileName gets the value of FileName for the instance
func (instance *Registry_HiveInitialize) GetPropertyFileName() (value string, err error) {
	retValue, err := instance.GetProperty("FileName")
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

// SetHive sets the value of Hive for the instance
func (instance *Registry_HiveInitialize) SetPropertyHive(value uint32) (err error) {
	return instance.SetProperty("Hive", (value))
}

// GetHive gets the value of Hive for the instance
func (instance *Registry_HiveInitialize) GetPropertyHive() (value uint32, err error) {
	retValue, err := instance.GetProperty("Hive")
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

// SetOperationType sets the value of OperationType for the instance
func (instance *Registry_HiveInitialize) SetPropertyOperationType(value uint32) (err error) {
	return instance.SetProperty("OperationType", (value))
}

// GetOperationType gets the value of OperationType for the instance
func (instance *Registry_HiveInitialize) GetPropertyOperationType() (value uint32, err error) {
	retValue, err := instance.GetProperty("OperationType")
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

// SetPoolTag sets the value of PoolTag for the instance
func (instance *Registry_HiveInitialize) SetPropertyPoolTag(value uint32) (err error) {
	return instance.SetProperty("PoolTag", (value))
}

// GetPoolTag gets the value of PoolTag for the instance
func (instance *Registry_HiveInitialize) GetPropertyPoolTag() (value uint32, err error) {
	retValue, err := instance.GetProperty("PoolTag")
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
func (instance *Registry_HiveInitialize) SetPropertySize(value uint32) (err error) {
	return instance.SetProperty("Size", (value))
}

// GetSize gets the value of Size for the instance
func (instance *Registry_HiveInitialize) GetPropertySize() (value uint32, err error) {
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
