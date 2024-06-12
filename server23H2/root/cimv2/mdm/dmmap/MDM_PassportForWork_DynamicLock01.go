// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
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

// MDM_PassportForWork_DynamicLock01 struct
type MDM_PassportForWork_DynamicLock01 struct {
	*cim.WmiInstance

	//
	DynamicLock bool

	//
	InstanceID string

	//
	ParentID string

	//
	Plugins string
}

func NewMDM_PassportForWork_DynamicLock01Ex1(instance *cim.WmiInstance) (newInstance *MDM_PassportForWork_DynamicLock01, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_PassportForWork_DynamicLock01{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_PassportForWork_DynamicLock01Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_PassportForWork_DynamicLock01, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_PassportForWork_DynamicLock01{
		WmiInstance: tmp,
	}
	return
}

// SetDynamicLock sets the value of DynamicLock for the instance
func (instance *MDM_PassportForWork_DynamicLock01) SetPropertyDynamicLock(value bool) (err error) {
	return instance.SetProperty("DynamicLock", (value))
}

// GetDynamicLock gets the value of DynamicLock for the instance
func (instance *MDM_PassportForWork_DynamicLock01) GetPropertyDynamicLock() (value bool, err error) {
	retValue, err := instance.GetProperty("DynamicLock")
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

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_PassportForWork_DynamicLock01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_PassportForWork_DynamicLock01) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_PassportForWork_DynamicLock01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_PassportForWork_DynamicLock01) GetPropertyParentID() (value string, err error) {
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

// SetPlugins sets the value of Plugins for the instance
func (instance *MDM_PassportForWork_DynamicLock01) SetPropertyPlugins(value string) (err error) {
	return instance.SetProperty("Plugins", (value))
}

// GetPlugins gets the value of Plugins for the instance
func (instance *MDM_PassportForWork_DynamicLock01) GetPropertyPlugins() (value string, err error) {
	retValue, err := instance.GetProperty("Plugins")
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
