// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.Microsoft.Windows.DesiredStateConfiguration
//////////////////////////////////////////////
package desiredstateconfiguration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_DSCConfigurationOutputWriteArray struct
type MSFT_DSCConfigurationOutputWriteArray struct {
	*MSFT_DSCConfigurationOutput

	//
	Array []interface{}

	//
	ParameterName string
}

func NewMSFT_DSCConfigurationOutputWriteArrayEx1(instance *cim.WmiInstance) (newInstance *MSFT_DSCConfigurationOutputWriteArray, err error) {
	tmp, err := NewMSFT_DSCConfigurationOutputEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_DSCConfigurationOutputWriteArray{
		MSFT_DSCConfigurationOutput: tmp,
	}
	return
}

func NewMSFT_DSCConfigurationOutputWriteArrayEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_DSCConfigurationOutputWriteArray, err error) {
	tmp, err := NewMSFT_DSCConfigurationOutputEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_DSCConfigurationOutputWriteArray{
		MSFT_DSCConfigurationOutput: tmp,
	}
	return
}

// SetArray sets the value of Array for the instance
func (instance *MSFT_DSCConfigurationOutputWriteArray) SetPropertyArray(value []interface{}) (err error) {
	return instance.SetProperty("Array", (value))
}

// GetArray gets the value of Array for the instance
func (instance *MSFT_DSCConfigurationOutputWriteArray) GetPropertyArray() (value []interface{}, err error) {
	retValue, err := instance.GetProperty("Array")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(interface{})
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " interface{} is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, interface{}(valuetmp))
	}

	return
}

// SetParameterName sets the value of ParameterName for the instance
func (instance *MSFT_DSCConfigurationOutputWriteArray) SetPropertyParameterName(value string) (err error) {
	return instance.SetProperty("ParameterName", (value))
}

// GetParameterName gets the value of ParameterName for the instance
func (instance *MSFT_DSCConfigurationOutputWriteArray) GetPropertyParameterName() (value string, err error) {
	retValue, err := instance.GetProperty("ParameterName")
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
