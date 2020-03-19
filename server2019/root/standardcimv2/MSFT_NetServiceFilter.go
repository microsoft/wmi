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

// MSFT_NetServiceFilter struct
type MSFT_NetServiceFilter struct {
	*CIM_FilterEntryBase

	//
	ServiceName string
}

func NewMSFT_NetServiceFilterEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetServiceFilter, err error) {
	tmp, err := NewCIM_FilterEntryBaseEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetServiceFilter{
		CIM_FilterEntryBase: tmp,
	}
	return
}

func NewMSFT_NetServiceFilterEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetServiceFilter, err error) {
	tmp, err := NewCIM_FilterEntryBaseEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetServiceFilter{
		CIM_FilterEntryBase: tmp,
	}
	return
}

// SetServiceName sets the value of ServiceName for the instance
func (instance *MSFT_NetServiceFilter) SetPropertyServiceName(value string) (err error) {
	return instance.SetProperty("ServiceName", value)
}

// GetServiceName gets the value of ServiceName for the instance
func (instance *MSFT_NetServiceFilter) GetPropertyServiceName() (value string, err error) {
	retValue, err := instance.GetProperty("ServiceName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
