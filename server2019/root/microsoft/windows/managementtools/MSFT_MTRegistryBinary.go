// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.ManagementTools
//
// ////////////////////////////////////////////
package managementtools

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_MTRegistryBinary struct
type MSFT_MTRegistryBinary struct {
	*MSFT_MTRegistryValue

	//
	Data []uint8
}

func NewMSFT_MTRegistryBinaryEx1(instance *cim.WmiInstance) (newInstance *MSFT_MTRegistryBinary, err error) {
	tmp, err := NewMSFT_MTRegistryValueEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_MTRegistryBinary{
		MSFT_MTRegistryValue: tmp,
	}
	return
}

func NewMSFT_MTRegistryBinaryEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_MTRegistryBinary, err error) {
	tmp, err := NewMSFT_MTRegistryValueEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_MTRegistryBinary{
		MSFT_MTRegistryValue: tmp,
	}
	return
}

// SetData sets the value of Data for the instance
func (instance *MSFT_MTRegistryBinary) SetPropertyData(value []uint8) (err error) {
	return instance.SetProperty("Data", (value))
}

// GetData gets the value of Data for the instance
func (instance *MSFT_MTRegistryBinary) GetPropertyData() (value []uint8, err error) {
	retValue, err := instance.GetProperty("Data")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint8)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint8(valuetmp))
	}

	return
}
