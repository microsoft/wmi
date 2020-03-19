// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.ServerManager
//////////////////////////////////////////////
package servermanager

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
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
	return instance.SetProperty("ErrorCategory", value)
}

// GetErrorCategory gets the value of ErrorCategory for the instance
func (instance *MSFT_ServerManagerDeploymentError) GetPropertyErrorCategory() (value uint8, err error) {
	retValue, err := instance.GetProperty("ErrorCategory")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetErrorId sets the value of ErrorId for the instance
func (instance *MSFT_ServerManagerDeploymentError) SetPropertyErrorId(value string) (err error) {
	return instance.SetProperty("ErrorId", value)
}

// GetErrorId gets the value of ErrorId for the instance
func (instance *MSFT_ServerManagerDeploymentError) GetPropertyErrorId() (value string, err error) {
	retValue, err := instance.GetProperty("ErrorId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetErrorMessage sets the value of ErrorMessage for the instance
func (instance *MSFT_ServerManagerDeploymentError) SetPropertyErrorMessage(value string) (err error) {
	return instance.SetProperty("ErrorMessage", value)
}

// GetErrorMessage gets the value of ErrorMessage for the instance
func (instance *MSFT_ServerManagerDeploymentError) GetPropertyErrorMessage() (value string, err error) {
	retValue, err := instance.GetProperty("ErrorMessage")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
