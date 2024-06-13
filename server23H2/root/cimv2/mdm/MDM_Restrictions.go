// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.CIMV2.mdm
//////////////////////////////////////////////
package mdm

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MDM_Restrictions struct
type MDM_Restrictions struct {
	*cim.WmiInstance

	//
	BluetoothEnabled bool

	//
	DataRoamingEnabled bool

	//
	DiagnosticsSubmissionEnabled bool

	//
	EcsAutoProvisionEnabled bool

	//
	EcsSyncUrl string

	//
	IEEnterpriseModeEnabled bool

	//
	IEEnterpriseModeEnabledURL string

	//
	IEEnterpriseModeSitelist string

	//
	key uint32

	//
	PCSettingsMeteredNetworkSyncEnabled bool

	//
	PCSettingsPasswordSyncEnabled bool

	//
	PCSettingsSyncEnabled bool

	//
	SmartScreenEnabled bool

	//
	UserAccountControlStatus uint32

	//
	WifiEnabled bool
}

func NewMDM_RestrictionsEx1(instance *cim.WmiInstance) (newInstance *MDM_Restrictions, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Restrictions{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_RestrictionsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Restrictions, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Restrictions{
		WmiInstance: tmp,
	}
	return
}

// SetBluetoothEnabled sets the value of BluetoothEnabled for the instance
func (instance *MDM_Restrictions) SetPropertyBluetoothEnabled(value bool) (err error) {
	return instance.SetProperty("BluetoothEnabled", (value))
}

// GetBluetoothEnabled gets the value of BluetoothEnabled for the instance
func (instance *MDM_Restrictions) GetPropertyBluetoothEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("BluetoothEnabled")
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

// SetDataRoamingEnabled sets the value of DataRoamingEnabled for the instance
func (instance *MDM_Restrictions) SetPropertyDataRoamingEnabled(value bool) (err error) {
	return instance.SetProperty("DataRoamingEnabled", (value))
}

// GetDataRoamingEnabled gets the value of DataRoamingEnabled for the instance
func (instance *MDM_Restrictions) GetPropertyDataRoamingEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("DataRoamingEnabled")
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

// SetDiagnosticsSubmissionEnabled sets the value of DiagnosticsSubmissionEnabled for the instance
func (instance *MDM_Restrictions) SetPropertyDiagnosticsSubmissionEnabled(value bool) (err error) {
	return instance.SetProperty("DiagnosticsSubmissionEnabled", (value))
}

// GetDiagnosticsSubmissionEnabled gets the value of DiagnosticsSubmissionEnabled for the instance
func (instance *MDM_Restrictions) GetPropertyDiagnosticsSubmissionEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("DiagnosticsSubmissionEnabled")
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

// SetEcsAutoProvisionEnabled sets the value of EcsAutoProvisionEnabled for the instance
func (instance *MDM_Restrictions) SetPropertyEcsAutoProvisionEnabled(value bool) (err error) {
	return instance.SetProperty("EcsAutoProvisionEnabled", (value))
}

// GetEcsAutoProvisionEnabled gets the value of EcsAutoProvisionEnabled for the instance
func (instance *MDM_Restrictions) GetPropertyEcsAutoProvisionEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("EcsAutoProvisionEnabled")
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

// SetEcsSyncUrl sets the value of EcsSyncUrl for the instance
func (instance *MDM_Restrictions) SetPropertyEcsSyncUrl(value string) (err error) {
	return instance.SetProperty("EcsSyncUrl", (value))
}

// GetEcsSyncUrl gets the value of EcsSyncUrl for the instance
func (instance *MDM_Restrictions) GetPropertyEcsSyncUrl() (value string, err error) {
	retValue, err := instance.GetProperty("EcsSyncUrl")
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

// SetIEEnterpriseModeEnabled sets the value of IEEnterpriseModeEnabled for the instance
func (instance *MDM_Restrictions) SetPropertyIEEnterpriseModeEnabled(value bool) (err error) {
	return instance.SetProperty("IEEnterpriseModeEnabled", (value))
}

// GetIEEnterpriseModeEnabled gets the value of IEEnterpriseModeEnabled for the instance
func (instance *MDM_Restrictions) GetPropertyIEEnterpriseModeEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("IEEnterpriseModeEnabled")
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

// SetIEEnterpriseModeEnabledURL sets the value of IEEnterpriseModeEnabledURL for the instance
func (instance *MDM_Restrictions) SetPropertyIEEnterpriseModeEnabledURL(value string) (err error) {
	return instance.SetProperty("IEEnterpriseModeEnabledURL", (value))
}

// GetIEEnterpriseModeEnabledURL gets the value of IEEnterpriseModeEnabledURL for the instance
func (instance *MDM_Restrictions) GetPropertyIEEnterpriseModeEnabledURL() (value string, err error) {
	retValue, err := instance.GetProperty("IEEnterpriseModeEnabledURL")
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

// SetIEEnterpriseModeSitelist sets the value of IEEnterpriseModeSitelist for the instance
func (instance *MDM_Restrictions) SetPropertyIEEnterpriseModeSitelist(value string) (err error) {
	return instance.SetProperty("IEEnterpriseModeSitelist", (value))
}

// GetIEEnterpriseModeSitelist gets the value of IEEnterpriseModeSitelist for the instance
func (instance *MDM_Restrictions) GetPropertyIEEnterpriseModeSitelist() (value string, err error) {
	retValue, err := instance.GetProperty("IEEnterpriseModeSitelist")
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

// Setkey sets the value of key for the instance
func (instance *MDM_Restrictions) SetPropertykey(value uint32) (err error) {
	return instance.SetProperty("key", (value))
}

// Getkey gets the value of key for the instance
func (instance *MDM_Restrictions) GetPropertykey() (value uint32, err error) {
	retValue, err := instance.GetProperty("key")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetPCSettingsMeteredNetworkSyncEnabled sets the value of PCSettingsMeteredNetworkSyncEnabled for the instance
func (instance *MDM_Restrictions) SetPropertyPCSettingsMeteredNetworkSyncEnabled(value bool) (err error) {
	return instance.SetProperty("PCSettingsMeteredNetworkSyncEnabled", (value))
}

// GetPCSettingsMeteredNetworkSyncEnabled gets the value of PCSettingsMeteredNetworkSyncEnabled for the instance
func (instance *MDM_Restrictions) GetPropertyPCSettingsMeteredNetworkSyncEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("PCSettingsMeteredNetworkSyncEnabled")
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

// SetPCSettingsPasswordSyncEnabled sets the value of PCSettingsPasswordSyncEnabled for the instance
func (instance *MDM_Restrictions) SetPropertyPCSettingsPasswordSyncEnabled(value bool) (err error) {
	return instance.SetProperty("PCSettingsPasswordSyncEnabled", (value))
}

// GetPCSettingsPasswordSyncEnabled gets the value of PCSettingsPasswordSyncEnabled for the instance
func (instance *MDM_Restrictions) GetPropertyPCSettingsPasswordSyncEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("PCSettingsPasswordSyncEnabled")
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

// SetPCSettingsSyncEnabled sets the value of PCSettingsSyncEnabled for the instance
func (instance *MDM_Restrictions) SetPropertyPCSettingsSyncEnabled(value bool) (err error) {
	return instance.SetProperty("PCSettingsSyncEnabled", (value))
}

// GetPCSettingsSyncEnabled gets the value of PCSettingsSyncEnabled for the instance
func (instance *MDM_Restrictions) GetPropertyPCSettingsSyncEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("PCSettingsSyncEnabled")
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

// SetSmartScreenEnabled sets the value of SmartScreenEnabled for the instance
func (instance *MDM_Restrictions) SetPropertySmartScreenEnabled(value bool) (err error) {
	return instance.SetProperty("SmartScreenEnabled", (value))
}

// GetSmartScreenEnabled gets the value of SmartScreenEnabled for the instance
func (instance *MDM_Restrictions) GetPropertySmartScreenEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("SmartScreenEnabled")
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

// SetUserAccountControlStatus sets the value of UserAccountControlStatus for the instance
func (instance *MDM_Restrictions) SetPropertyUserAccountControlStatus(value uint32) (err error) {
	return instance.SetProperty("UserAccountControlStatus", (value))
}

// GetUserAccountControlStatus gets the value of UserAccountControlStatus for the instance
func (instance *MDM_Restrictions) GetPropertyUserAccountControlStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("UserAccountControlStatus")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetWifiEnabled sets the value of WifiEnabled for the instance
func (instance *MDM_Restrictions) SetPropertyWifiEnabled(value bool) (err error) {
	return instance.SetProperty("WifiEnabled", (value))
}

// GetWifiEnabled gets the value of WifiEnabled for the instance
func (instance *MDM_Restrictions) GetPropertyWifiEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("WifiEnabled")
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
