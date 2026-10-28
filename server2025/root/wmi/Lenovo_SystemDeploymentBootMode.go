// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 10/28/2026 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Lenovo_SystemDeploymentBootMode struct
type Lenovo_SystemDeploymentBootMode struct {
	*Lenovo_BIOSElement

	//
	Active bool

	// System Deployment Boot Mode
	CurrentSetting string

	//
	InstanceName string
}

func NewLenovo_SystemDeploymentBootModeEx1(instance *cim.WmiInstance) (newInstance *Lenovo_SystemDeploymentBootMode, err error) {
	tmp, err := NewLenovo_BIOSElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Lenovo_SystemDeploymentBootMode{
		Lenovo_BIOSElement: tmp,
	}
	return
}

func NewLenovo_SystemDeploymentBootModeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Lenovo_SystemDeploymentBootMode, err error) {
	tmp, err := NewLenovo_BIOSElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Lenovo_SystemDeploymentBootMode{
		Lenovo_BIOSElement: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *Lenovo_SystemDeploymentBootMode) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *Lenovo_SystemDeploymentBootMode) GetPropertyActive() (value bool, err error) {
	retValue, err := instance.GetProperty("Active")
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

// SetCurrentSetting sets the value of CurrentSetting for the instance
func (instance *Lenovo_SystemDeploymentBootMode) SetPropertyCurrentSetting(value string) (err error) {
	return instance.SetProperty("CurrentSetting", (value))
}

// GetCurrentSetting gets the value of CurrentSetting for the instance
func (instance *Lenovo_SystemDeploymentBootMode) GetPropertyCurrentSetting() (value string, err error) {
	retValue, err := instance.GetProperty("CurrentSetting")
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *Lenovo_SystemDeploymentBootMode) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *Lenovo_SystemDeploymentBootMode) GetPropertyInstanceName() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceName")
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
