// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.Microsoft.Windows.ServerManager
//////////////////////////////////////////////
package servermanager

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_ServerBpaResult struct
type MSFT_ServerBpaResult struct {
	*cim.WmiInstance

	//
	BpaXPath string

	//
	Values []string
}

func NewMSFT_ServerBpaResultEx1(instance *cim.WmiInstance) (newInstance *MSFT_ServerBpaResult, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_ServerBpaResult{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_ServerBpaResultEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_ServerBpaResult, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_ServerBpaResult{
		WmiInstance: tmp,
	}
	return
}

// SetBpaXPath sets the value of BpaXPath for the instance
func (instance *MSFT_ServerBpaResult) SetPropertyBpaXPath(value string) (err error) {
	return instance.SetProperty("BpaXPath", (value))
}

// GetBpaXPath gets the value of BpaXPath for the instance
func (instance *MSFT_ServerBpaResult) GetPropertyBpaXPath() (value string, err error) {
	retValue, err := instance.GetProperty("BpaXPath")
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

// SetValues sets the value of Values for the instance
func (instance *MSFT_ServerBpaResult) SetPropertyValues(value []string) (err error) {
	return instance.SetProperty("Values", (value))
}

// GetValues gets the value of Values for the instance
func (instance *MSFT_ServerBpaResult) GetPropertyValues() (value []string, err error) {
	retValue, err := instance.GetProperty("Values")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}
