// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.virtualization.v2
//
// ////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_SecuritySettingData struct
type Msvm_SecuritySettingData struct {
	*CIM_SettingData

	//
	AppContainerLaunchOptOut bool

	//
	BindToHostTpm bool

	//
	DataProtectionRequested bool

	//
	EncryptStateAndVmMigrationTraffic bool

	//
	KsdEnabled bool

	//
	ShieldingRequested bool

	//
	TpmEnabled bool

	//
	VirtualizationBasedSecurityOptOut bool
}

func NewMsvm_SecuritySettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_SecuritySettingData, err error) {
	tmp, err := NewCIM_SettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_SecuritySettingData{
		CIM_SettingData: tmp,
	}
	return
}

func NewMsvm_SecuritySettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_SecuritySettingData, err error) {
	tmp, err := NewCIM_SettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_SecuritySettingData{
		CIM_SettingData: tmp,
	}
	return
}

// SetAppContainerLaunchOptOut sets the value of AppContainerLaunchOptOut for the instance
func (instance *Msvm_SecuritySettingData) SetPropertyAppContainerLaunchOptOut(value bool) (err error) {
	return instance.SetProperty("AppContainerLaunchOptOut", (value))
}

// GetAppContainerLaunchOptOut gets the value of AppContainerLaunchOptOut for the instance
func (instance *Msvm_SecuritySettingData) GetPropertyAppContainerLaunchOptOut() (value bool, err error) {
	retValue, err := instance.GetProperty("AppContainerLaunchOptOut")
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

// SetBindToHostTpm sets the value of BindToHostTpm for the instance
func (instance *Msvm_SecuritySettingData) SetPropertyBindToHostTpm(value bool) (err error) {
	return instance.SetProperty("BindToHostTpm", (value))
}

// GetBindToHostTpm gets the value of BindToHostTpm for the instance
func (instance *Msvm_SecuritySettingData) GetPropertyBindToHostTpm() (value bool, err error) {
	retValue, err := instance.GetProperty("BindToHostTpm")
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

// SetDataProtectionRequested sets the value of DataProtectionRequested for the instance
func (instance *Msvm_SecuritySettingData) SetPropertyDataProtectionRequested(value bool) (err error) {
	return instance.SetProperty("DataProtectionRequested", (value))
}

// GetDataProtectionRequested gets the value of DataProtectionRequested for the instance
func (instance *Msvm_SecuritySettingData) GetPropertyDataProtectionRequested() (value bool, err error) {
	retValue, err := instance.GetProperty("DataProtectionRequested")
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

// SetEncryptStateAndVmMigrationTraffic sets the value of EncryptStateAndVmMigrationTraffic for the instance
func (instance *Msvm_SecuritySettingData) SetPropertyEncryptStateAndVmMigrationTraffic(value bool) (err error) {
	return instance.SetProperty("EncryptStateAndVmMigrationTraffic", (value))
}

// GetEncryptStateAndVmMigrationTraffic gets the value of EncryptStateAndVmMigrationTraffic for the instance
func (instance *Msvm_SecuritySettingData) GetPropertyEncryptStateAndVmMigrationTraffic() (value bool, err error) {
	retValue, err := instance.GetProperty("EncryptStateAndVmMigrationTraffic")
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

// SetKsdEnabled sets the value of KsdEnabled for the instance
func (instance *Msvm_SecuritySettingData) SetPropertyKsdEnabled(value bool) (err error) {
	return instance.SetProperty("KsdEnabled", (value))
}

// GetKsdEnabled gets the value of KsdEnabled for the instance
func (instance *Msvm_SecuritySettingData) GetPropertyKsdEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("KsdEnabled")
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

// SetShieldingRequested sets the value of ShieldingRequested for the instance
func (instance *Msvm_SecuritySettingData) SetPropertyShieldingRequested(value bool) (err error) {
	return instance.SetProperty("ShieldingRequested", (value))
}

// GetShieldingRequested gets the value of ShieldingRequested for the instance
func (instance *Msvm_SecuritySettingData) GetPropertyShieldingRequested() (value bool, err error) {
	retValue, err := instance.GetProperty("ShieldingRequested")
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

// SetTpmEnabled sets the value of TpmEnabled for the instance
func (instance *Msvm_SecuritySettingData) SetPropertyTpmEnabled(value bool) (err error) {
	return instance.SetProperty("TpmEnabled", (value))
}

// GetTpmEnabled gets the value of TpmEnabled for the instance
func (instance *Msvm_SecuritySettingData) GetPropertyTpmEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("TpmEnabled")
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

// SetVirtualizationBasedSecurityOptOut sets the value of VirtualizationBasedSecurityOptOut for the instance
func (instance *Msvm_SecuritySettingData) SetPropertyVirtualizationBasedSecurityOptOut(value bool) (err error) {
	return instance.SetProperty("VirtualizationBasedSecurityOptOut", (value))
}

// GetVirtualizationBasedSecurityOptOut gets the value of VirtualizationBasedSecurityOptOut for the instance
func (instance *Msvm_SecuritySettingData) GetPropertyVirtualizationBasedSecurityOptOut() (value bool, err error) {
	retValue, err := instance.GetProperty("VirtualizationBasedSecurityOptOut")
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
