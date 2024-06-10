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

// MDM_Policy_Config01_HumanPresence02 struct
type MDM_Policy_Config01_HumanPresence02 struct {
	*cim.WmiInstance

	//
	ForceInstantLock int32

	//
	ForceInstantWake int32

	//
	ForceLockTimeout int32

	//
	InstanceID string

	//
	ParentID string
}

func NewMDM_Policy_Config01_HumanPresence02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_Config01_HumanPresence02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Config01_HumanPresence02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Policy_Config01_HumanPresence02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Policy_Config01_HumanPresence02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Config01_HumanPresence02{
		WmiInstance: tmp,
	}
	return
}

// SetForceInstantLock sets the value of ForceInstantLock for the instance
func (instance *MDM_Policy_Config01_HumanPresence02) SetPropertyForceInstantLock(value int32) (err error) {
	return instance.SetProperty("ForceInstantLock", (value))
}

// GetForceInstantLock gets the value of ForceInstantLock for the instance
func (instance *MDM_Policy_Config01_HumanPresence02) GetPropertyForceInstantLock() (value int32, err error) {
	retValue, err := instance.GetProperty("ForceInstantLock")
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

// SetForceInstantWake sets the value of ForceInstantWake for the instance
func (instance *MDM_Policy_Config01_HumanPresence02) SetPropertyForceInstantWake(value int32) (err error) {
	return instance.SetProperty("ForceInstantWake", (value))
}

// GetForceInstantWake gets the value of ForceInstantWake for the instance
func (instance *MDM_Policy_Config01_HumanPresence02) GetPropertyForceInstantWake() (value int32, err error) {
	retValue, err := instance.GetProperty("ForceInstantWake")
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

// SetForceLockTimeout sets the value of ForceLockTimeout for the instance
func (instance *MDM_Policy_Config01_HumanPresence02) SetPropertyForceLockTimeout(value int32) (err error) {
	return instance.SetProperty("ForceLockTimeout", (value))
}

// GetForceLockTimeout gets the value of ForceLockTimeout for the instance
func (instance *MDM_Policy_Config01_HumanPresence02) GetPropertyForceLockTimeout() (value int32, err error) {
	retValue, err := instance.GetProperty("ForceLockTimeout")
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

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Policy_Config01_HumanPresence02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Config01_HumanPresence02) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_Policy_Config01_HumanPresence02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Config01_HumanPresence02) GetPropertyParentID() (value string, err error) {
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
