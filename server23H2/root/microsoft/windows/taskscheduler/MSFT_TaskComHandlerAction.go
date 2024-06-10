// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.TaskScheduler
//
// ////////////////////////////////////////////
package taskscheduler

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_TaskComHandlerAction struct
type MSFT_TaskComHandlerAction struct {
	*MSFT_TaskAction

	//
	ClassId string

	//
	Data string
}

func NewMSFT_TaskComHandlerActionEx1(instance *cim.WmiInstance) (newInstance *MSFT_TaskComHandlerAction, err error) {
	tmp, err := NewMSFT_TaskActionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_TaskComHandlerAction{
		MSFT_TaskAction: tmp,
	}
	return
}

func NewMSFT_TaskComHandlerActionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_TaskComHandlerAction, err error) {
	tmp, err := NewMSFT_TaskActionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_TaskComHandlerAction{
		MSFT_TaskAction: tmp,
	}
	return
}

// SetClassId sets the value of ClassId for the instance
func (instance *MSFT_TaskComHandlerAction) SetPropertyClassId(value string) (err error) {
	return instance.SetProperty("ClassId", (value))
}

// GetClassId gets the value of ClassId for the instance
func (instance *MSFT_TaskComHandlerAction) GetPropertyClassId() (value string, err error) {
	retValue, err := instance.GetProperty("ClassId")
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

// SetData sets the value of Data for the instance
func (instance *MSFT_TaskComHandlerAction) SetPropertyData(value string) (err error) {
	return instance.SetProperty("Data", (value))
}

// GetData gets the value of Data for the instance
func (instance *MSFT_TaskComHandlerAction) GetPropertyData() (value string, err error) {
	retValue, err := instance.GetProperty("Data")
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
