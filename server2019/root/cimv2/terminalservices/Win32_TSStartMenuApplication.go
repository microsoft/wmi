// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.CIMV2.TerminalServices
//
// ////////////////////////////////////////////
package terminalservices

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_TSStartMenuApplication struct
type Win32_TSStartMenuApplication struct {
	*CIM_LogicalElement

	// Command line arguments
	CommandLineArguments string

	// Index of the icon
	IconIndex int32

	// Path to the application icon
	IconPath string

	// Path to the application
	Path string

	// Virtual Path to the application (includes Environment Variables)
	VPath string
}

func NewWin32_TSStartMenuApplicationEx1(instance *cim.WmiInstance) (newInstance *Win32_TSStartMenuApplication, err error) {
	tmp, err := NewCIM_LogicalElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_TSStartMenuApplication{
		CIM_LogicalElement: tmp,
	}
	return
}

func NewWin32_TSStartMenuApplicationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_TSStartMenuApplication, err error) {
	tmp, err := NewCIM_LogicalElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_TSStartMenuApplication{
		CIM_LogicalElement: tmp,
	}
	return
}

// SetCommandLineArguments sets the value of CommandLineArguments for the instance
func (instance *Win32_TSStartMenuApplication) SetPropertyCommandLineArguments(value string) (err error) {
	return instance.SetProperty("CommandLineArguments", (value))
}

// GetCommandLineArguments gets the value of CommandLineArguments for the instance
func (instance *Win32_TSStartMenuApplication) GetPropertyCommandLineArguments() (value string, err error) {
	retValue, err := instance.GetProperty("CommandLineArguments")
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

// SetIconIndex sets the value of IconIndex for the instance
func (instance *Win32_TSStartMenuApplication) SetPropertyIconIndex(value int32) (err error) {
	return instance.SetProperty("IconIndex", (value))
}

// GetIconIndex gets the value of IconIndex for the instance
func (instance *Win32_TSStartMenuApplication) GetPropertyIconIndex() (value int32, err error) {
	retValue, err := instance.GetProperty("IconIndex")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// SetIconPath sets the value of IconPath for the instance
func (instance *Win32_TSStartMenuApplication) SetPropertyIconPath(value string) (err error) {
	return instance.SetProperty("IconPath", (value))
}

// GetIconPath gets the value of IconPath for the instance
func (instance *Win32_TSStartMenuApplication) GetPropertyIconPath() (value string, err error) {
	retValue, err := instance.GetProperty("IconPath")
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

// SetPath sets the value of Path for the instance
func (instance *Win32_TSStartMenuApplication) SetPropertyPath(value string) (err error) {
	return instance.SetProperty("Path", (value))
}

// GetPath gets the value of Path for the instance
func (instance *Win32_TSStartMenuApplication) GetPropertyPath() (value string, err error) {
	retValue, err := instance.GetProperty("Path")
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

// SetVPath sets the value of VPath for the instance
func (instance *Win32_TSStartMenuApplication) SetPropertyVPath(value string) (err error) {
	return instance.SetProperty("VPath", (value))
}

// GetVPath gets the value of VPath for the instance
func (instance *Win32_TSStartMenuApplication) GetPropertyVPath() (value string, err error) {
	retValue, err := instance.GetProperty("VPath")
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
