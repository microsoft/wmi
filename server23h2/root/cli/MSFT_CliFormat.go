// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.Cli
//////////////////////////////////////////////
package cli

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_CliFormat struct
type MSFT_CliFormat struct {
	*cim.WmiInstance

	//
	Format string

	//
	Name string

	//
	Properties []MSFT_CliProperty
}

func NewMSFT_CliFormatEx1(instance *cim.WmiInstance) (newInstance *MSFT_CliFormat, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_CliFormat{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_CliFormatEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_CliFormat, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_CliFormat{
		WmiInstance: tmp,
	}
	return
}

// SetFormat sets the value of Format for the instance
func (instance *MSFT_CliFormat) SetPropertyFormat(value string) (err error) {
	return instance.SetProperty("Format", (value))
}

// GetFormat gets the value of Format for the instance
func (instance *MSFT_CliFormat) GetPropertyFormat() (value string, err error) {
	retValue, err := instance.GetProperty("Format")
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
func (instance *MSFT_CliFormat) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *MSFT_CliFormat) GetPropertyName() (value string, err error) {
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

// SetProperties sets the value of Properties for the instance
func (instance *MSFT_CliFormat) SetPropertyProperties(value []MSFT_CliProperty) (err error) {
	return instance.SetProperty("Properties", (value))
}

// GetProperties gets the value of Properties for the instance
func (instance *MSFT_CliFormat) GetPropertyProperties() (value []MSFT_CliProperty, err error) {
	retValue, err := instance.GetProperty("Properties")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(MSFT_CliProperty)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " MSFT_CliProperty is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, MSFT_CliProperty(valuetmp))
	}

	return
}
