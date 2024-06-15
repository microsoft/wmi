// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
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

// MSFT_ServerManagerDeploymentError struct
type MSFT_ServerManagerDeploymentError struct {
	*cim.WmiInstance

	//
	ErrorCategory uint8

	//
	ErrorId string

	//
	ErrorMessage string
}

func NewMSFT_ServerManagerDeploymentErrorEx1(instance *cim.WmiInstance) (newInstance *MSFT_ServerManagerDeploymentError, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_ServerManagerDeploymentError{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_ServerManagerDeploymentErrorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_ServerManagerDeploymentError, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_ServerManagerDeploymentError{
		WmiInstance: tmp,
	}
	return
}

// SetErrorCategory sets the value of ErrorCategory for the instance
func (instance *MSFT_ServerManagerDeploymentError) SetPropertyErrorCategory(value uint8) (err error) {
	return instance.SetProperty("ErrorCategory", (value))
}

// GetErrorCategory gets the value of ErrorCategory for the instance
func (instance *MSFT_ServerManagerDeploymentError) GetPropertyErrorCategory() (value uint8, err error) {
	retValue, err := instance.GetProperty("ErrorCategory")
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

// SetErrorId sets the value of ErrorId for the instance
func (instance *MSFT_ServerManagerDeploymentError) SetPropertyErrorId(value string) (err error) {
	return instance.SetProperty("ErrorId", (value))
}

// GetErrorId gets the value of ErrorId for the instance
func (instance *MSFT_ServerManagerDeploymentError) GetPropertyErrorId() (value string, err error) {
	retValue, err := instance.GetProperty("ErrorId")
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

// SetErrorMessage sets the value of ErrorMessage for the instance
func (instance *MSFT_ServerManagerDeploymentError) SetPropertyErrorMessage(value string) (err error) {
	return instance.SetProperty("ErrorMessage", (value))
}

// GetErrorMessage gets the value of ErrorMessage for the instance
func (instance *MSFT_ServerManagerDeploymentError) GetPropertyErrorMessage() (value string, err error) {
	retValue, err := instance.GetProperty("ErrorMessage")
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
