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

// MDM_Policy_Config01_WindowsDefenderSecurityCenter02 struct
type MDM_Policy_Config01_WindowsDefenderSecurityCenter02 struct {
	*cim.WmiInstance

	//
	CompanyName string

	//
	DisableAccountProtectionUI int32

	//
	DisableAppBrowserUI int32

	//
	DisableClearTpmButton int32

	//
	DisableDeviceSecurityUI int32

	//
	DisableEnhancedNotifications int32

	//
	DisableFamilyUI int32

	//
	DisableHealthUI int32

	//
	DisableNetworkUI int32

	//
	DisableNotifications int32

	//
	DisableTpmFirmwareUpdateWarning int32

	//
	DisableVirusUI int32

	//
	DisallowExploitProtectionOverride int32

	//
	Email string

	//
	EnableCustomizedToasts int32

	//
	EnableInAppCustomization int32

	//
	HideRansomwareDataRecovery int32

	//
	HideSecureBoot int32

	//
	HideTPMTroubleshooting int32

	//
	HideWindowsSecurityNotificationAreaControl int32

	//
	InstanceID string

	//
	ParentID string

	//
	Phone string

	//
	URL string
}

func NewMDM_Policy_Config01_WindowsDefenderSecurityCenter02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_Config01_WindowsDefenderSecurityCenter02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Config01_WindowsDefenderSecurityCenter02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Policy_Config01_WindowsDefenderSecurityCenter02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Policy_Config01_WindowsDefenderSecurityCenter02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Config01_WindowsDefenderSecurityCenter02{
		WmiInstance: tmp,
	}
	return
}

// SetCompanyName sets the value of CompanyName for the instance
func (instance *MDM_Policy_Config01_WindowsDefenderSecurityCenter02) SetPropertyCompanyName(value string) (err error) {
	return instance.SetProperty("CompanyName", (value))
}

