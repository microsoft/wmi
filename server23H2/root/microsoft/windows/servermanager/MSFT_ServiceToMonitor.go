// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.ServerManager
//
// ////////////////////////////////////////////
package servermanager

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_ServiceToMonitor struct
type MSFT_ServiceToMonitor struct {
	*cim.WmiInstance

	//
	DefaultMonitoring bool

	//
	Name string
}

func NewMSFT_ServiceToMonitorEx1(instance *cim.WmiInstance) (newInstance *MSFT_ServiceToMonitor, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_ServiceToMonitor{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_ServiceToMonitorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_ServiceToMonitor, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_ServiceToMonitor{
		WmiInstance: tmp,
	}
	return
}

// SetDefaultMonitoring sets the value of DefaultMonitoring for the instance
func (instance *MSFT_ServiceToMonitor) SetPropertyDefaultMonitoring(value bool) (err error) {
	return instance.SetProperty("DefaultMonitoring", (value))
}

// GetDefaultMonitoring gets the value of DefaultMonitoring for the instance
func (instance *MSFT_ServiceToMonitor) GetPropertyDefaultMonitoring() (value bool, err error) {
	retValue, err := instance.GetProperty("DefaultMonitoring")
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

// SetName sets the value of Name for the instance
func (instance *MSFT_ServiceToMonitor) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *MSFT_ServiceToMonitor) GetPropertyName() (value string, err error) {
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
