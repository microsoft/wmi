// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.virtualization.v2
//
// ////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_KvpExchangeDataItem struct
type Msvm_KvpExchangeDataItem struct {
	*CIM_ManagedElement

	//
	Data string

	//
	Name string

	//
	Source uint16
}

func NewMsvm_KvpExchangeDataItemEx1(instance *cim.WmiInstance) (newInstance *Msvm_KvpExchangeDataItem, err error) {
	tmp, err := NewCIM_ManagedElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_KvpExchangeDataItem{
		CIM_ManagedElement: tmp,
	}
	return
}

func NewMsvm_KvpExchangeDataItemEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_KvpExchangeDataItem, err error) {
	tmp, err := NewCIM_ManagedElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_KvpExchangeDataItem{
		CIM_ManagedElement: tmp,
	}
	return
}

// SetData sets the value of Data for the instance
func (instance *Msvm_KvpExchangeDataItem) SetPropertyData(value string) (err error) {
	return instance.SetProperty("Data", (value))
}

// GetData gets the value of Data for the instance
func (instance *Msvm_KvpExchangeDataItem) GetPropertyData() (value string, err error) {
	retValue, err := instance.GetProperty("Data")
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

// SetName sets the value of Name for the instance
func (instance *Msvm_KvpExchangeDataItem) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *Msvm_KvpExchangeDataItem) GetPropertyName() (value string, err error) {
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

// SetSource sets the value of Source for the instance
func (instance *Msvm_KvpExchangeDataItem) SetPropertySource(value uint16) (err error) {
	return instance.SetProperty("Source", (value))
}

// GetSource gets the value of Source for the instance
func (instance *Msvm_KvpExchangeDataItem) GetPropertySource() (value uint16, err error) {
	retValue, err := instance.GetProperty("Source")
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
