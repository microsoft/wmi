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

// MDM_Policy_User_Config01_ApplicationManagement02 struct
type MDM_Policy_User_Config01_ApplicationManagement02 struct {
	*cim.WmiInstance

	//
	InstanceID string

	//
	MSIAlwaysInstallWithElevatedPrivileges int32

	//
	ParentID string

	//
	RequirePrivateStoreOnly int32
}

func NewMDM_Policy_User_Config01_ApplicationManagement02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_User_Config01_ApplicationManagement02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_User_Config01_ApplicationManagement02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Policy_User_Config01_ApplicationManagement02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Policy_User_Config01_ApplicationManagement02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_User_Config01_ApplicationManagement02{
		WmiInstance: tmp,
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Policy_User_Config01_ApplicationManagement02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_User_Config01_ApplicationManagement02) GetPropertyInstanceID() (value string, err error) {
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

// SetMSIAlwaysInstallWithElevatedPrivileges sets the value of MSIAlwaysInstallWithElevatedPrivileges for the instance
func (instance *MDM_Policy_User_Config01_ApplicationManagement02) SetPropertyMSIAlwaysInstallWithElevatedPrivileges(value int32) (err error) {
	return instance.SetProperty("MSIAlwaysInstallWithElevatedPrivileges", (value))
}

// GetMSIAlwaysInstallWithElevatedPrivileges gets the value of MSIAlwaysInstallWithElevatedPrivileges for the instance
func (instance *MDM_Policy_User_Config01_ApplicationManagement02) GetPropertyMSIAlwaysInstallWithElevatedPrivileges() (value int32, err error) {
	retValue, err := instance.GetProperty("MSIAlwaysInstallWithElevatedPrivileges")
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

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_Policy_User_Config01_ApplicationManagement02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_User_Config01_ApplicationManagement02) GetPropertyParentID() (value string, err error) {
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

// SetRequirePrivateStoreOnly sets the value of RequirePrivateStoreOnly for the instance
func (instance *MDM_Policy_User_Config01_ApplicationManagement02) SetPropertyRequirePrivateStoreOnly(value int32) (err error) {
	return instance.SetProperty("RequirePrivateStoreOnly", (value))
}

// GetRequirePrivateStoreOnly gets the value of RequirePrivateStoreOnly for the instance
func (instance *MDM_Policy_User_Config01_ApplicationManagement02) GetPropertyRequirePrivateStoreOnly() (value int32, err error) {
	retValue, err := instance.GetProperty("RequirePrivateStoreOnly")
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
