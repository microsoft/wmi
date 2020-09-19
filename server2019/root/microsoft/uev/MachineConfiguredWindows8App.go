// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.Microsoft.Uev
//////////////////////////////////////////////
package uev

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MachineConfiguredWindows8App struct
type MachineConfiguredWindows8App struct {
	*cim.WmiInstance

	//
	DisplayName string

	//
	Enabled bool

	//
	Installed bool

	//
	PackageFamilyName string
}

func NewMachineConfiguredWindows8AppEx1(instance *cim.WmiInstance) (newInstance *MachineConfiguredWindows8App, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MachineConfiguredWindows8App{
		WmiInstance: tmp,
	}
	return
}

func NewMachineConfiguredWindows8AppEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MachineConfiguredWindows8App, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MachineConfiguredWindows8App{
		WmiInstance: tmp,
	}
	return
}

// SetDisplayName sets the value of DisplayName for the instance
func (instance *MachineConfiguredWindows8App) SetPropertyDisplayName(value string) (err error) {
	return instance.SetProperty("DisplayName", (value))
}

// GetDisplayName gets the value of DisplayName for the instance
func (instance *MachineConfiguredWindows8App) GetPropertyDisplayName() (value string, err error) {
	retValue, err := instance.GetProperty("DisplayName")
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

// SetEnabled sets the value of Enabled for the instance
func (instance *MachineConfiguredWindows8App) SetPropertyEnabled(value bool) (err error) {
	return instance.SetProperty("Enabled", (value))
}

// GetEnabled gets the value of Enabled for the instance
func (instance *MachineConfiguredWindows8App) GetPropertyEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("Enabled")
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

// SetInstalled sets the value of Installed for the instance
func (instance *MachineConfiguredWindows8App) SetPropertyInstalled(value bool) (err error) {
	return instance.SetProperty("Installed", (value))
}

// GetInstalled gets the value of Installed for the instance
func (instance *MachineConfiguredWindows8App) GetPropertyInstalled() (value bool, err error) {
	retValue, err := instance.GetProperty("Installed")
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

// SetPackageFamilyName sets the value of PackageFamilyName for the instance
func (instance *MachineConfiguredWindows8App) SetPropertyPackageFamilyName(value string) (err error) {
	return instance.SetProperty("PackageFamilyName", (value))
}

// GetPackageFamilyName gets the value of PackageFamilyName for the instance
func (instance *MachineConfiguredWindows8App) GetPropertyPackageFamilyName() (value string, err error) {
	retValue, err := instance.GetProperty("PackageFamilyName")
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

//

// <param name="packageFamilyName" type="string "></param>
func (instance *MachineConfiguredWindows8App) EnableApp( /* IN */ packageFamilyName string) (err error) {
	_, err = instance.InvokeMethodWithReturn("EnableApp", packageFamilyName)
	if err != nil {
		return
	}
	return

}

//

// <param name="packageFamilyName" type="string "></param>
func (instance *MachineConfiguredWindows8App) DisableApp( /* IN */ packageFamilyName string) (err error) {
	_, err = instance.InvokeMethodWithReturn("DisableApp", packageFamilyName)
	if err != nil {
		return
	}
	return

}

//

// <param name="packageFamilyName" type="string "></param>
func (instance *MachineConfiguredWindows8App) RemoveApp( /* IN */ packageFamilyName string) (err error) {
	_, err = instance.InvokeMethodWithReturn("RemoveApp", packageFamilyName)
	if err != nil {
		return
	}
	return

}