// GetCompanyName gets the value of CompanyName for the instance
func (instance *MDM_Policy_Config01_WindowsDefenderSecurityCenter02) GetPropertyCompanyName() (value string, err error) {
	retValue, err := instance.GetProperty("CompanyName")
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

// SetDisableAccountProtectionUI sets the value of DisableAccountProtectionUI for the instance
func (instance *MDM_Policy_Config01_WindowsDefenderSecurityCenter02) SetPropertyDisableAccountProtectionUI(value int32) (err error) {
	return instance.SetProperty("DisableAccountProtectionUI", (value))
}

// GetDisableAccountProtectionUI gets the value of DisableAccountProtectionUI for the instance
func (instance *MDM_Policy_Config01_WindowsDefenderSecurityCenter02) GetPropertyDisableAccountProtectionUI() (value int32, err error) {
	retValue, err := instance.GetProperty("DisableAccountProtectionUI")
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

// SetDisableAppBrowserUI sets the value of DisableAppBrowserUI for the instance
func (instance *MDM_Policy_Config01_WindowsDefenderSecurityCenter02) SetPropertyDisableAppBrowserUI(value int32) (err error) {
	return instance.SetProperty("DisableAppBrowserUI", (value))
}

// GetDisableAppBrowserUI gets the value of DisableAppBrowserUI for the instance
func (instance *MDM_Policy_Config01_WindowsDefenderSecurityCenter02) GetPropertyDisableAppBrowserUI() (value int32, err error) {
	retValue, err := instance.GetProperty("DisableAppBrowserUI")
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

// SetDisableClearTpmButton sets the value of DisableClearTpmButton for the instance
func (instance *MDM_Policy_Config01_WindowsDefenderSecurityCenter02) SetPropertyDisableClearTpmButton(value int32) (err error) {
	return instance.SetProperty("DisableClearTpmButton", (value))
}

// GetDisableClearTpmButton gets the value of DisableClearTpmButton for the instance
func (instance *MDM_Policy_Config01_WindowsDefenderSecurityCenter02) GetPropertyDisableClearTpmButton() (value int32, err error) {
	retValue, err := instance.GetProperty("DisableClearTpmButton")
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

// SetDisableDeviceSecurityUI sets the value of DisableDeviceSecurityUI for the instance
func (instance *MDM_Policy_Config01_WindowsDefenderSecurityCenter02) SetPropertyDisableDeviceSecurityUI(value int32) (err error) {
	return instance.SetProperty("DisableDeviceSecurityUI", (value))
}

// GetDisableDeviceSecurityUI gets the value of DisableDeviceSecurityUI for the instance
func (instance *MDM_Policy_Config01_WindowsDefenderSecurityCenter02) GetPropertyDisableDeviceSecurityUI() (value int32, err error) {
	retValue, err := instance.GetProperty("DisableDeviceSecurityUI")
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

// SetDisableEnhancedNotifications sets the value of DisableEnhancedNotifications for the instance
func (instance *MDM_Policy_Config01_WindowsDefenderSecurityCenter02) SetPropertyDisableEnhancedNotifications(value int32) (err error) {
	return instance.SetProperty("DisableEnhancedNotifications", (value))
}

// GetDisableEnhancedNotifications gets the value of DisableEnhancedNotifications for the instance
func (instance *MDM_Policy_Config01_WindowsDefenderSecurityCenter02) GetPropertyDisableEnhancedNotifications() (value int32, err error) {
	retValue, err := instance.GetProperty("DisableEnhancedNotifications")
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

// SetDisableFamilyUI sets the value of DisableFamilyUI for the instance
func (instance *MDM_Policy_Config01_WindowsDefenderSecurityCenter02) SetPropertyDisableFamilyUI(value int32) (err error) {
	return instance.SetProperty("DisableFamilyUI", (value))
}

// GetDisableFamilyUI gets the value of DisableFamilyUI for the instance
func (instance *MDM_Policy_Config01_WindowsDefenderSecurityCenter02) GetPropertyDisableFamilyUI() (value int32, err error) {
	retValue, err := instance.GetProperty("DisableFamilyUI")
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

// SetDisableHealthUI sets the value of DisableHealthUI for the instance
func (instance *MDM_Policy_Config01_WindowsDefenderSecurityCenter02) SetPropertyDisableHealthUI(value int32) (err error) {
	return instance.SetProperty("DisableHealthUI", (value))
}

// GetDisableHealthUI gets the value of DisableHealthUI for the instance
func (instance *MDM_Policy_Config01_WindowsDefenderSecurityCenter02) GetPropertyDisableHealthUI() (value int32, err error) {
	retValue, err := instance.GetProperty("DisableHealthUI")
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

// SetDisableNetworkUI sets the value of DisableNetworkUI for the instance
func (instance *MDM_Policy_Config01_WindowsDefenderSecurityCenter02) SetPropertyDisableNetworkUI(value int32) (err error) {
	return instance.SetProperty("DisableNetworkUI", (value))
}

// GetDisableNetworkUI gets the value of DisableNetworkUI for the instance
func (instance *MDM_Policy_Config01_WindowsDefenderSecurityCenter02) GetPropertyDisableNetworkUI() (value int32, err error) {
	retValue, err := instance.GetProperty("DisableNetworkUI")
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

// SetDisableNotifications sets the value of DisableNotifications for the instance
func (instance *MDM_Policy_Config01_WindowsDefenderSecurityCenter02) SetPropertyDisableNotifications(value int32) (err error) {
	return instance.SetProperty("DisableNotifications", (value))
}

// GetDisableNotifications gets the value of DisableNotifications for the instance
func (instance *MDM_Policy_Config01_WindowsDefenderSecurityCenter02) GetPropertyDisableNotifications() (value int32, err error) {
	retValue, err := instance.GetProperty("DisableNotifications")
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

// SetDisableTpmFirmwareUpdateWarning sets the value of DisableTpmFirmwareUpdateWarning for the instance
func (instance *MDM_Policy_Config01_WindowsDefenderSecurityCenter02) SetPropertyDisableTpmFirmwareUpdateWarning(value int32) (err error) {
	return instance.SetProperty("DisableTpmFirmwareUpdateWarning", (value))
}

// GetDisableTpmFirmwareUpdateWarning gets the value of DisableTpmFirmwareUpdateWarning for the instance
func (instance *MDM_Policy_Config01_WindowsDefenderSecurityCenter02) GetPropertyDisableTpmFirmwareUpdateWarning() (value int32, err error) {
	retValue, err := instance.GetProperty("DisableTpmFirmwareUpdateWarning")
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

// SetDisableVirusUI sets the value of DisableVirusUI for the instance
func (instance *MDM_Policy_Config01_WindowsDefenderSecurityCenter02) SetPropertyDisableVirusUI(value int32) (err error) {
	return instance.SetProperty("DisableVirusUI", (value))
}

// GetDisableVirusUI gets the value of DisableVirusUI for the instance
func (instance *MDM_Policy_Config01_WindowsDefenderSecurityCenter02) GetPropertyDisableVirusUI() (value int32, err error) {
	retValue, err := instance.GetProperty("DisableVirusUI")
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

// SetDisallowExploitProtectionOverride sets the value of DisallowExploitProtectionOverride for the instance
func (instance *MDM_Policy_Config01_WindowsDefenderSecurityCenter02) SetPropertyDisallowExploitProtectionOverride(value int32) (err error) {
	return instance.SetProperty("DisallowExploitProtectionOverride", (value))
}

// GetDisallowExploitProtectionOverride gets the value of DisallowExploitProtectionOverride for the instance
func (instance *MDM_Policy_Config01_WindowsDefenderSecurityCenter02) GetPropertyDisallowExploitProtectionOverride() (value int32, err error) {
	retValue, err := instance.GetProperty("DisallowExploitProtectionOverride")
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

// SetEmail sets the value of Email for the instance
func (instance *MDM_Policy_Config01_WindowsDefenderSecurityCenter02) SetPropertyEmail(value string) (err error) {
	return instance.SetProperty("Email", (value))
}

// GetEmail gets the value of Email for the instance
func (instance *MDM_Policy_Config01_WindowsDefenderSecurityCenter02) GetPropertyEmail() (value string, err error) {
	retValue, err := instance.GetProperty("Email")
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

// SetEnableCustomizedToasts sets the value of EnableCustomizedToasts for the instance
func (instance *MDM_Policy_Config01_WindowsDefenderSecurityCenter02) SetPropertyEnableCustomizedToasts(value int32) (err error) {
	return instance.SetProperty("EnableCustomizedToasts", (value))
}

// GetEnableCustomizedToasts gets the value of EnableCustomizedToasts for the instance
func (instance *MDM_Policy_Config01_WindowsDefenderSecurityCenter02) GetPropertyEnableCustomizedToasts() (value int32, err error) {
	retValue, err := instance.GetProperty("EnableCustomizedToasts")
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

// SetEnableInAppCustomization sets the value of EnableInAppCustomization for the instance
func (instance *MDM_Policy_Config01_WindowsDefenderSecurityCenter02) SetPropertyEnableInAppCustomization(value int32) (err error) {
	return instance.SetProperty("EnableInAppCustomization", (value))
}

// GetEnableInAppCustomization gets the value of EnableInAppCustomization for the instance
func (instance *MDM_Policy_Config01_WindowsDefenderSecurityCenter02) GetPropertyEnableInAppCustomization() (value int32, err error) {
	retValue, err := instance.GetProperty("EnableInAppCustomization")
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

// SetHideRansomwareDataRecovery sets the value of HideRansomwareDataRecovery for the instance
func (instance *MDM_Policy_Config01_WindowsDefenderSecurityCenter02) SetPropertyHideRansomwareDataRecovery(value int32) (err error) {
	return instance.SetProperty("HideRansomwareDataRecovery", (value))
}

// GetHideRansomwareDataRecovery gets the value of HideRansomwareDataRecovery for the instance
func (instance *MDM_Policy_Config01_WindowsDefenderSecurityCenter02) GetPropertyHideRansomwareDataRecovery() (value int32, err error) {
	retValue, err := instance.GetProperty("HideRansomwareDataRecovery")
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

// SetHideSecureBoot sets the value of HideSecureBoot for the instance
func (instance *MDM_Policy_Config01_WindowsDefenderSecurityCenter02) SetPropertyHideSecureBoot(value int32) (err error) {
	return instance.SetProperty("HideSecureBoot", (value))
}

// GetHideSecureBoot gets the value of HideSecureBoot for the instance
func (instance *MDM_Policy_Config01_WindowsDefenderSecurityCenter02) GetPropertyHideSecureBoot() (value int32, err error) {
	retValue, err := instance.GetProperty("HideSecureBoot")
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

// SetHideTPMTroubleshooting sets the value of HideTPMTroubleshooting for the instance
func (instance *MDM_Policy_Config01_WindowsDefenderSecurityCenter02) SetPropertyHideTPMTroubleshooting(value int32) (err error) {
	return instance.SetProperty("HideTPMTroubleshooting", (value))
}

// GetHideTPMTroubleshooting gets the value of HideTPMTroubleshooting for the instance
func (instance *MDM_Policy_Config01_WindowsDefenderSecurityCenter02) GetPropertyHideTPMTroubleshooting() (value int32, err error) {
	retValue, err := instance.GetProperty("HideTPMTroubleshooting")
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

// SetHideWindowsSecurityNotificationAreaControl sets the value of HideWindowsSecurityNotificationAreaControl for the instance
func (instance *MDM_Policy_Config01_WindowsDefenderSecurityCenter02) SetPropertyHideWindowsSecurityNotificationAreaControl(value int32) (err error) {
	return instance.SetProperty("HideWindowsSecurityNotificationAreaControl", (value))
}

// GetHideWindowsSecurityNotificationAreaControl gets the value of HideWindowsSecurityNotificationAreaControl for the instance
func (instance *MDM_Policy_Config01_WindowsDefenderSecurityCenter02) GetPropertyHideWindowsSecurityNotificationAreaControl() (value int32, err error) {
	retValue, err := instance.GetProperty("HideWindowsSecurityNotificationAreaControl")
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
func (instance *MDM_Policy_Config01_WindowsDefenderSecurityCenter02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Config01_WindowsDefenderSecurityCenter02) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_Policy_Config01_WindowsDefenderSecurityCenter02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Config01_WindowsDefenderSecurityCenter02) GetPropertyParentID() (value string, err error) {
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

// SetPhone sets the value of Phone for the instance
func (instance *MDM_Policy_Config01_WindowsDefenderSecurityCenter02) SetPropertyPhone(value string) (err error) {
	return instance.SetProperty("Phone", (value))
}

// GetPhone gets the value of Phone for the instance
func (instance *MDM_Policy_Config01_WindowsDefenderSecurityCenter02) GetPropertyPhone() (value string, err error) {
	retValue, err := instance.GetProperty("Phone")
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

// SetURL sets the value of URL for the instance
func (instance *MDM_Policy_Config01_WindowsDefenderSecurityCenter02) SetPropertyURL(value string) (err error) {
	return instance.SetProperty("URL", (value))
}

// GetURL gets the value of URL for the instance
func (instance *MDM_Policy_Config01_WindowsDefenderSecurityCenter02) GetPropertyURL() (value string, err error) {
	retValue, err := instance.GetProperty("URL")
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
