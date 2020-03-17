// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.mdm
//////////////////////////////////////////////
package mdm

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MDM_SecurityStatus struct
type MDM_SecurityStatus struct {
	cim.WmiInstance

	//
	AntiVirusSignatureStatus uint32

	//
	AntiVirusStatus uint32

	//
	ApplicationContentUriRules string

	//
	AutoUpdateStatus uint32

	//
	FirewallStatus uint32

	//
	IsMicrosoftAccountOptional bool

	//
	key uint32

	//
	MaintenanceScheduleAllowWakeup bool

	//
	MaintenanceScheduleDelayPattern string

	//
	MaintenanceScheduleStartHour uint32

	//
	RequireEncryption bool
}

// SetAntiVirusSignatureStatus sets the value of AntiVirusSignatureStatus for the instance
func (instance *MDM_SecurityStatus) SetPropertyAntiVirusSignatureStatus(value uint32) (err error) {
	return instance.SetProperty("AntiVirusSignatureStatus", value)
}

// GetAntiVirusSignatureStatus gets the value of AntiVirusSignatureStatus for the instance
func (instance *MDM_SecurityStatus) GetPropertyAntiVirusSignatureStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("AntiVirusSignatureStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAntiVirusStatus sets the value of AntiVirusStatus for the instance
func (instance *MDM_SecurityStatus) SetPropertyAntiVirusStatus(value uint32) (err error) {
	return instance.SetProperty("AntiVirusStatus", value)
}

// GetAntiVirusStatus gets the value of AntiVirusStatus for the instance
func (instance *MDM_SecurityStatus) GetPropertyAntiVirusStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("AntiVirusStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetApplicationContentUriRules sets the value of ApplicationContentUriRules for the instance
func (instance *MDM_SecurityStatus) SetPropertyApplicationContentUriRules(value string) (err error) {
	return instance.SetProperty("ApplicationContentUriRules", value)
}

// GetApplicationContentUriRules gets the value of ApplicationContentUriRules for the instance
func (instance *MDM_SecurityStatus) GetPropertyApplicationContentUriRules() (value string, err error) {
	retValue, err := instance.GetProperty("ApplicationContentUriRules")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAutoUpdateStatus sets the value of AutoUpdateStatus for the instance
func (instance *MDM_SecurityStatus) SetPropertyAutoUpdateStatus(value uint32) (err error) {
	return instance.SetProperty("AutoUpdateStatus", value)
}

// GetAutoUpdateStatus gets the value of AutoUpdateStatus for the instance
func (instance *MDM_SecurityStatus) GetPropertyAutoUpdateStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("AutoUpdateStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFirewallStatus sets the value of FirewallStatus for the instance
func (instance *MDM_SecurityStatus) SetPropertyFirewallStatus(value uint32) (err error) {
	return instance.SetProperty("FirewallStatus", value)
}

// GetFirewallStatus gets the value of FirewallStatus for the instance
func (instance *MDM_SecurityStatus) GetPropertyFirewallStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("FirewallStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsMicrosoftAccountOptional sets the value of IsMicrosoftAccountOptional for the instance
func (instance *MDM_SecurityStatus) SetPropertyIsMicrosoftAccountOptional(value bool) (err error) {
	return instance.SetProperty("IsMicrosoftAccountOptional", value)
}

// GetIsMicrosoftAccountOptional gets the value of IsMicrosoftAccountOptional for the instance
func (instance *MDM_SecurityStatus) GetPropertyIsMicrosoftAccountOptional() (value bool, err error) {
	retValue, err := instance.GetProperty("IsMicrosoftAccountOptional")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// Setkey sets the value of key for the instance
func (instance *MDM_SecurityStatus) SetPropertykey(value uint32) (err error) {
	return instance.SetProperty("key", value)
}

// Getkey gets the value of key for the instance
func (instance *MDM_SecurityStatus) GetPropertykey() (value uint32, err error) {
	retValue, err := instance.GetProperty("key")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaintenanceScheduleAllowWakeup sets the value of MaintenanceScheduleAllowWakeup for the instance
func (instance *MDM_SecurityStatus) SetPropertyMaintenanceScheduleAllowWakeup(value bool) (err error) {
	return instance.SetProperty("MaintenanceScheduleAllowWakeup", value)
}

// GetMaintenanceScheduleAllowWakeup gets the value of MaintenanceScheduleAllowWakeup for the instance
func (instance *MDM_SecurityStatus) GetPropertyMaintenanceScheduleAllowWakeup() (value bool, err error) {
	retValue, err := instance.GetProperty("MaintenanceScheduleAllowWakeup")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaintenanceScheduleDelayPattern sets the value of MaintenanceScheduleDelayPattern for the instance
func (instance *MDM_SecurityStatus) SetPropertyMaintenanceScheduleDelayPattern(value string) (err error) {
	return instance.SetProperty("MaintenanceScheduleDelayPattern", value)
}

// GetMaintenanceScheduleDelayPattern gets the value of MaintenanceScheduleDelayPattern for the instance
func (instance *MDM_SecurityStatus) GetPropertyMaintenanceScheduleDelayPattern() (value string, err error) {
	retValue, err := instance.GetProperty("MaintenanceScheduleDelayPattern")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaintenanceScheduleStartHour sets the value of MaintenanceScheduleStartHour for the instance
func (instance *MDM_SecurityStatus) SetPropertyMaintenanceScheduleStartHour(value uint32) (err error) {
	return instance.SetProperty("MaintenanceScheduleStartHour", value)
}

// GetMaintenanceScheduleStartHour gets the value of MaintenanceScheduleStartHour for the instance
func (instance *MDM_SecurityStatus) GetPropertyMaintenanceScheduleStartHour() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaintenanceScheduleStartHour")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRequireEncryption sets the value of RequireEncryption for the instance
func (instance *MDM_SecurityStatus) SetPropertyRequireEncryption(value bool) (err error) {
	return instance.SetProperty("RequireEncryption", value)
}

// GetRequireEncryption gets the value of RequireEncryption for the instance
func (instance *MDM_SecurityStatus) GetPropertyRequireEncryption() (value bool, err error) {
	retValue, err := instance.GetProperty("RequireEncryption")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
