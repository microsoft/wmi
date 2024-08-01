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

// MSFT_CliTranslateTableEntry struct
type MSFT_CliTranslateTableEntry struct {
	*cim.WmiInstance

	//
	FromValue string

	//
	ToValue string
}

func NewMSFT_CliTranslateTableEntryEx1(instance *cim.WmiInstance) (newInstance *MSFT_CliTranslateTableEntry, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_CliTranslateTableEntry{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_CliTranslateTableEntryEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_CliTranslateTableEntry, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_CliTranslateTableEntry{
		WmiInstance: tmp,
	}
	return
}

// SetFromValue sets the value of FromValue for the instance
func (instance *MSFT_CliTranslateTableEntry) SetPropertyFromValue(value string) (err error) {
	return instance.SetProperty("FromValue", (value))
}

// GetFromValue gets the value of FromValue for the instance
func (instance *MSFT_CliTranslateTableEntry) GetPropertyFromValue() (value string, err error) {
	retValue, err := instance.GetProperty("FromValue")
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

// SetToValue sets the value of ToValue for the instance
func (instance *MSFT_CliTranslateTableEntry) SetPropertyToValue(value string) (err error) {
	return instance.SetProperty("ToValue", (value))
}

// GetToValue gets the value of ToValue for the instance
func (instance *MSFT_CliTranslateTableEntry) GetPropertyToValue() (value string, err error) {
	retValue, err := instance.GetProperty("ToValue")
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
