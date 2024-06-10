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

// MDM_Policy_Result01_ApplicationManagement02 struct
type MDM_Policy_Result01_ApplicationManagement02 struct {
	*cim.WmiInstance

	//
	AllowAllTrustedApps int32

	//
	AllowAppStoreAutoUpdate int32

	//
	AllowDeveloperUnlock int32

	//
	AllowGameDVR int32

	//
	AllowSharedUserAppData int32

	//
	DisableStoreOriginatedApps int32

	//
	InstanceID string

	//
	LaunchAppAfterLogOn string

	//
	MSIAllowUserControlOverInstall int32

	//
	MSIAlwaysInstallWithElevatedPrivileges int32

	//
	ParentID string

	//
	RequirePrivateStoreOnly int32

	//
	RestrictAppDataToSystemVolume int32

	//
	RestrictAppToSystemVolume int32

	//
	ScheduleForceRestartForUpdateFailures string
}

func NewMDM_Policy_Result01_ApplicationManagement02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_Result01_ApplicationManagement02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Result01_ApplicationManagement02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Policy_Result01_ApplicationManagement02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Policy_Result01_ApplicationManagement02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Result01_ApplicationManagement02{
		WmiInstance: tmp,
	}
	return
}

// SetAllowAllTrustedApps sets the value of AllowAllTrustedApps for the instance
func (instance *MDM_Policy_Result01_ApplicationManagement02) SetPropertyAllowAllTrustedApps(value int32) (err error) {
	return instance.SetProperty("AllowAllTrustedApps", (value))
}

// GetAllowAllTrustedApps gets the value of AllowAllTrustedApps for the instance
func (instance *MDM_Policy_Result01_ApplicationManagement02) GetPropertyAllowAllTrustedApps() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowAllTrustedApps")
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

// SetAllowAppStoreAutoUpdate sets the value of AllowAppStoreAutoUpdate for the instance
func (instance *MDM_Policy_Result01_ApplicationManagement02) SetPropertyAllowAppStoreAutoUpdate(value int32) (err error) {
	return instance.SetProperty("AllowAppStoreAutoUpdate", (value))
}

// GetAllowAppStoreAutoUpdate gets the value of AllowAppStoreAutoUpdate for the instance
func (instance *MDM_Policy_Result01_ApplicationManagement02) GetPropertyAllowAppStoreAutoUpdate() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowAppStoreAutoUpdate")
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

// SetAllowDeveloperUnlock sets the value of AllowDeveloperUnlock for the instance
func (instance *MDM_Policy_Result01_ApplicationManagement02) SetPropertyAllowDeveloperUnlock(value int32) (err error) {
	return instance.SetProperty("AllowDeveloperUnlock", (value))
}

// GetAllowDeveloperUnlock gets the value of AllowDeveloperUnlock for the instance
func (instance *MDM_Policy_Result01_ApplicationManagement02) GetPropertyAllowDeveloperUnlock() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowDeveloperUnlock")
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

// SetAllowGameDVR sets the value of AllowGameDVR for the instance
func (instance *MDM_Policy_Result01_ApplicationManagement02) SetPropertyAllowGameDVR(value int32) (err error) {
	return instance.SetProperty("AllowGameDVR", (value))
}

// GetAllowGameDVR gets the value of AllowGameDVR for the instance
func (instance *MDM_Policy_Result01_ApplicationManagement02) GetPropertyAllowGameDVR() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowGameDVR")
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

// SetAllowSharedUserAppData sets the value of AllowSharedUserAppData for the instance
func (instance *MDM_Policy_Result01_ApplicationManagement02) SetPropertyAllowSharedUserAppData(value int32) (err error) {
	return instance.SetProperty("AllowSharedUserAppData", (value))
}

// GetAllowSharedUserAppData gets the value of AllowSharedUserAppData for the instance
func (instance *MDM_Policy_Result01_ApplicationManagement02) GetPropertyAllowSharedUserAppData() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowSharedUserAppData")
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

// SetDisableStoreOriginatedApps sets the value of DisableStoreOriginatedApps for the instance
func (instance *MDM_Policy_Result01_ApplicationManagement02) SetPropertyDisableStoreOriginatedApps(value int32) (err error) {
	return instance.SetProperty("DisableStoreOriginatedApps", (value))
}

