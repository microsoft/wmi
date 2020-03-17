// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.ServerManager
//////////////////////////////////////////////
package servermanager

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_ServiceToMonitor struct
type MSFT_ServiceToMonitor struct {
	cim.WmiInstance

	//
	DefaultMonitoring bool

	//
	Name string
}

// SetDefaultMonitoring sets the value of DefaultMonitoring for the instance
func (instance *MSFT_ServiceToMonitor) SetPropertyDefaultMonitoring(value bool) (err error) {
	return instance.SetProperty("DefaultMonitoring", value)
}

// GetDefaultMonitoring gets the value of DefaultMonitoring for the instance
func (instance *MSFT_ServiceToMonitor) GetPropertyDefaultMonitoring() (value bool, err error) {
	retValue, err := instance.GetProperty("DefaultMonitoring")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *MSFT_ServiceToMonitor) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *MSFT_ServiceToMonitor) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
