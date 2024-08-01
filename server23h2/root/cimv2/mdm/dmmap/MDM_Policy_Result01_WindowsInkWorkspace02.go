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

// MDM_Policy_Result01_WindowsInkWorkspace02 struct
type MDM_Policy_Result01_WindowsInkWorkspace02 struct {
	*cim.WmiInstance

	//
	AllowSuggestedAppsInWindowsInkWorkspace int32

	//
	AllowWindowsInkWorkspace int32

	//
	InstanceID string

	//
	ParentID string
}

func NewMDM_Policy_Result01_WindowsInkWorkspace02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_Result01_WindowsInkWorkspace02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Result01_WindowsInkWorkspace02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Policy_Result01_WindowsInkWorkspace02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Policy_Result01_WindowsInkWorkspace02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Result01_WindowsInkWorkspace02{
		WmiInstance: tmp,
	}
	return
}

// SetAllowSuggestedAppsInWindowsInkWorkspace sets the value of AllowSuggestedAppsInWindowsInkWorkspace for the instance
func (instance *MDM_Policy_Result01_WindowsInkWorkspace02) SetPropertyAllowSuggestedAppsInWindowsInkWorkspace(value int32) (err error) {
	return instance.SetProperty("AllowSuggestedAppsInWindowsInkWorkspace", (value))
}

// GetAllowSuggestedAppsInWindowsInkWorkspace gets the value of AllowSuggestedAppsInWindowsInkWorkspace for the instance
func (instance *MDM_Policy_Result01_WindowsInkWorkspace02) GetPropertyAllowSuggestedAppsInWindowsInkWorkspace() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowSuggestedAppsInWindowsInkWorkspace")
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

// SetAllowWindowsInkWorkspace sets the value of AllowWindowsInkWorkspace for the instance
func (instance *MDM_Policy_Result01_WindowsInkWorkspace02) SetPropertyAllowWindowsInkWorkspace(value int32) (err error) {
	return instance.SetProperty("AllowWindowsInkWorkspace", (value))
}

// GetAllowWindowsInkWorkspace gets the value of AllowWindowsInkWorkspace for the instance
func (instance *MDM_Policy_Result01_WindowsInkWorkspace02) GetPropertyAllowWindowsInkWorkspace() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowWindowsInkWorkspace")
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
func (instance *MDM_Policy_Result01_WindowsInkWorkspace02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_WindowsInkWorkspace02) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_Policy_Result01_WindowsInkWorkspace02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_WindowsInkWorkspace02) GetPropertyParentID() (value string, err error) {
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
