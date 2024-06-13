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

// MDM_Policy_Result01_WindowsLogon02 struct
type MDM_Policy_Result01_WindowsLogon02 struct {
	*cim.WmiInstance

	//
	AllowAutomaticRestartSignOn string

	//
	ConfigAutomaticRestartSignOn string

	//
	DisableLockScreenAppNotifications string

	//
	DontDisplayNetworkSelectionUI string

	//
	EnableFirstLogonAnimation int32

	//
	EnumerateLocalUsersOnDomainJoinedComputers string

	//
	HideFastUserSwitching int32

	//
	InstanceID string

	//
	ParentID string
}

func NewMDM_Policy_Result01_WindowsLogon02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_Result01_WindowsLogon02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Result01_WindowsLogon02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Policy_Result01_WindowsLogon02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Policy_Result01_WindowsLogon02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Result01_WindowsLogon02{
		WmiInstance: tmp,
	}
	return
}

// SetAllowAutomaticRestartSignOn sets the value of AllowAutomaticRestartSignOn for the instance
func (instance *MDM_Policy_Result01_WindowsLogon02) SetPropertyAllowAutomaticRestartSignOn(value string) (err error) {
	return instance.SetProperty("AllowAutomaticRestartSignOn", (value))
}

// GetAllowAutomaticRestartSignOn gets the value of AllowAutomaticRestartSignOn for the instance
func (instance *MDM_Policy_Result01_WindowsLogon02) GetPropertyAllowAutomaticRestartSignOn() (value string, err error) {
	retValue, err := instance.GetProperty("AllowAutomaticRestartSignOn")
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

// SetConfigAutomaticRestartSignOn sets the value of ConfigAutomaticRestartSignOn for the instance
func (instance *MDM_Policy_Result01_WindowsLogon02) SetPropertyConfigAutomaticRestartSignOn(value string) (err error) {
	return instance.SetProperty("ConfigAutomaticRestartSignOn", (value))
}

// GetConfigAutomaticRestartSignOn gets the value of ConfigAutomaticRestartSignOn for the instance
func (instance *MDM_Policy_Result01_WindowsLogon02) GetPropertyConfigAutomaticRestartSignOn() (value string, err error) {
	retValue, err := instance.GetProperty("ConfigAutomaticRestartSignOn")
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

// SetDisableLockScreenAppNotifications sets the value of DisableLockScreenAppNotifications for the instance
func (instance *MDM_Policy_Result01_WindowsLogon02) SetPropertyDisableLockScreenAppNotifications(value string) (err error) {
	return instance.SetProperty("DisableLockScreenAppNotifications", (value))
}

// GetDisableLockScreenAppNotifications gets the value of DisableLockScreenAppNotifications for the instance
func (instance *MDM_Policy_Result01_WindowsLogon02) GetPropertyDisableLockScreenAppNotifications() (value string, err error) {
	retValue, err := instance.GetProperty("DisableLockScreenAppNotifications")
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

// SetDontDisplayNetworkSelectionUI sets the value of DontDisplayNetworkSelectionUI for the instance
func (instance *MDM_Policy_Result01_WindowsLogon02) SetPropertyDontDisplayNetworkSelectionUI(value string) (err error) {
	return instance.SetProperty("DontDisplayNetworkSelectionUI", (value))
}

// GetDontDisplayNetworkSelectionUI gets the value of DontDisplayNetworkSelectionUI for the instance
func (instance *MDM_Policy_Result01_WindowsLogon02) GetPropertyDontDisplayNetworkSelectionUI() (value string, err error) {
	retValue, err := instance.GetProperty("DontDisplayNetworkSelectionUI")
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

// SetEnableFirstLogonAnimation sets the value of EnableFirstLogonAnimation for the instance
func (instance *MDM_Policy_Result01_WindowsLogon02) SetPropertyEnableFirstLogonAnimation(value int32) (err error) {
	return instance.SetProperty("EnableFirstLogonAnimation", (value))
}

// GetEnableFirstLogonAnimation gets the value of EnableFirstLogonAnimation for the instance
func (instance *MDM_Policy_Result01_WindowsLogon02) GetPropertyEnableFirstLogonAnimation() (value int32, err error) {
	retValue, err := instance.GetProperty("EnableFirstLogonAnimation")
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

// SetEnumerateLocalUsersOnDomainJoinedComputers sets the value of EnumerateLocalUsersOnDomainJoinedComputers for the instance
func (instance *MDM_Policy_Result01_WindowsLogon02) SetPropertyEnumerateLocalUsersOnDomainJoinedComputers(value string) (err error) {
	return instance.SetProperty("EnumerateLocalUsersOnDomainJoinedComputers", (value))
}

// GetEnumerateLocalUsersOnDomainJoinedComputers gets the value of EnumerateLocalUsersOnDomainJoinedComputers for the instance
func (instance *MDM_Policy_Result01_WindowsLogon02) GetPropertyEnumerateLocalUsersOnDomainJoinedComputers() (value string, err error) {
	retValue, err := instance.GetProperty("EnumerateLocalUsersOnDomainJoinedComputers")
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

// SetHideFastUserSwitching sets the value of HideFastUserSwitching for the instance
func (instance *MDM_Policy_Result01_WindowsLogon02) SetPropertyHideFastUserSwitching(value int32) (err error) {
	return instance.SetProperty("HideFastUserSwitching", (value))
}

// GetHideFastUserSwitching gets the value of HideFastUserSwitching for the instance
func (instance *MDM_Policy_Result01_WindowsLogon02) GetPropertyHideFastUserSwitching() (value int32, err error) {
	retValue, err := instance.GetProperty("HideFastUserSwitching")
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
func (instance *MDM_Policy_Result01_WindowsLogon02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_WindowsLogon02) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_Policy_Result01_WindowsLogon02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_WindowsLogon02) GetPropertyParentID() (value string, err error) {
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
