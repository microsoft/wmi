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

// AcpiPss struct
type AcpiPss struct {
	*cim.WmiInstance

	//
	Count uint32

	//
	State []AcpiPssState
}

func NewAcpiPssEx1(instance *cim.WmiInstance) (newInstance *AcpiPss, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &AcpiPss{
		WmiInstance: tmp,
	}
	return
}

func NewAcpiPssEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *AcpiPss, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &AcpiPss{
		WmiInstance: tmp,
	}
	return
}

// SetCount sets the value of Count for the instance
func (instance *AcpiPss) SetPropertyCount(value uint32) (err error) {
	return instance.SetProperty("Count", (value))
}

// GetCount gets the value of Count for the instance
func (instance *AcpiPss) GetPropertyCount() (value uint32, err error) {
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

// SetState sets the value of State for the instance
func (instance *AcpiPss) SetPropertyState(value []AcpiPssState) (err error) {
	return instance.SetProperty("State", (value))
}

// GetState gets the value of State for the instance
func (instance *AcpiPss) GetPropertyState() (value []AcpiPssState, err error) {
	retValue, err := instance.GetProperty("State")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(AcpiPssState)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " AcpiPssState is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, AcpiPssState(valuetmp))
	}

	return
}
