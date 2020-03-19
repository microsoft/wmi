// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.ManagementTools
//////////////////////////////////////////////
package managementtools

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_MTRegistryString struct
type MSFT_MTRegistryString struct {
	*MSFT_MTRegistryValue

	//
	Data string
}

func NewMSFT_MTRegistryStringEx1(instance *cim.WmiInstance) (newInstance *MSFT_MTRegistryString, err error) {
	tmp, err := NewMSFT_MTRegistryValueEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_MTRegistryString{
		MSFT_MTRegistryValue: tmp,
	}
	return
}

func NewMSFT_MTRegistryStringEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_MTRegistryString, err error) {
	tmp, err := NewMSFT_MTRegistryValueEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_MTRegistryString{
		MSFT_MTRegistryValue: tmp,
	}
	return
}

// SetData sets the value of Data for the instance
func (instance *MSFT_MTRegistryString) SetPropertyData(value string) (err error) {
	return instance.SetProperty("Data", value)
}

// GetData gets the value of Data for the instance
func (instance *MSFT_MTRegistryString) GetPropertyData() (value string, err error) {
	retValue, err := instance.GetProperty("Data")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
