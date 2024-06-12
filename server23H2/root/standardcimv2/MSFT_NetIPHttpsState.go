// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
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
	return instance.SetProperty("InterfaceStatus", (value))
}

// GetInterfaceStatus gets the value of InterfaceStatus for the instance
func (instance *MSFT_NetIPHttpsState) GetPropertyInterfaceStatus() (value string, err error) {
	retValue, err := instance.GetProperty("InterfaceStatus")
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

// SetLastErrorCode sets the value of LastErrorCode for the instance
func (instance *MSFT_NetIPHttpsState) SetPropertyLastErrorCode(value uint32) (err error) {
	return instance.SetProperty("LastErrorCode", (value))
}

// GetLastErrorCode gets the value of LastErrorCode for the instance
func (instance *MSFT_NetIPHttpsState) GetPropertyLastErrorCode() (value uint32, err error) {
	retValue, err := instance.GetProperty("LastErrorCode")
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
