// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.Windows.TaskScheduler
//////////////////////////////////////////////
package taskscheduler

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_TaskSettings2 struct
type MSFT_TaskSettings2 struct {
	*MSFT_TaskSettings

	//
	DisallowStartOnRemoteAppSession bool

	//
	UseUnifiedSchedulingEngine bool
}

func NewMSFT_TaskSettings2Ex1(instance *cim.WmiInstance) (newInstance *MSFT_TaskSettings2, err error) {
	tmp, err := NewMSFT_TaskSettingsEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_TaskSettings2{
		MSFT_TaskSettings: tmp,
	}
	return
}

func NewMSFT_TaskSettings2Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_TaskSettings2, err error) {
	tmp, err := NewMSFT_TaskSettingsEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_TaskSettings2{
		MSFT_TaskSettings: tmp,
	}
	return
}

// SetDisallowStartOnRemoteAppSession sets the value of DisallowStartOnRemoteAppSession for the instance
func (instance *MSFT_TaskSettings2) SetPropertyDisallowStartOnRemoteAppSession(value bool) (err error) {
	return instance.SetProperty("DisallowStartOnRemoteAppSession", (value))
}

// GetDisallowStartOnRemoteAppSession gets the value of DisallowStartOnRemoteAppSession for the instance
func (instance *MSFT_TaskSettings2) GetPropertyDisallowStartOnRemoteAppSession() (value bool, err error) {
	retValue, err := instance.GetProperty("DisallowStartOnRemoteAppSession")
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

// SetUseUnifiedSchedulingEngine sets the value of UseUnifiedSchedulingEngine for the instance
func (instance *MSFT_TaskSettings2) SetPropertyUseUnifiedSchedulingEngine(value bool) (err error) {
	return instance.SetProperty("UseUnifiedSchedulingEngine", (value))
}

// GetUseUnifiedSchedulingEngine gets the value of UseUnifiedSchedulingEngine for the instance
func (instance *MSFT_TaskSettings2) GetPropertyUseUnifiedSchedulingEngine() (value bool, err error) {
	retValue, err := instance.GetProperty("UseUnifiedSchedulingEngine")
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
