// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.Microsoft.Windows.AppBackgroundTask
//////////////////////////////////////////////
package appbackgroundtask

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_BackgroundTask struct
type MSFT_BackgroundTask struct {
	*cim.WmiInstance

	//
	EntryPoint []string

	//
	PackageFullName string

	//
	PerfInfo []string

	//
	TaskID []string

	//
	TaskName []string
}

func NewMSFT_BackgroundTaskEx1(instance *cim.WmiInstance) (newInstance *MSFT_BackgroundTask, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_BackgroundTask{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_BackgroundTaskEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_BackgroundTask, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_BackgroundTask{
		WmiInstance: tmp,
	}
	return
}

// SetEntryPoint sets the value of EntryPoint for the instance
func (instance *MSFT_BackgroundTask) SetPropertyEntryPoint(value []string) (err error) {
	return instance.SetProperty("EntryPoint", (value))
}

// GetEntryPoint gets the value of EntryPoint for the instance
func (instance *MSFT_BackgroundTask) GetPropertyEntryPoint() (value []string, err error) {
	retValue, err := instance.GetProperty("EntryPoint")
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

// SetPackageFullName sets the value of PackageFullName for the instance
func (instance *MSFT_BackgroundTask) SetPropertyPackageFullName(value string) (err error) {
	return instance.SetProperty("PackageFullName", (value))
}

// GetPackageFullName gets the value of PackageFullName for the instance
func (instance *MSFT_BackgroundTask) GetPropertyPackageFullName() (value string, err error) {
	retValue, err := instance.GetProperty("PackageFullName")
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

// SetPerfInfo sets the value of PerfInfo for the instance
func (instance *MSFT_BackgroundTask) SetPropertyPerfInfo(value []string) (err error) {
	return instance.SetProperty("PerfInfo", (value))
}

// GetPerfInfo gets the value of PerfInfo for the instance
func (instance *MSFT_BackgroundTask) GetPropertyPerfInfo() (value []string, err error) {
	retValue, err := instance.GetProperty("PerfInfo")
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

// SetTaskID sets the value of TaskID for the instance
func (instance *MSFT_BackgroundTask) SetPropertyTaskID(value []string) (err error) {
	return instance.SetProperty("TaskID", (value))
}

// GetTaskID gets the value of TaskID for the instance
func (instance *MSFT_BackgroundTask) GetPropertyTaskID() (value []string, err error) {
	retValue, err := instance.GetProperty("TaskID")
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

// SetTaskName sets the value of TaskName for the instance
func (instance *MSFT_BackgroundTask) SetPropertyTaskName(value []string) (err error) {
	return instance.SetProperty("TaskName", (value))
}

// GetTaskName gets the value of TaskName for the instance
func (instance *MSFT_BackgroundTask) GetPropertyTaskName() (value []string, err error) {
	retValue, err := instance.GetProperty("TaskName")
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
