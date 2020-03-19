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

// MSFT_NetIPHttpsState struct
type MSFT_NetIPHttpsState struct {
	*CIM_ElementSettingData

	//
	InterfaceStatus string

	//
	LastErrorCode uint32
}

func NewMSFT_NetIPHttpsStateEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetIPHttpsState, err error) {
	tmp, err := NewCIM_ElementSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetIPHttpsState{
		CIM_ElementSettingData: tmp,
	}
	return
}

func NewMSFT_NetIPHttpsStateEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetIPHttpsState, err error) {
	tmp, err := NewCIM_ElementSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetIPHttpsState{
		CIM_ElementSettingData: tmp,
	}
	return
}

// SetInterfaceStatus sets the value of InterfaceStatus for the instance
func (instance *MSFT_NetIPHttpsState) SetPropertyInterfaceStatus(value string) (err error) {
	return instance.SetProperty("InterfaceStatus", value)
}

// GetInterfaceStatus gets the value of InterfaceStatus for the instance
func (instance *MSFT_NetIPHttpsState) GetPropertyInterfaceStatus() (value string, err error) {
	retValue, err := instance.GetProperty("InterfaceStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLastErrorCode sets the value of LastErrorCode for the instance
func (instance *MSFT_NetIPHttpsState) SetPropertyLastErrorCode(value uint32) (err error) {
	return instance.SetProperty("LastErrorCode", value)
}

// GetLastErrorCode gets the value of LastErrorCode for the instance
func (instance *MSFT_NetIPHttpsState) GetPropertyLastErrorCode() (value uint32, err error) {
	retValue, err := instance.GetProperty("LastErrorCode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
