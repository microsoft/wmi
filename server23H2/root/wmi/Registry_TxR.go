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

// Registry_TxR struct
type Registry_TxR struct {
	*Registry

	//
	Hive string

	//
	OperationTime uint64

	//
	Status uint32

	//
	TxrGUID interface{}

	//
	UowCount uint32
}

func NewRegistry_TxREx1(instance *cim.WmiInstance) (newInstance *Registry_TxR, err error) {
	tmp, err := NewRegistryEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Registry_TxR{
		Registry: tmp,
	}
	return
}

func NewRegistry_TxREx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Registry_TxR, err error) {
	tmp, err := NewRegistryEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Registry_TxR{
		Registry: tmp,
	}
	return
}

// SetHive sets the value of Hive for the instance
func (instance *Registry_TxR) SetPropertyHive(value string) (err error) {
	return instance.SetProperty("Hive", (value))
}

// GetHive gets the value of Hive for the instance
func (instance *Registry_TxR) GetPropertyHive() (value string, err error) {
	retValue, err := instance.GetProperty("Hive")
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

// SetOperationTime sets the value of OperationTime for the instance
func (instance *Registry_TxR) SetPropertyOperationTime(value uint64) (err error) {
	return instance.SetProperty("OperationTime", (value))
}

// GetOperationTime gets the value of OperationTime for the instance
func (instance *Registry_TxR) GetPropertyOperationTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("OperationTime")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetStatus sets the value of Status for the instance
func (instance *Registry_TxR) SetPropertyStatus(value uint32) (err error) {
	return instance.SetProperty("Status", (value))
}

// GetStatus gets the value of Status for the instance
func (instance *Registry_TxR) GetPropertyStatus() (value uint32, err error) {
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

// SetTxrGUID sets the value of TxrGUID for the instance
func (instance *Registry_TxR) SetPropertyTxrGUID(value interface{}) (err error) {
	return instance.SetProperty("TxrGUID", (value))
}

// GetTxrGUID gets the value of TxrGUID for the instance
func (instance *Registry_TxR) GetPropertyTxrGUID() (value interface{}, err error) {
	retValue, err := instance.GetProperty("TxrGUID")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(interface{})
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " interface{} is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = interface{}(valuetmp)

	return
}

// SetUowCount sets the value of UowCount for the instance
func (instance *Registry_TxR) SetPropertyUowCount(value uint32) (err error) {
	return instance.SetProperty("UowCount", (value))
}

// GetUowCount gets the value of UowCount for the instance
func (instance *Registry_TxR) GetPropertyUowCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("UowCount")
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
