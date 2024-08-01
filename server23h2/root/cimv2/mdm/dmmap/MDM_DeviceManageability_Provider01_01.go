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

// MDM_DeviceManageability_Provider01_01 struct
type MDM_DeviceManageability_Provider01_01 struct {
	*cim.WmiInstance

	//
	ConfigInfo string

	//
	EnrollmentInfo string

	//
	InstanceID string

	//
	ParentID string
}

func NewMDM_DeviceManageability_Provider01_01Ex1(instance *cim.WmiInstance) (newInstance *MDM_DeviceManageability_Provider01_01, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_DeviceManageability_Provider01_01{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_DeviceManageability_Provider01_01Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_DeviceManageability_Provider01_01, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_DeviceManageability_Provider01_01{
		WmiInstance: tmp,
	}
	return
}

// SetConfigInfo sets the value of ConfigInfo for the instance
func (instance *MDM_DeviceManageability_Provider01_01) SetPropertyConfigInfo(value string) (err error) {
	return instance.SetProperty("ConfigInfo", (value))
}

// GetConfigInfo gets the value of ConfigInfo for the instance
func (instance *MDM_DeviceManageability_Provider01_01) GetPropertyConfigInfo() (value string, err error) {
	retValue, err := instance.GetProperty("ConfigInfo")
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

// SetEnrollmentInfo sets the value of EnrollmentInfo for the instance
func (instance *MDM_DeviceManageability_Provider01_01) SetPropertyEnrollmentInfo(value string) (err error) {
	return instance.SetProperty("EnrollmentInfo", (value))
}

// GetEnrollmentInfo gets the value of EnrollmentInfo for the instance
func (instance *MDM_DeviceManageability_Provider01_01) GetPropertyEnrollmentInfo() (value string, err error) {
	retValue, err := instance.GetProperty("EnrollmentInfo")
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
func (instance *MDM_DeviceManageability_Provider01_01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_DeviceManageability_Provider01_01) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_DeviceManageability_Provider01_01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_DeviceManageability_Provider01_01) GetPropertyParentID() (value string, err error) {
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
