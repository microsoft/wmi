// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NetInterfaceTypeFilter struct
type MSFT_NetInterfaceTypeFilter struct {
	*CIM_FilterEntryBase

	//
	InterfaceType uint32
}

func NewMSFT_NetInterfaceTypeFilterEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetInterfaceTypeFilter, err error) {
	tmp, err := NewCIM_FilterEntryBaseEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetInterfaceTypeFilter{
		CIM_FilterEntryBase: tmp,
	}
	return
}

func NewMSFT_NetInterfaceTypeFilterEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetInterfaceTypeFilter, err error) {
	tmp, err := NewCIM_FilterEntryBaseEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetInterfaceTypeFilter{
		CIM_FilterEntryBase: tmp,
	}
	return
}

// SetInterfaceType sets the value of InterfaceType for the instance
func (instance *MSFT_NetInterfaceTypeFilter) SetPropertyInterfaceType(value uint32) (err error) {
	return instance.SetProperty("InterfaceType", value)
}

// GetInterfaceType gets the value of InterfaceType for the instance
func (instance *MSFT_NetInterfaceTypeFilter) GetPropertyInterfaceType() (value uint32, err error) {
	retValue, err := instance.GetProperty("InterfaceType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
