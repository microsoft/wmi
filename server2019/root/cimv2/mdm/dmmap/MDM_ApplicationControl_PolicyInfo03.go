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

// MDM_ApplicationControl_PolicyInfo03 struct
type MDM_ApplicationControl_PolicyInfo03 struct {
	*cim.WmiInstance

	//
	FriendlyName string

	//
	InstanceID string

	//
	IsAuthorized bool

	//
	IsBasePolicy bool

	//
	IsDeployed bool

	//
	IsEffective bool

	//
	IsSystemPolicy bool

	//
	ParentID string

	//
	Status int32

	//
	Version string
}

func NewMDM_ApplicationControl_PolicyInfo03Ex1(instance *cim.WmiInstance) (newInstance *MDM_ApplicationControl_PolicyInfo03, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_ApplicationControl_PolicyInfo03{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_ApplicationControl_PolicyInfo03Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_ApplicationControl_PolicyInfo03, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_ApplicationControl_PolicyInfo03{
		WmiInstance: tmp,
	}
	return
}

// SetFriendlyName sets the value of FriendlyName for the instance
func (instance *MDM_ApplicationControl_PolicyInfo03) SetPropertyFriendlyName(value string) (err error) {
	return instance.SetProperty("FriendlyName", (value))
}

// GetFriendlyName gets the value of FriendlyName for the instance
func (instance *MDM_ApplicationControl_PolicyInfo03) GetPropertyFriendlyName() (value string, err error) {
	retValue, err := instance.GetProperty("FriendlyName")
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
func (instance *MDM_ApplicationControl_PolicyInfo03) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_ApplicationControl_PolicyInfo03) GetPropertyInstanceID() (value string, err error) {
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

// SetIsAuthorized sets the value of IsAuthorized for the instance
func (instance *MDM_ApplicationControl_PolicyInfo03) SetPropertyIsAuthorized(value bool) (err error) {
	return instance.SetProperty("IsAuthorized", (value))
}

// GetIsAuthorized gets the value of IsAuthorized for the instance
func (instance *MDM_ApplicationControl_PolicyInfo03) GetPropertyIsAuthorized() (value bool, err error) {
	retValue, err := instance.GetProperty("IsAuthorized")
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

// SetIsBasePolicy sets the value of IsBasePolicy for the instance
func (instance *MDM_ApplicationControl_PolicyInfo03) SetPropertyIsBasePolicy(value bool) (err error) {
	return instance.SetProperty("IsBasePolicy", (value))
}

// GetIsBasePolicy gets the value of IsBasePolicy for the instance
func (instance *MDM_ApplicationControl_PolicyInfo03) GetPropertyIsBasePolicy() (value bool, err error) {
	retValue, err := instance.GetProperty("IsBasePolicy")
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

// SetIsDeployed sets the value of IsDeployed for the instance
func (instance *MDM_ApplicationControl_PolicyInfo03) SetPropertyIsDeployed(value bool) (err error) {
	return instance.SetProperty("IsDeployed", (value))
}

// GetIsDeployed gets the value of IsDeployed for the instance
func (instance *MDM_ApplicationControl_PolicyInfo03) GetPropertyIsDeployed() (value bool, err error) {
	retValue, err := instance.GetProperty("IsDeployed")
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

// SetIsEffective sets the value of IsEffective for the instance
func (instance *MDM_ApplicationControl_PolicyInfo03) SetPropertyIsEffective(value bool) (err error) {
	return instance.SetProperty("IsEffective", (value))
}

// GetIsEffective gets the value of IsEffective for the instance
func (instance *MDM_ApplicationControl_PolicyInfo03) GetPropertyIsEffective() (value bool, err error) {
	retValue, err := instance.GetProperty("IsEffective")
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

// SetIsSystemPolicy sets the value of IsSystemPolicy for the instance
func (instance *MDM_ApplicationControl_PolicyInfo03) SetPropertyIsSystemPolicy(value bool) (err error) {
	return instance.SetProperty("IsSystemPolicy", (value))
}

// GetIsSystemPolicy gets the value of IsSystemPolicy for the instance
func (instance *MDM_ApplicationControl_PolicyInfo03) GetPropertyIsSystemPolicy() (value bool, err error) {
	retValue, err := instance.GetProperty("IsSystemPolicy")
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
func (instance *MDM_ApplicationControl_PolicyInfo03) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_ApplicationControl_PolicyInfo03) GetPropertyParentID() (value string, err error) {
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

// SetStatus sets the value of Status for the instance
func (instance *MDM_ApplicationControl_PolicyInfo03) SetPropertyStatus(value int32) (err error) {
	return instance.SetProperty("Status", (value))
}

// GetStatus gets the value of Status for the instance
func (instance *MDM_ApplicationControl_PolicyInfo03) GetPropertyStatus() (value int32, err error) {
	retValue, err := instance.GetProperty("Status")
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

// SetVersion sets the value of Version for the instance
func (instance *MDM_ApplicationControl_PolicyInfo03) SetPropertyVersion(value string) (err error) {
	return instance.SetProperty("Version", (value))
}

// GetVersion gets the value of Version for the instance
func (instance *MDM_ApplicationControl_PolicyInfo03) GetPropertyVersion() (value string, err error) {
	retValue, err := instance.GetProperty("Version")
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
