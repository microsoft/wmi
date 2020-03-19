// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_ClientApplicationSetting struct
type Win32_ClientApplicationSetting struct {
	*cim.WmiInstance

	//
	Application Win32_DCOMApplication

	//
	Client CIM_DataFile
}

func NewWin32_ClientApplicationSettingEx1(instance *cim.WmiInstance) (newInstance *Win32_ClientApplicationSetting, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &Win32_ClientApplicationSetting{
		WmiInstance: tmp,
	}
	return
}

func NewWin32_ClientApplicationSettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_ClientApplicationSetting, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_ClientApplicationSetting{
		WmiInstance: tmp,
	}
	return
}

// SetApplication sets the value of Application for the instance
func (instance *Win32_ClientApplicationSetting) SetPropertyApplication(value Win32_DCOMApplication) (err error) {
	return instance.SetProperty("Application", value)
}

// GetApplication gets the value of Application for the instance
func (instance *Win32_ClientApplicationSetting) GetPropertyApplication() (value Win32_DCOMApplication, err error) {
	retValue, err := instance.GetProperty("Application")
	if err != nil {
		return
	}
	value, ok := retValue.(Win32_DCOMApplication)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetClient sets the value of Client for the instance
func (instance *Win32_ClientApplicationSetting) SetPropertyClient(value CIM_DataFile) (err error) {
	return instance.SetProperty("Client", value)
}

// GetClient gets the value of Client for the instance
func (instance *Win32_ClientApplicationSetting) GetPropertyClient() (value CIM_DataFile, err error) {
	retValue, err := instance.GetProperty("Client")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_DataFile)
	if !ok {
		// TODO: Set an error
	}
	return
}
