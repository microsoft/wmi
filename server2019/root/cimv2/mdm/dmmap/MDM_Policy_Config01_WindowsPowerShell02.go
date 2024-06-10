// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.CIMV2.mdm.dmmap
//
// ////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MDM_Policy_Config01_WindowsPowerShell02 struct
type MDM_Policy_Config01_WindowsPowerShell02 struct {
	*cim.WmiInstance

	//
	InstanceID string

	//
	ParentID string

	//
	TurnOnPowerShellScriptBlockLogging string
}

func NewMDM_Policy_Config01_WindowsPowerShell02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_Config01_WindowsPowerShell02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Config01_WindowsPowerShell02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Policy_Config01_WindowsPowerShell02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Policy_Config01_WindowsPowerShell02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Config01_WindowsPowerShell02{
		WmiInstance: tmp,
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Policy_Config01_WindowsPowerShell02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Config01_WindowsPowerShell02) GetPropertyInstanceID() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceID")
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

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_Policy_Config01_WindowsPowerShell02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Config01_WindowsPowerShell02) GetPropertyParentID() (value string, err error) {
	retValue, err := instance.GetProperty("ParentID")
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

// SetTurnOnPowerShellScriptBlockLogging sets the value of TurnOnPowerShellScriptBlockLogging for the instance
func (instance *MDM_Policy_Config01_WindowsPowerShell02) SetPropertyTurnOnPowerShellScriptBlockLogging(value string) (err error) {
	return instance.SetProperty("TurnOnPowerShellScriptBlockLogging", (value))
}

// GetTurnOnPowerShellScriptBlockLogging gets the value of TurnOnPowerShellScriptBlockLogging for the instance
func (instance *MDM_Policy_Config01_WindowsPowerShell02) GetPropertyTurnOnPowerShellScriptBlockLogging() (value string, err error) {
	retValue, err := instance.GetProperty("TurnOnPowerShellScriptBlockLogging")
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
