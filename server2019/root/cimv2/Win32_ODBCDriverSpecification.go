// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_ODBCDriverSpecification struct
type Win32_ODBCDriverSpecification struct {
	*CIM_Check

	//
	Driver string

	//
	File string

	//
	SetupFile string
}

func NewWin32_ODBCDriverSpecificationEx1(instance *cim.WmiInstance) (newInstance *Win32_ODBCDriverSpecification, err error) {
	tmp, err := NewCIM_CheckEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_ODBCDriverSpecification{
		CIM_Check: tmp,
	}
	return
}

func NewWin32_ODBCDriverSpecificationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_ODBCDriverSpecification, err error) {
	tmp, err := NewCIM_CheckEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_ODBCDriverSpecification{
		CIM_Check: tmp,
	}
	return
}

// SetDriver sets the value of Driver for the instance
func (instance *Win32_ODBCDriverSpecification) SetPropertyDriver(value string) (err error) {
	return instance.SetProperty("Driver", value)
}

// GetDriver gets the value of Driver for the instance
func (instance *Win32_ODBCDriverSpecification) GetPropertyDriver() (value string, err error) {
	retValue, err := instance.GetProperty("Driver")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFile sets the value of File for the instance
func (instance *Win32_ODBCDriverSpecification) SetPropertyFile(value string) (err error) {
	return instance.SetProperty("File", value)
}

// GetFile gets the value of File for the instance
func (instance *Win32_ODBCDriverSpecification) GetPropertyFile() (value string, err error) {
	retValue, err := instance.GetProperty("File")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSetupFile sets the value of SetupFile for the instance
func (instance *Win32_ODBCDriverSpecification) SetPropertySetupFile(value string) (err error) {
	return instance.SetProperty("SetupFile", value)
}

// GetSetupFile gets the value of SetupFile for the instance
func (instance *Win32_ODBCDriverSpecification) GetPropertySetupFile() (value string, err error) {
	retValue, err := instance.GetProperty("SetupFile")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
