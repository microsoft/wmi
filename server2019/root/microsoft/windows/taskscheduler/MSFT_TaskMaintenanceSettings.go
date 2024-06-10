// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.Microsoft.Windows.TaskScheduler
//////////////////////////////////////////////
package taskscheduler

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_TaskMaintenanceSettings struct
type MSFT_TaskMaintenanceSettings struct {
	*cim.WmiInstance

	//
	Deadline string

	//
	Exclusive bool

	//
	Period string
}

func NewMSFT_TaskMaintenanceSettingsEx1(instance *cim.WmiInstance) (newInstance *MSFT_TaskMaintenanceSettings, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_TaskMaintenanceSettings{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_TaskMaintenanceSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_TaskMaintenanceSettings, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_TaskMaintenanceSettings{
		WmiInstance: tmp,
	}
	return
}

// SetDeadline sets the value of Deadline for the instance
func (instance *MSFT_TaskMaintenanceSettings) SetPropertyDeadline(value string) (err error) {
	return instance.SetProperty("Deadline", (value))
}

// GetDeadline gets the value of Deadline for the instance
func (instance *MSFT_TaskMaintenanceSettings) GetPropertyDeadline() (value string, err error) {
	retValue, err := instance.GetProperty("Deadline")
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

// SetExclusive sets the value of Exclusive for the instance
func (instance *MSFT_TaskMaintenanceSettings) SetPropertyExclusive(value bool) (err error) {
	return instance.SetProperty("Exclusive", (value))
}

// GetExclusive gets the value of Exclusive for the instance
func (instance *MSFT_TaskMaintenanceSettings) GetPropertyExclusive() (value bool, err error) {
	retValue, err := instance.GetProperty("Exclusive")
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

// SetPeriod sets the value of Period for the instance
func (instance *MSFT_TaskMaintenanceSettings) SetPropertyPeriod(value string) (err error) {
	return instance.SetProperty("Period", (value))
}

// GetPeriod gets the value of Period for the instance
func (instance *MSFT_TaskMaintenanceSettings) GetPropertyPeriod() (value string, err error) {
	retValue, err := instance.GetProperty("Period")
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
