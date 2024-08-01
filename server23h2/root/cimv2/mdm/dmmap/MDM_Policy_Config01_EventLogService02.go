// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MDM_Policy_Config01_EventLogService02 struct
type MDM_Policy_Config01_EventLogService02 struct {
	*cim.WmiInstance

	//
	ControlEventLogBehavior string

	//
	InstanceID string

	//
	ParentID string

	//
	SpecifyMaximumFileSizeApplicationLog string

	//
	SpecifyMaximumFileSizeSecurityLog string

	//
	SpecifyMaximumFileSizeSystemLog string
}

func NewMDM_Policy_Config01_EventLogService02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_Config01_EventLogService02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Config01_EventLogService02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Policy_Config01_EventLogService02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Policy_Config01_EventLogService02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Config01_EventLogService02{
		WmiInstance: tmp,
	}
	return
}

// SetControlEventLogBehavior sets the value of ControlEventLogBehavior for the instance
func (instance *MDM_Policy_Config01_EventLogService02) SetPropertyControlEventLogBehavior(value string) (err error) {
	return instance.SetProperty("ControlEventLogBehavior", (value))
}

// GetControlEventLogBehavior gets the value of ControlEventLogBehavior for the instance
func (instance *MDM_Policy_Config01_EventLogService02) GetPropertyControlEventLogBehavior() (value string, err error) {
	retValue, err := instance.GetProperty("ControlEventLogBehavior")
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

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Policy_Config01_EventLogService02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Config01_EventLogService02) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_Policy_Config01_EventLogService02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Config01_EventLogService02) GetPropertyParentID() (value string, err error) {
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

// SetSpecifyMaximumFileSizeApplicationLog sets the value of SpecifyMaximumFileSizeApplicationLog for the instance
func (instance *MDM_Policy_Config01_EventLogService02) SetPropertySpecifyMaximumFileSizeApplicationLog(value string) (err error) {
	return instance.SetProperty("SpecifyMaximumFileSizeApplicationLog", (value))
}

// GetSpecifyMaximumFileSizeApplicationLog gets the value of SpecifyMaximumFileSizeApplicationLog for the instance
func (instance *MDM_Policy_Config01_EventLogService02) GetPropertySpecifyMaximumFileSizeApplicationLog() (value string, err error) {
	retValue, err := instance.GetProperty("SpecifyMaximumFileSizeApplicationLog")
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

// SetSpecifyMaximumFileSizeSecurityLog sets the value of SpecifyMaximumFileSizeSecurityLog for the instance
func (instance *MDM_Policy_Config01_EventLogService02) SetPropertySpecifyMaximumFileSizeSecurityLog(value string) (err error) {
	return instance.SetProperty("SpecifyMaximumFileSizeSecurityLog", (value))
}

// GetSpecifyMaximumFileSizeSecurityLog gets the value of SpecifyMaximumFileSizeSecurityLog for the instance
func (instance *MDM_Policy_Config01_EventLogService02) GetPropertySpecifyMaximumFileSizeSecurityLog() (value string, err error) {
	retValue, err := instance.GetProperty("SpecifyMaximumFileSizeSecurityLog")
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

// SetSpecifyMaximumFileSizeSystemLog sets the value of SpecifyMaximumFileSizeSystemLog for the instance
func (instance *MDM_Policy_Config01_EventLogService02) SetPropertySpecifyMaximumFileSizeSystemLog(value string) (err error) {
	return instance.SetProperty("SpecifyMaximumFileSizeSystemLog", (value))
}

// GetSpecifyMaximumFileSizeSystemLog gets the value of SpecifyMaximumFileSizeSystemLog for the instance
func (instance *MDM_Policy_Config01_EventLogService02) GetPropertySpecifyMaximumFileSizeSystemLog() (value string, err error) {
	retValue, err := instance.GetProperty("SpecifyMaximumFileSizeSystemLog")
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
