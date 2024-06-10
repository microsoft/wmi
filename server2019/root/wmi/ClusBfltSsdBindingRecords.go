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
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ClusBfltSsdBindingRecords struct
type ClusBfltSsdBindingRecords struct {
	*cim.WmiInstance

	// Cache Store Binding Records.
	BindingRecords []ClusBfltHddBindingRecord

	// Number of Binding Records.
	NumberOfBindingRecords uint32
}

func NewClusBfltSsdBindingRecordsEx1(instance *cim.WmiInstance) (newInstance *ClusBfltSsdBindingRecords, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &ClusBfltSsdBindingRecords{
		WmiInstance: tmp,
	}
	return
}

func NewClusBfltSsdBindingRecordsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ClusBfltSsdBindingRecords, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ClusBfltSsdBindingRecords{
		WmiInstance: tmp,
	}
	return
}

// SetBindingRecords sets the value of BindingRecords for the instance
func (instance *ClusBfltSsdBindingRecords) SetPropertyBindingRecords(value []ClusBfltHddBindingRecord) (err error) {
	return instance.SetProperty("BindingRecords", (value))
}

// GetBindingRecords gets the value of BindingRecords for the instance
func (instance *ClusBfltSsdBindingRecords) GetPropertyBindingRecords() (value []ClusBfltHddBindingRecord, err error) {
	retValue, err := instance.GetProperty("BindingRecords")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(ClusBfltHddBindingRecord)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " ClusBfltHddBindingRecord is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, ClusBfltHddBindingRecord(valuetmp))
	}

	return
}

// SetNumberOfBindingRecords sets the value of NumberOfBindingRecords for the instance
func (instance *ClusBfltSsdBindingRecords) SetPropertyNumberOfBindingRecords(value uint32) (err error) {
	return instance.SetProperty("NumberOfBindingRecords", (value))
}

// GetNumberOfBindingRecords gets the value of NumberOfBindingRecords for the instance
func (instance *ClusBfltSsdBindingRecords) GetPropertyNumberOfBindingRecords() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfBindingRecords")
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
