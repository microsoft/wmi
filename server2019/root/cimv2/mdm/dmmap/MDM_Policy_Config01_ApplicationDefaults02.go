// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
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

// MDM_Policy_Config01_ApplicationDefaults02 struct
type MDM_Policy_Config01_ApplicationDefaults02 struct {
	*cim.WmiInstance

	//
	DefaultAssociationsConfiguration string

	//
	EnableAppUriHandlers int32

	//
	InstanceID string

	//
	ParentID string
}

func NewMDM_Policy_Config01_ApplicationDefaults02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_Config01_ApplicationDefaults02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Config01_ApplicationDefaults02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Policy_Config01_ApplicationDefaults02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Policy_Config01_ApplicationDefaults02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Config01_ApplicationDefaults02{
		WmiInstance: tmp,
	}
	return
}

// SetDefaultAssociationsConfiguration sets the value of DefaultAssociationsConfiguration for the instance
func (instance *MDM_Policy_Config01_ApplicationDefaults02) SetPropertyDefaultAssociationsConfiguration(value string) (err error) {
	return instance.SetProperty("DefaultAssociationsConfiguration", (value))
}

// GetDefaultAssociationsConfiguration gets the value of DefaultAssociationsConfiguration for the instance
func (instance *MDM_Policy_Config01_ApplicationDefaults02) GetPropertyDefaultAssociationsConfiguration() (value string, err error) {
	retValue, err := instance.GetProperty("DefaultAssociationsConfiguration")
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

// SetEnableAppUriHandlers sets the value of EnableAppUriHandlers for the instance
func (instance *MDM_Policy_Config01_ApplicationDefaults02) SetPropertyEnableAppUriHandlers(value int32) (err error) {
	return instance.SetProperty("EnableAppUriHandlers", (value))
}

// GetEnableAppUriHandlers gets the value of EnableAppUriHandlers for the instance
func (instance *MDM_Policy_Config01_ApplicationDefaults02) GetPropertyEnableAppUriHandlers() (value int32, err error) {
	retValue, err := instance.GetProperty("EnableAppUriHandlers")
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
func (instance *MDM_Policy_Config01_ApplicationDefaults02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Config01_ApplicationDefaults02) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_Policy_Config01_ApplicationDefaults02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Config01_ApplicationDefaults02) GetPropertyParentID() (value string, err error) {
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
