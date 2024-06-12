// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.Microsoft.Windows.ServerManager
//////////////////////////////////////////////
package servermanager

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_ServerManagerRequestState struct
type MSFT_ServerManagerRequestState struct {
	*cim.WmiInstance

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

func NewMSFT_ServerManagerRequestStateEx1(instance *cim.WmiInstance) (newInstance *MSFT_ServerManagerRequestState, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_ServerManagerRequestState{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_ServerManagerRequestStateEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_ServerManagerRequestState, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_ServerManagerRequestState{
		WmiInstance: tmp,
	}
	return
}

// SetError sets the value of Error for the instance
func (instance *MSFT_ServerManagerRequestState) SetPropertyError(value MSFT_ServerManagerDeploymentError) (err error) {
	return instance.SetProperty("Error", (value))
}

// GetError gets the value of Error for the instance
func (instance *MSFT_ServerManagerRequestState) GetPropertyError() (value MSFT_ServerManagerDeploymentError, err error) {
	retValue, err := instance.GetProperty("Error")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSFT_ServerManagerDeploymentError)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSFT_ServerManagerDeploymentError is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSFT_ServerManagerDeploymentError(valuetmp)

	return
}

// SetProgressTicks sets the value of ProgressTicks for the instance
func (instance *MSFT_ServerManagerRequestState) SetPropertyProgressTicks(value uint32) (err error) {
	return instance.SetProperty("ProgressTicks", (value))
}

// GetProgressTicks gets the value of ProgressTicks for the instance
func (instance *MSFT_ServerManagerRequestState) GetPropertyProgressTicks() (value uint32, err error) {
	retValue, err := instance.GetProperty("ProgressTicks")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetRequestState sets the value of RequestState for the instance
func (instance *MSFT_ServerManagerRequestState) SetPropertyRequestState(value uint8) (err error) {
	return instance.SetProperty("RequestState", (value))
}

// GetRequestState gets the value of RequestState for the instance
func (instance *MSFT_ServerManagerRequestState) GetPropertyRequestState() (value uint8, err error) {
	retValue, err := instance.GetProperty("RequestState")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetRestartRequired sets the value of RestartRequired for the instance
func (instance *MSFT_ServerManagerRequestState) SetPropertyRestartRequired(value bool) (err error) {
	return instance.SetProperty("RestartRequired", (value))
}

// GetRestartRequired gets the value of RestartRequired for the instance
func (instance *MSFT_ServerManagerRequestState) GetPropertyRestartRequired() (value bool, err error) {
	retValue, err := instance.GetProperty("RestartRequired")
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

// SetTotalTicks sets the value of TotalTicks for the instance
func (instance *MSFT_ServerManagerRequestState) SetPropertyTotalTicks(value uint32) (err error) {
	return instance.SetProperty("TotalTicks", (value))
}

// GetTotalTicks gets the value of TotalTicks for the instance
func (instance *MSFT_ServerManagerRequestState) GetPropertyTotalTicks() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalTicks")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetWarnings sets the value of Warnings for the instance
func (instance *MSFT_ServerManagerRequestState) SetPropertyWarnings(value []string) (err error) {
	return instance.SetProperty("Warnings", (value))
}

// GetWarnings gets the value of Warnings for the instance
func (instance *MSFT_ServerManagerRequestState) GetPropertyWarnings() (value []string, err error) {
	retValue, err := instance.GetProperty("Warnings")
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
