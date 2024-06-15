// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.Microsoft.Windows.TaskScheduler
//////////////////////////////////////////////
package taskscheduler

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_TaskExecAction struct
type MSFT_TaskExecAction struct {
	*MSFT_TaskAction

	//
	Arguments string

	//
	Execute string

	//
	WorkingDirectory string
}

func NewMSFT_TaskExecActionEx1(instance *cim.WmiInstance) (newInstance *MSFT_TaskExecAction, err error) {
	tmp, err := NewMSFT_TaskActionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_TaskExecAction{
		MSFT_TaskAction: tmp,
	}
	return
}

func NewMSFT_TaskExecActionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_TaskExecAction, err error) {
	tmp, err := NewMSFT_TaskActionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_TaskExecAction{
		MSFT_TaskAction: tmp,
	}
	return
}

// SetArguments sets the value of Arguments for the instance
func (instance *MSFT_TaskExecAction) SetPropertyArguments(value string) (err error) {
	return instance.SetProperty("Arguments", (value))
}

// GetArguments gets the value of Arguments for the instance
func (instance *MSFT_TaskExecAction) GetPropertyArguments() (value string, err error) {
	retValue, err := instance.GetProperty("Arguments")
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

// SetExecute sets the value of Execute for the instance
func (instance *MSFT_TaskExecAction) SetPropertyExecute(value string) (err error) {
	return instance.SetProperty("Execute", (value))
}

// GetExecute gets the value of Execute for the instance
func (instance *MSFT_TaskExecAction) GetPropertyExecute() (value string, err error) {
	retValue, err := instance.GetProperty("Execute")
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

// SetWorkingDirectory sets the value of WorkingDirectory for the instance
func (instance *MSFT_TaskExecAction) SetPropertyWorkingDirectory(value string) (err error) {
	return instance.SetProperty("WorkingDirectory", (value))
}

// GetWorkingDirectory gets the value of WorkingDirectory for the instance
func (instance *MSFT_TaskExecAction) GetPropertyWorkingDirectory() (value string, err error) {
	retValue, err := instance.GetProperty("WorkingDirectory")
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
