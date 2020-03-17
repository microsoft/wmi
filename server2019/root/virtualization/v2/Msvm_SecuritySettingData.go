// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

// Msvm_SecuritySettingData struct
type Msvm_SecuritySettingData struct {
	CIM_SettingData

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

// SetAppContainerLaunchOptOut sets the value of AppContainerLaunchOptOut for the instance
func (instance *Msvm_SecuritySettingData) SetPropertyAppContainerLaunchOptOut(value bool) (err error) {
	return instance.SetProperty("AppContainerLaunchOptOut", value)
}

// GetAppContainerLaunchOptOut gets the value of AppContainerLaunchOptOut for the instance
func (instance *Msvm_SecuritySettingData) GetPropertyAppContainerLaunchOptOut() (value bool, err error) {
	retValue, err := instance.GetProperty("AppContainerLaunchOptOut")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBindToHostTpm sets the value of BindToHostTpm for the instance
func (instance *Msvm_SecuritySettingData) SetPropertyBindToHostTpm(value bool) (err error) {
	return instance.SetProperty("BindToHostTpm", value)
}

// GetBindToHostTpm gets the value of BindToHostTpm for the instance
func (instance *Msvm_SecuritySettingData) GetPropertyBindToHostTpm() (value bool, err error) {
	retValue, err := instance.GetProperty("BindToHostTpm")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDataProtectionRequested sets the value of DataProtectionRequested for the instance
func (instance *Msvm_SecuritySettingData) SetPropertyDataProtectionRequested(value bool) (err error) {
	return instance.SetProperty("DataProtectionRequested", value)
}

// GetDataProtectionRequested gets the value of DataProtectionRequested for the instance
func (instance *Msvm_SecuritySettingData) GetPropertyDataProtectionRequested() (value bool, err error) {
	retValue, err := instance.GetProperty("DataProtectionRequested")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEncryptStateAndVmMigrationTraffic sets the value of EncryptStateAndVmMigrationTraffic for the instance
func (instance *Msvm_SecuritySettingData) SetPropertyEncryptStateAndVmMigrationTraffic(value bool) (err error) {
	return instance.SetProperty("EncryptStateAndVmMigrationTraffic", value)
}

// GetEncryptStateAndVmMigrationTraffic gets the value of EncryptStateAndVmMigrationTraffic for the instance
func (instance *Msvm_SecuritySettingData) GetPropertyEncryptStateAndVmMigrationTraffic() (value bool, err error) {
	retValue, err := instance.GetProperty("EncryptStateAndVmMigrationTraffic")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetKsdEnabled sets the value of KsdEnabled for the instance
func (instance *Msvm_SecuritySettingData) SetPropertyKsdEnabled(value bool) (err error) {
	return instance.SetProperty("KsdEnabled", value)
}

// GetKsdEnabled gets the value of KsdEnabled for the instance
func (instance *Msvm_SecuritySettingData) GetPropertyKsdEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("KsdEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetShieldingRequested sets the value of ShieldingRequested for the instance
func (instance *Msvm_SecuritySettingData) SetPropertyShieldingRequested(value bool) (err error) {
	return instance.SetProperty("ShieldingRequested", value)
}

// GetShieldingRequested gets the value of ShieldingRequested for the instance
func (instance *Msvm_SecuritySettingData) GetPropertyShieldingRequested() (value bool, err error) {
	retValue, err := instance.GetProperty("ShieldingRequested")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTpmEnabled sets the value of TpmEnabled for the instance
func (instance *Msvm_SecuritySettingData) SetPropertyTpmEnabled(value bool) (err error) {
	return instance.SetProperty("TpmEnabled", value)
}

// GetTpmEnabled gets the value of TpmEnabled for the instance
func (instance *Msvm_SecuritySettingData) GetPropertyTpmEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("TpmEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVirtualizationBasedSecurityOptOut sets the value of VirtualizationBasedSecurityOptOut for the instance
func (instance *Msvm_SecuritySettingData) SetPropertyVirtualizationBasedSecurityOptOut(value bool) (err error) {
	return instance.SetProperty("VirtualizationBasedSecurityOptOut", value)
}

// GetVirtualizationBasedSecurityOptOut gets the value of VirtualizationBasedSecurityOptOut for the instance
func (instance *Msvm_SecuritySettingData) GetPropertyVirtualizationBasedSecurityOptOut() (value bool, err error) {
	retValue, err := instance.GetProperty("VirtualizationBasedSecurityOptOut")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
