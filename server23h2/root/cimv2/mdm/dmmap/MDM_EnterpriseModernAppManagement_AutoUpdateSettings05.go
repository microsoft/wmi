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

// MDM_EnterpriseModernAppManagement_AutoUpdateSettings05 struct
type MDM_EnterpriseModernAppManagement_AutoUpdateSettings05 struct {
	*cim.WmiInstance

	//
	AutomaticBackgroundTask bool

	//
	Disable bool

	//
	ForceUpdateFromAnyVersion bool

	//
	HoursBetweenUpdateChecks bool

	//
	InstanceID string

	//
	OnLaunchUpdateCheck bool

	//
	PackageSource string

	//
	ParentID string

	//
	ShowPrompt bool

	//
	UpdateBlocksActivation bool
}

func NewMDM_EnterpriseModernAppManagement_AutoUpdateSettings05Ex1(instance *cim.WmiInstance) (newInstance *MDM_EnterpriseModernAppManagement_AutoUpdateSettings05, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_EnterpriseModernAppManagement_AutoUpdateSettings05{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_EnterpriseModernAppManagement_AutoUpdateSettings05Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_EnterpriseModernAppManagement_AutoUpdateSettings05, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_EnterpriseModernAppManagement_AutoUpdateSettings05{
		WmiInstance: tmp,
	}
	return
}

// SetAutomaticBackgroundTask sets the value of AutomaticBackgroundTask for the instance
func (instance *MDM_EnterpriseModernAppManagement_AutoUpdateSettings05) SetPropertyAutomaticBackgroundTask(value bool) (err error) {
	return instance.SetProperty("AutomaticBackgroundTask", (value))
}

// GetAutomaticBackgroundTask gets the value of AutomaticBackgroundTask for the instance
func (instance *MDM_EnterpriseModernAppManagement_AutoUpdateSettings05) GetPropertyAutomaticBackgroundTask() (value bool, err error) {
	retValue, err := instance.GetProperty("AutomaticBackgroundTask")
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

// SetDisable sets the value of Disable for the instance
func (instance *MDM_EnterpriseModernAppManagement_AutoUpdateSettings05) SetPropertyDisable(value bool) (err error) {
	return instance.SetProperty("Disable", (value))
}

// GetDisable gets the value of Disable for the instance
func (instance *MDM_EnterpriseModernAppManagement_AutoUpdateSettings05) GetPropertyDisable() (value bool, err error) {
	retValue, err := instance.GetProperty("Disable")
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

// SetForceUpdateFromAnyVersion sets the value of ForceUpdateFromAnyVersion for the instance
func (instance *MDM_EnterpriseModernAppManagement_AutoUpdateSettings05) SetPropertyForceUpdateFromAnyVersion(value bool) (err error) {
	return instance.SetProperty("ForceUpdateFromAnyVersion", (value))
}

// GetForceUpdateFromAnyVersion gets the value of ForceUpdateFromAnyVersion for the instance
func (instance *MDM_EnterpriseModernAppManagement_AutoUpdateSettings05) GetPropertyForceUpdateFromAnyVersion() (value bool, err error) {
	retValue, err := instance.GetProperty("ForceUpdateFromAnyVersion")
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

// SetHoursBetweenUpdateChecks sets the value of HoursBetweenUpdateChecks for the instance
func (instance *MDM_EnterpriseModernAppManagement_AutoUpdateSettings05) SetPropertyHoursBetweenUpdateChecks(value bool) (err error) {
	return instance.SetProperty("HoursBetweenUpdateChecks", (value))
}

// GetHoursBetweenUpdateChecks gets the value of HoursBetweenUpdateChecks for the instance
func (instance *MDM_EnterpriseModernAppManagement_AutoUpdateSettings05) GetPropertyHoursBetweenUpdateChecks() (value bool, err error) {
	retValue, err := instance.GetProperty("HoursBetweenUpdateChecks")
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
func (instance *MDM_EnterpriseModernAppManagement_AutoUpdateSettings05) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_EnterpriseModernAppManagement_AutoUpdateSettings05) GetPropertyInstanceID() (value string, err error) {
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

// SetOnLaunchUpdateCheck sets the value of OnLaunchUpdateCheck for the instance
func (instance *MDM_EnterpriseModernAppManagement_AutoUpdateSettings05) SetPropertyOnLaunchUpdateCheck(value bool) (err error) {
	return instance.SetProperty("OnLaunchUpdateCheck", (value))
}

// GetOnLaunchUpdateCheck gets the value of OnLaunchUpdateCheck for the instance
func (instance *MDM_EnterpriseModernAppManagement_AutoUpdateSettings05) GetPropertyOnLaunchUpdateCheck() (value bool, err error) {
	retValue, err := instance.GetProperty("OnLaunchUpdateCheck")
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

// SetPackageSource sets the value of PackageSource for the instance
func (instance *MDM_EnterpriseModernAppManagement_AutoUpdateSettings05) SetPropertyPackageSource(value string) (err error) {
	return instance.SetProperty("PackageSource", (value))
}

// GetPackageSource gets the value of PackageSource for the instance
func (instance *MDM_EnterpriseModernAppManagement_AutoUpdateSettings05) GetPropertyPackageSource() (value string, err error) {
	retValue, err := instance.GetProperty("PackageSource")
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
func (instance *MDM_EnterpriseModernAppManagement_AutoUpdateSettings05) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_EnterpriseModernAppManagement_AutoUpdateSettings05) GetPropertyParentID() (value string, err error) {
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

// SetShowPrompt sets the value of ShowPrompt for the instance
func (instance *MDM_EnterpriseModernAppManagement_AutoUpdateSettings05) SetPropertyShowPrompt(value bool) (err error) {
	return instance.SetProperty("ShowPrompt", (value))
}

// GetShowPrompt gets the value of ShowPrompt for the instance
func (instance *MDM_EnterpriseModernAppManagement_AutoUpdateSettings05) GetPropertyShowPrompt() (value bool, err error) {
	retValue, err := instance.GetProperty("ShowPrompt")
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

// SetUpdateBlocksActivation sets the value of UpdateBlocksActivation for the instance
func (instance *MDM_EnterpriseModernAppManagement_AutoUpdateSettings05) SetPropertyUpdateBlocksActivation(value bool) (err error) {
	return instance.SetProperty("UpdateBlocksActivation", (value))
}

// GetUpdateBlocksActivation gets the value of UpdateBlocksActivation for the instance
func (instance *MDM_EnterpriseModernAppManagement_AutoUpdateSettings05) GetPropertyUpdateBlocksActivation() (value bool, err error) {
	retValue, err := instance.GetProperty("UpdateBlocksActivation")
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
