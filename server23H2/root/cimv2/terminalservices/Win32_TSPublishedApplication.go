// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.CIMV2.TerminalServices
//////////////////////////////////////////////
package terminalservices

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_TSPublishedApplication struct
type Win32_TSPublishedApplication struct {
	*CIM_LogicalElement

	// Alias of the application
	Alias string

	// Command Line Arguments setting
	CommandLineSetting TSPublishedApplication_CommandLineSetting

	// Contents of the icon corresponding to the application
	IconContents []uint8

	// Index of the icon
	IconIndex int32

	// Path to the application icon
	IconPath string

	// Path to the application
	Path string

	// Whether application path is valid
	PathExists bool

	// Contents of the RDP file corresponding to the application
	RDPFileContents string

	// Command Line Arguments required for this application
	RequiredCommandLine string

	// Security Descriptor controlling access to the application, in SDDL Format. Empty string implies allow all access. Does not support DENY ACEs, or ACEs referring to non-domain users or groups.
	SecurityDescriptor string

	// Whether this application should be shown in the TS Web Access
	ShowInPortal bool

	// Virtual Path to the application
	VPath string
}

func NewWin32_TSPublishedApplicationEx1(instance *cim.WmiInstance) (newInstance *Win32_TSPublishedApplication, err error) {
	tmp, err := NewCIM_LogicalElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_TSPublishedApplication{
		CIM_LogicalElement: tmp,
	}
	return
}

func NewWin32_TSPublishedApplicationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_TSPublishedApplication, err error) {
	tmp, err := NewCIM_LogicalElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_TSPublishedApplication{
		CIM_LogicalElement: tmp,
	}
	return
}

// SetAlias sets the value of Alias for the instance
func (instance *Win32_TSPublishedApplication) SetPropertyAlias(value string) (err error) {
	return instance.SetProperty("Alias", (value))
}

// GetAlias gets the value of Alias for the instance
func (instance *Win32_TSPublishedApplication) GetPropertyAlias() (value string, err error) {
	retValue, err := instance.GetProperty("Alias")
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

// SetCommandLineSetting sets the value of CommandLineSetting for the instance
func (instance *Win32_TSPublishedApplication) SetPropertyCommandLineSetting(value TSPublishedApplication_CommandLineSetting) (err error) {
	return instance.SetProperty("CommandLineSetting", (value))
}

// GetCommandLineSetting gets the value of CommandLineSetting for the instance
func (instance *Win32_TSPublishedApplication) GetPropertyCommandLineSetting() (value TSPublishedApplication_CommandLineSetting, err error) {
	retValue, err := instance.GetProperty("CommandLineSetting")
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

	value = TSPublishedApplication_CommandLineSetting(valuetmp)

	return
}

// SetIconContents sets the value of IconContents for the instance
func (instance *Win32_TSPublishedApplication) SetPropertyIconContents(value []uint8) (err error) {
	return instance.SetProperty("IconContents", (value))
}

// GetIconContents gets the value of IconContents for the instance
func (instance *Win32_TSPublishedApplication) GetPropertyIconContents() (value []uint8, err error) {
	retValue, err := instance.GetProperty("IconContents")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint8)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint8(valuetmp))
	}

	return
}

// SetIconIndex sets the value of IconIndex for the instance
func (instance *Win32_TSPublishedApplication) SetPropertyIconIndex(value int32) (err error) {
	return instance.SetProperty("IconIndex", (value))
}

// GetIconIndex gets the value of IconIndex for the instance
func (instance *Win32_TSPublishedApplication) GetPropertyIconIndex() (value int32, err error) {
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
func (instance *Win32_TSPublishedApplication) SetPropertyIconPath(value string) (err error) {
	return instance.SetProperty("IconPath", (value))
}

// GetIconPath gets the value of IconPath for the instance
func (instance *Win32_TSPublishedApplication) GetPropertyIconPath() (value string, err error) {
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
func (instance *Win32_TSPublishedApplication) SetPropertyPath(value string) (err error) {
	return instance.SetProperty("Path", (value))
}

// GetPath gets the value of Path for the instance
func (instance *Win32_TSPublishedApplication) GetPropertyPath() (value string, err error) {
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

// SetPathExists sets the value of PathExists for the instance
func (instance *Win32_TSPublishedApplication) SetPropertyPathExists(value bool) (err error) {
	return instance.SetProperty("PathExists", (value))
}

// GetPathExists gets the value of PathExists for the instance
func (instance *Win32_TSPublishedApplication) GetPropertyPathExists() (value bool, err error) {
	retValue, err := instance.GetProperty("PathExists")
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

// SetRDPFileContents sets the value of RDPFileContents for the instance
func (instance *Win32_TSPublishedApplication) SetPropertyRDPFileContents(value string) (err error) {
	return instance.SetProperty("RDPFileContents", (value))
}

// GetRDPFileContents gets the value of RDPFileContents for the instance
func (instance *Win32_TSPublishedApplication) GetPropertyRDPFileContents() (value string, err error) {
	retValue, err := instance.GetProperty("RDPFileContents")
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

// SetRequiredCommandLine sets the value of RequiredCommandLine for the instance
func (instance *Win32_TSPublishedApplication) SetPropertyRequiredCommandLine(value string) (err error) {
	return instance.SetProperty("RequiredCommandLine", (value))
}

// GetRequiredCommandLine gets the value of RequiredCommandLine for the instance
func (instance *Win32_TSPublishedApplication) GetPropertyRequiredCommandLine() (value string, err error) {
	retValue, err := instance.GetProperty("RequiredCommandLine")
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

// SetSecurityDescriptor sets the value of SecurityDescriptor for the instance
func (instance *Win32_TSPublishedApplication) SetPropertySecurityDescriptor(value string) (err error) {
	return instance.SetProperty("SecurityDescriptor", (value))
}

// GetSecurityDescriptor gets the value of SecurityDescriptor for the instance
func (instance *Win32_TSPublishedApplication) GetPropertySecurityDescriptor() (value string, err error) {
	retValue, err := instance.GetProperty("SecurityDescriptor")
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

// SetShowInPortal sets the value of ShowInPortal for the instance
func (instance *Win32_TSPublishedApplication) SetPropertyShowInPortal(value bool) (err error) {
	return instance.SetProperty("ShowInPortal", (value))
}

// GetShowInPortal gets the value of ShowInPortal for the instance
func (instance *Win32_TSPublishedApplication) GetPropertyShowInPortal() (value bool, err error) {
	retValue, err := instance.GetProperty("ShowInPortal")
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

// SetVPath sets the value of VPath for the instance
func (instance *Win32_TSPublishedApplication) SetPropertyVPath(value string) (err error) {
	return instance.SetProperty("VPath", (value))
}

// GetVPath gets the value of VPath for the instance
func (instance *Win32_TSPublishedApplication) GetPropertyVPath() (value string, err error) {
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
