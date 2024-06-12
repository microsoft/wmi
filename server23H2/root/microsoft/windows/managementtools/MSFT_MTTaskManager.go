// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.Microsoft.Windows.ManagementTools
//////////////////////////////////////////////
package managementtools

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_MTTaskManager struct
type MSFT_MTTaskManager struct {
	*CIM_ManagedElement

	//
	CurrentIndex uint16

	//
	IntervalSeconds uint16

	//
	Name string
}

func NewMSFT_MTTaskManagerEx1(instance *cim.WmiInstance) (newInstance *MSFT_MTTaskManager, err error) {
	tmp, err := NewCIM_ManagedElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_MTTaskManager{
		CIM_ManagedElement: tmp,
	}
	return
}

func NewMSFT_MTTaskManagerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_MTTaskManager, err error) {
	tmp, err := NewCIM_ManagedElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_MTTaskManager{
		CIM_ManagedElement: tmp,
	}
	return
}

// SetCurrentIndex sets the value of CurrentIndex for the instance
func (instance *MSFT_MTTaskManager) SetPropertyCurrentIndex(value uint16) (err error) {
	return instance.SetProperty("CurrentIndex", (value))
}

// GetCurrentIndex gets the value of CurrentIndex for the instance
func (instance *MSFT_MTTaskManager) GetPropertyCurrentIndex() (value uint16, err error) {
	retValue, err := instance.GetProperty("CurrentIndex")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetIntervalSeconds sets the value of IntervalSeconds for the instance
func (instance *MSFT_MTTaskManager) SetPropertyIntervalSeconds(value uint16) (err error) {
	return instance.SetProperty("IntervalSeconds", (value))
}

// GetIntervalSeconds gets the value of IntervalSeconds for the instance
func (instance *MSFT_MTTaskManager) GetPropertyIntervalSeconds() (value uint16, err error) {
	retValue, err := instance.GetProperty("IntervalSeconds")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetName sets the value of Name for the instance
func (instance *MSFT_MTTaskManager) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *MSFT_MTTaskManager) GetPropertyName() (value string, err error) {
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

//

// <param name="Seconds" type="uint16 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_MTTaskManager) SetInterval( /* IN */ Seconds uint16) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetInterval", Seconds)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_MTTaskManager) ForceRefresh() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("ForceRefresh")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
