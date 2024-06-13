// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
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

// MSFT_CliTranslateTable struct
type MSFT_CliTranslateTable struct {
	*cim.WmiInstance

	//
	Name string

	//
	Tbl []MSFT_CliTranslateTableEntry
}

func NewMSFT_CliTranslateTableEx1(instance *cim.WmiInstance) (newInstance *MSFT_CliTranslateTable, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_CliTranslateTable{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_CliTranslateTableEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_CliTranslateTable, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_CliTranslateTable{
		WmiInstance: tmp,
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *MSFT_CliTranslateTable) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *MSFT_CliTranslateTable) GetPropertyName() (value string, err error) {
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

// SetTbl sets the value of Tbl for the instance
func (instance *MSFT_CliTranslateTable) SetPropertyTbl(value []MSFT_CliTranslateTableEntry) (err error) {
	return instance.SetProperty("Tbl", (value))
}

// GetTbl gets the value of Tbl for the instance
func (instance *MSFT_CliTranslateTable) GetPropertyTbl() (value []MSFT_CliTranslateTableEntry, err error) {
	retValue, err := instance.GetProperty("Tbl")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(MSFT_CliTranslateTableEntry)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " MSFT_CliTranslateTableEntry is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, MSFT_CliTranslateTableEntry(valuetmp))
	}

	return
}
