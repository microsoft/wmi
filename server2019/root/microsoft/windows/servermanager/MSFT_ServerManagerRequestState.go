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

// MSFT_ServerManagerRequestState struct
type MSFT_ServerManagerRequestState struct {
	cim.WmiInstance

	//
	Error MSFT_ServerManagerDeploymentError

	//
	ProgressTicks uint32

	//
	RequestState uint8

	//
	RestartRequired bool

	//
	TotalTicks uint32

	//
	Warnings []string
}

// SetError sets the value of Error for the instance
func (instance *MSFT_ServerManagerRequestState) SetPropertyError(value MSFT_ServerManagerDeploymentError) (err error) {
	return instance.SetProperty("Error", value)
}

// GetError gets the value of Error for the instance
func (instance *MSFT_ServerManagerRequestState) GetPropertyError() (value MSFT_ServerManagerDeploymentError, err error) {
	retValue, err := instance.GetProperty("Error")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_ServerManagerDeploymentError)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProgressTicks sets the value of ProgressTicks for the instance
func (instance *MSFT_ServerManagerRequestState) SetPropertyProgressTicks(value uint32) (err error) {
	return instance.SetProperty("ProgressTicks", value)
}

// GetProgressTicks gets the value of ProgressTicks for the instance
func (instance *MSFT_ServerManagerRequestState) GetPropertyProgressTicks() (value uint32, err error) {
	retValue, err := instance.GetProperty("ProgressTicks")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRequestState sets the value of RequestState for the instance
func (instance *MSFT_ServerManagerRequestState) SetPropertyRequestState(value uint8) (err error) {
	return instance.SetProperty("RequestState", value)
}

// GetRequestState gets the value of RequestState for the instance
func (instance *MSFT_ServerManagerRequestState) GetPropertyRequestState() (value uint8, err error) {
	retValue, err := instance.GetProperty("RequestState")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRestartRequired sets the value of RestartRequired for the instance
func (instance *MSFT_ServerManagerRequestState) SetPropertyRestartRequired(value bool) (err error) {
	return instance.SetProperty("RestartRequired", value)
}

// GetRestartRequired gets the value of RestartRequired for the instance
func (instance *MSFT_ServerManagerRequestState) GetPropertyRestartRequired() (value bool, err error) {
	retValue, err := instance.GetProperty("RestartRequired")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalTicks sets the value of TotalTicks for the instance
func (instance *MSFT_ServerManagerRequestState) SetPropertyTotalTicks(value uint32) (err error) {
	return instance.SetProperty("TotalTicks", value)
}

// GetTotalTicks gets the value of TotalTicks for the instance
func (instance *MSFT_ServerManagerRequestState) GetPropertyTotalTicks() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalTicks")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWarnings sets the value of Warnings for the instance
func (instance *MSFT_ServerManagerRequestState) SetPropertyWarnings(value []string) (err error) {
	return instance.SetProperty("Warnings", value)
}

// GetWarnings gets the value of Warnings for the instance
func (instance *MSFT_ServerManagerRequestState) GetPropertyWarnings() (value []string, err error) {
	retValue, err := instance.GetProperty("Warnings")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}
