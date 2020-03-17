// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Uev
//////////////////////////////////////////////
package uev

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// UserConfiguredWindows8App struct
type UserConfiguredWindows8App struct {
	cim.WmiInstance

	//
	DisplayName string

	//
	Enabled bool

	//
	Installed bool

	//
	PackageFamilyName string
}

// SetDisplayName sets the value of DisplayName for the instance
func (instance *UserConfiguredWindows8App) SetPropertyDisplayName(value string) (err error) {
	return instance.SetProperty("DisplayName", value)
}

// GetDisplayName gets the value of DisplayName for the instance
func (instance *UserConfiguredWindows8App) GetPropertyDisplayName() (value string, err error) {
	retValue, err := instance.GetProperty("DisplayName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnabled sets the value of Enabled for the instance
func (instance *UserConfiguredWindows8App) SetPropertyEnabled(value bool) (err error) {
	return instance.SetProperty("Enabled", value)
}

// GetEnabled gets the value of Enabled for the instance
func (instance *UserConfiguredWindows8App) GetPropertyEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("Enabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstalled sets the value of Installed for the instance
func (instance *UserConfiguredWindows8App) SetPropertyInstalled(value bool) (err error) {
	return instance.SetProperty("Installed", value)
}

// GetInstalled gets the value of Installed for the instance
func (instance *UserConfiguredWindows8App) GetPropertyInstalled() (value bool, err error) {
	retValue, err := instance.GetProperty("Installed")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPackageFamilyName sets the value of PackageFamilyName for the instance
func (instance *UserConfiguredWindows8App) SetPropertyPackageFamilyName(value string) (err error) {
	return instance.SetProperty("PackageFamilyName", value)
}

// GetPackageFamilyName gets the value of PackageFamilyName for the instance
func (instance *UserConfiguredWindows8App) GetPropertyPackageFamilyName() (value string, err error) {
	retValue, err := instance.GetProperty("PackageFamilyName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="packageFamilyName" type="string "></param>
func (instance *UserConfiguredWindows8App) EnableApp( /* IN */ packageFamilyName string) (err error) {
	_, err = instance.InvokeMethodWithReturn("EnableApp", packageFamilyName)
	if err != nil {
		return
	}
	return

}

//

// <param name="packageFamilyName" type="string "></param>
func (instance *UserConfiguredWindows8App) DisableApp( /* IN */ packageFamilyName string) (err error) {
	_, err = instance.InvokeMethodWithReturn("DisableApp", packageFamilyName)
	if err != nil {
		return
	}
	return

}

//

// <param name="packageFamilyName" type="string "></param>
func (instance *UserConfiguredWindows8App) RemoveApp( /* IN */ packageFamilyName string) (err error) {
	_, err = instance.InvokeMethodWithReturn("RemoveApp", packageFamilyName)
	if err != nil {
		return
	}
	return

}

//

// <param name="packageFamilyName" type="string "></param>

// <param name="ReturnValue" type="bool "></param>
func (instance *UserConfiguredWindows8App) CheckApp( /* IN */ packageFamilyName string) (result bool, err error) {
	retVal, err := instance.InvokeMethodWithReturn("CheckApp", packageFamilyName)
	if err != nil {
		return
	}
	result = (retVal > 0)
	return

}
