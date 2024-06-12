// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.CIMV2.power
//////////////////////////////////////////////
package power

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_PowerPlan struct
type Win32_PowerPlan struct {
	*CIM_SettingData

	//
	IsActive bool
}

func NewWin32_PowerPlanEx1(instance *cim.WmiInstance) (newInstance *Win32_PowerPlan, err error) {
	tmp, err := NewCIM_SettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PowerPlan{
		CIM_SettingData: tmp,
	}
	return
}

func NewWin32_PowerPlanEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PowerPlan, err error) {
	tmp, err := NewCIM_SettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PowerPlan{
		CIM_SettingData: tmp,
	}
	return
}

// SetIsActive sets the value of IsActive for the instance
func (instance *Win32_PowerPlan) SetPropertyIsActive(value bool) (err error) {
	return instance.SetProperty("IsActive", (value))
}

// GetIsActive gets the value of IsActive for the instance
func (instance *Win32_PowerPlan) GetPropertyIsActive() (value bool, err error) {
	retValue, err := instance.GetProperty("IsActive")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

//

// <param name="ReturnValue" type="bool "></param>
func (instance *Win32_PowerPlan) Activate() (result bool, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Activate")
	if err != nil {
		return
	}
	result = (retVal > 0)
	return

}
