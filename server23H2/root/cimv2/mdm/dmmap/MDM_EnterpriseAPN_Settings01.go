// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
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

// MDM_EnterpriseAPN_Settings01 struct
type MDM_EnterpriseAPN_Settings01 struct {
	*cim.WmiInstance

	//
	AllowUserControl bool

	//
	HideView bool

	//
	InstanceID string

	//
	ParentID string
}

func NewMDM_EnterpriseAPN_Settings01Ex1(instance *cim.WmiInstance) (newInstance *MDM_EnterpriseAPN_Settings01, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_EnterpriseAPN_Settings01{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_EnterpriseAPN_Settings01Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_EnterpriseAPN_Settings01, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_EnterpriseAPN_Settings01{
		WmiInstance: tmp,
	}
	return
}

// SetAllowUserControl sets the value of AllowUserControl for the instance
func (instance *MDM_EnterpriseAPN_Settings01) SetPropertyAllowUserControl(value bool) (err error) {
	return instance.SetProperty("AllowUserControl", (value))
}

// GetAllowUserControl gets the value of AllowUserControl for the instance
func (instance *MDM_EnterpriseAPN_Settings01) GetPropertyAllowUserControl() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowUserControl")
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

// SetHideView sets the value of HideView for the instance
func (instance *MDM_EnterpriseAPN_Settings01) SetPropertyHideView(value bool) (err error) {
	return instance.SetProperty("HideView", (value))
}

// GetHideView gets the value of HideView for the instance
func (instance *MDM_EnterpriseAPN_Settings01) GetPropertyHideView() (value bool, err error) {
	retValue, err := instance.GetProperty("HideView")
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
func (instance *MDM_EnterpriseAPN_Settings01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_EnterpriseAPN_Settings01) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_EnterpriseAPN_Settings01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_EnterpriseAPN_Settings01) GetPropertyParentID() (value string, err error) {
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
