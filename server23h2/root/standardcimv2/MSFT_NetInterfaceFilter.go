// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NetInterfaceFilter struct
type MSFT_NetInterfaceFilter struct {
	*CIM_FilterEntryBase

	//
	InterfaceAlias []string
}

func NewMSFT_NetInterfaceFilterEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetInterfaceFilter, err error) {
	tmp, err := NewCIM_FilterEntryBaseEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetInterfaceFilter{
		CIM_FilterEntryBase: tmp,
	}
	return
}

func NewMSFT_NetInterfaceFilterEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetInterfaceFilter, err error) {
	tmp, err := NewCIM_FilterEntryBaseEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetInterfaceFilter{
		CIM_FilterEntryBase: tmp,
	}
	return
}

// SetInterfaceAlias sets the value of InterfaceAlias for the instance
func (instance *MSFT_NetInterfaceFilter) SetPropertyInterfaceAlias(value []string) (err error) {
	return instance.SetProperty("InterfaceAlias", (value))
}

// GetInterfaceAlias gets the value of InterfaceAlias for the instance
func (instance *MSFT_NetInterfaceFilter) GetPropertyInterfaceAlias() (value []string, err error) {
	retValue, err := instance.GetProperty("InterfaceAlias")
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
