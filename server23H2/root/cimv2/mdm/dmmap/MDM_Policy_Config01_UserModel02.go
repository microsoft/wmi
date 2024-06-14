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

// MDM_Policy_Config01_UserModel02 struct
type MDM_Policy_Config01_UserModel02 struct {
	*cim.WmiInstance

	//
	EnterpriseSupport string

	//
	InstanceID string

	//
	ParentID string

	//
	UserModel int32
}

func NewMDM_Policy_Config01_UserModel02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_Config01_UserModel02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Config01_UserModel02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Policy_Config01_UserModel02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Policy_Config01_UserModel02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Config01_UserModel02{
		WmiInstance: tmp,
	}
	return
}

// SetEnterpriseSupport sets the value of EnterpriseSupport for the instance
func (instance *MDM_Policy_Config01_UserModel02) SetPropertyEnterpriseSupport(value string) (err error) {
	return instance.SetProperty("EnterpriseSupport", (value))
}

// GetEnterpriseSupport gets the value of EnterpriseSupport for the instance
func (instance *MDM_Policy_Config01_UserModel02) GetPropertyEnterpriseSupport() (value string, err error) {
	retValue, err := instance.GetProperty("EnterpriseSupport")
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
func (instance *MDM_Policy_Config01_UserModel02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Config01_UserModel02) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_Policy_Config01_UserModel02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Config01_UserModel02) GetPropertyParentID() (value string, err error) {
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

// SetUserModel sets the value of UserModel for the instance
func (instance *MDM_Policy_Config01_UserModel02) SetPropertyUserModel(value int32) (err error) {
	return instance.SetProperty("UserModel", (value))
}

// GetUserModel gets the value of UserModel for the instance
func (instance *MDM_Policy_Config01_UserModel02) GetPropertyUserModel() (value int32, err error) {
	retValue, err := instance.GetProperty("UserModel")
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
