// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.Microsoft.Windows.ManagementTools
//////////////////////////////////////////////
package managementtools

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_MTRegistryDword struct
type MSFT_MTRegistryDword struct {
	*MSFT_MTRegistryValue

	//
	Data uint32
}

func NewMSFT_MTRegistryDwordEx1(instance *cim.WmiInstance) (newInstance *MSFT_MTRegistryDword, err error) {
	tmp, err := NewMSFT_MTRegistryValueEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_MTRegistryDword{
		MSFT_MTRegistryValue: tmp,
	}
	return
}

func NewMSFT_MTRegistryDwordEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_MTRegistryDword, err error) {
	tmp, err := NewMSFT_MTRegistryValueEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_MTRegistryDword{
		MSFT_MTRegistryValue: tmp,
	}
	return
}

// SetData sets the value of Data for the instance
func (instance *MSFT_MTRegistryDword) SetPropertyData(value uint32) (err error) {
	return instance.SetProperty("Data", (value))
}

// GetData gets the value of Data for the instance
func (instance *MSFT_MTRegistryDword) GetPropertyData() (value uint32, err error) {
	retValue, err := instance.GetProperty("Data")
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
