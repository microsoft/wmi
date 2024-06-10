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

// MDM_ApplicationControl_PolicyIDs03 struct
type MDM_ApplicationControl_PolicyIDs03 struct {
	*cim.WmiInstance

	//
	InstanceID string

	//
	IsVariableLeaf bool

	//
	ParentID string

	//
	Value string
}

func NewMDM_ApplicationControl_PolicyIDs03Ex1(instance *cim.WmiInstance) (newInstance *MDM_ApplicationControl_PolicyIDs03, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_ApplicationControl_PolicyIDs03{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_ApplicationControl_PolicyIDs03Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_ApplicationControl_PolicyIDs03, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_ApplicationControl_PolicyIDs03{
		WmiInstance: tmp,
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_ApplicationControl_PolicyIDs03) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_ApplicationControl_PolicyIDs03) GetPropertyInstanceID() (value string, err error) {
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

// SetIsVariableLeaf sets the value of IsVariableLeaf for the instance
func (instance *MDM_ApplicationControl_PolicyIDs03) SetPropertyIsVariableLeaf(value bool) (err error) {
	return instance.SetProperty("IsVariableLeaf", (value))
}

// GetIsVariableLeaf gets the value of IsVariableLeaf for the instance
func (instance *MDM_ApplicationControl_PolicyIDs03) GetPropertyIsVariableLeaf() (value bool, err error) {
	retValue, err := instance.GetProperty("IsVariableLeaf")
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

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_ApplicationControl_PolicyIDs03) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_ApplicationControl_PolicyIDs03) GetPropertyParentID() (value string, err error) {
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

// SetValue sets the value of Value for the instance
func (instance *MDM_ApplicationControl_PolicyIDs03) SetPropertyValue(value string) (err error) {
	return instance.SetProperty("Value", (value))
}

// GetValue gets the value of Value for the instance
func (instance *MDM_ApplicationControl_PolicyIDs03) GetPropertyValue() (value string, err error) {
	retValue, err := instance.GetProperty("Value")
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