// GetDisableStoreOriginatedApps gets the value of DisableStoreOriginatedApps for the instance
func (instance *MDM_Policy_Result01_ApplicationManagement02) GetPropertyDisableStoreOriginatedApps() (value int32, err error) {
	retValue, err := instance.GetProperty("DisableStoreOriginatedApps")
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
func (instance *MDM_Policy_Result01_ApplicationManagement02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_ApplicationManagement02) GetPropertyInstanceID() (value string, err error) {
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

// SetLaunchAppAfterLogOn sets the value of LaunchAppAfterLogOn for the instance
func (instance *MDM_Policy_Result01_ApplicationManagement02) SetPropertyLaunchAppAfterLogOn(value string) (err error) {
	return instance.SetProperty("LaunchAppAfterLogOn", (value))
}

// GetLaunchAppAfterLogOn gets the value of LaunchAppAfterLogOn for the instance
func (instance *MDM_Policy_Result01_ApplicationManagement02) GetPropertyLaunchAppAfterLogOn() (value string, err error) {
	retValue, err := instance.GetProperty("LaunchAppAfterLogOn")
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

// SetMSIAllowUserControlOverInstall sets the value of MSIAllowUserControlOverInstall for the instance
func (instance *MDM_Policy_Result01_ApplicationManagement02) SetPropertyMSIAllowUserControlOverInstall(value int32) (err error) {
	return instance.SetProperty("MSIAllowUserControlOverInstall", (value))
}

// GetMSIAllowUserControlOverInstall gets the value of MSIAllowUserControlOverInstall for the instance
func (instance *MDM_Policy_Result01_ApplicationManagement02) GetPropertyMSIAllowUserControlOverInstall() (value int32, err error) {
	retValue, err := instance.GetProperty("MSIAllowUserControlOverInstall")
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

// SetMSIAlwaysInstallWithElevatedPrivileges sets the value of MSIAlwaysInstallWithElevatedPrivileges for the instance
func (instance *MDM_Policy_Result01_ApplicationManagement02) SetPropertyMSIAlwaysInstallWithElevatedPrivileges(value int32) (err error) {
	return instance.SetProperty("MSIAlwaysInstallWithElevatedPrivileges", (value))
}

// GetMSIAlwaysInstallWithElevatedPrivileges gets the value of MSIAlwaysInstallWithElevatedPrivileges for the instance
func (instance *MDM_Policy_Result01_ApplicationManagement02) GetPropertyMSIAlwaysInstallWithElevatedPrivileges() (value int32, err error) {
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
func (instance *MDM_Policy_Result01_ApplicationManagement02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_ApplicationManagement02) GetPropertyParentID() (value string, err error) {
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
func (instance *MDM_Policy_Result01_ApplicationManagement02) SetPropertyRequirePrivateStoreOnly(value int32) (err error) {
	return instance.SetProperty("RequirePrivateStoreOnly", (value))
}

// GetRequirePrivateStoreOnly gets the value of RequirePrivateStoreOnly for the instance
func (instance *MDM_Policy_Result01_ApplicationManagement02) GetPropertyRequirePrivateStoreOnly() (value int32, err error) {
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

// SetRestrictAppDataToSystemVolume sets the value of RestrictAppDataToSystemVolume for the instance
func (instance *MDM_Policy_Result01_ApplicationManagement02) SetPropertyRestrictAppDataToSystemVolume(value int32) (err error) {
	return instance.SetProperty("RestrictAppDataToSystemVolume", (value))
}

// GetRestrictAppDataToSystemVolume gets the value of RestrictAppDataToSystemVolume for the instance
func (instance *MDM_Policy_Result01_ApplicationManagement02) GetPropertyRestrictAppDataToSystemVolume() (value int32, err error) {
	retValue, err := instance.GetProperty("RestrictAppDataToSystemVolume")
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

// SetRestrictAppToSystemVolume sets the value of RestrictAppToSystemVolume for the instance
func (instance *MDM_Policy_Result01_ApplicationManagement02) SetPropertyRestrictAppToSystemVolume(value int32) (err error) {
	return instance.SetProperty("RestrictAppToSystemVolume", (value))
}

// GetRestrictAppToSystemVolume gets the value of RestrictAppToSystemVolume for the instance
func (instance *MDM_Policy_Result01_ApplicationManagement02) GetPropertyRestrictAppToSystemVolume() (value int32, err error) {
	retValue, err := instance.GetProperty("RestrictAppToSystemVolume")
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

// SetScheduleForceRestartForUpdateFailures sets the value of ScheduleForceRestartForUpdateFailures for the instance
func (instance *MDM_Policy_Result01_ApplicationManagement02) SetPropertyScheduleForceRestartForUpdateFailures(value string) (err error) {
	return instance.SetProperty("ScheduleForceRestartForUpdateFailures", (value))
}

// GetScheduleForceRestartForUpdateFailures gets the value of ScheduleForceRestartForUpdateFailures for the instance
func (instance *MDM_Policy_Result01_ApplicationManagement02) GetPropertyScheduleForceRestartForUpdateFailures() (value string, err error) {
	retValue, err := instance.GetProperty("ScheduleForceRestartForUpdateFailures")
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
