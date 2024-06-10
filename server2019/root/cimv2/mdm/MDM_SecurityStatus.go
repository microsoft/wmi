// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
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

// MDM_SecurityStatus struct
type MDM_SecurityStatus struct {
	*cim.WmiInstance

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

func NewMDM_SecurityStatusEx1(instance *cim.WmiInstance) (newInstance *MDM_SecurityStatus, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_SecurityStatus{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_SecurityStatusEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_SecurityStatus, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_SecurityStatus{
		WmiInstance: tmp,
	}
	return
}

// SetAntiVirusSignatureStatus sets the value of AntiVirusSignatureStatus for the instance
func (instance *MDM_SecurityStatus) SetPropertyAntiVirusSignatureStatus(value uint32) (err error) {
	return instance.SetProperty("AntiVirusSignatureStatus", (value))
}

// GetAntiVirusSignatureStatus gets the value of AntiVirusSignatureStatus for the instance
func (instance *MDM_SecurityStatus) GetPropertyAntiVirusSignatureStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("AntiVirusSignatureStatus")
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

// SetAntiVirusStatus sets the value of AntiVirusStatus for the instance
func (instance *MDM_SecurityStatus) SetPropertyAntiVirusStatus(value uint32) (err error) {
	return instance.SetProperty("AntiVirusStatus", (value))
}

// GetAntiVirusStatus gets the value of AntiVirusStatus for the instance
func (instance *MDM_SecurityStatus) GetPropertyAntiVirusStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("AntiVirusStatus")
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

// SetApplicationContentUriRules sets the value of ApplicationContentUriRules for the instance
func (instance *MDM_SecurityStatus) SetPropertyApplicationContentUriRules(value string) (err error) {
	return instance.SetProperty("ApplicationContentUriRules", (value))
}

// GetApplicationContentUriRules gets the value of ApplicationContentUriRules for the instance
func (instance *MDM_SecurityStatus) GetPropertyApplicationContentUriRules() (value string, err error) {
	retValue, err := instance.GetProperty("ApplicationContentUriRules")
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

// SetAutoUpdateStatus sets the value of AutoUpdateStatus for the instance
func (instance *MDM_SecurityStatus) SetPropertyAutoUpdateStatus(value uint32) (err error) {
	return instance.SetProperty("AutoUpdateStatus", (value))
}

// GetAutoUpdateStatus gets the value of AutoUpdateStatus for the instance
func (instance *MDM_SecurityStatus) GetPropertyAutoUpdateStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("AutoUpdateStatus")
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

// SetFirewallStatus sets the value of FirewallStatus for the instance
func (instance *MDM_SecurityStatus) SetPropertyFirewallStatus(value uint32) (err error) {
	return instance.SetProperty("FirewallStatus", (value))
}

// GetFirewallStatus gets the value of FirewallStatus for the instance
func (instance *MDM_SecurityStatus) GetPropertyFirewallStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("FirewallStatus")
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

// SetIsMicrosoftAccountOptional sets the value of IsMicrosoftAccountOptional for the instance
func (instance *MDM_SecurityStatus) SetPropertyIsMicrosoftAccountOptional(value bool) (err error) {
	return instance.SetProperty("IsMicrosoftAccountOptional", (value))
}

// GetIsMicrosoftAccountOptional gets the value of IsMicrosoftAccountOptional for the instance
func (instance *MDM_SecurityStatus) GetPropertyIsMicrosoftAccountOptional() (value bool, err error) {
	retValue, err := instance.GetProperty("IsMicrosoftAccountOptional")
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

// Setkey sets the value of key for the instance
func (instance *MDM_SecurityStatus) SetPropertykey(value uint32) (err error) {
	return instance.SetProperty("key", (value))
}

// Getkey gets the value of key for the instance
func (instance *MDM_SecurityStatus) GetPropertykey() (value uint32, err error) {
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

// SetMaintenanceScheduleAllowWakeup sets the value of MaintenanceScheduleAllowWakeup for the instance
func (instance *MDM_SecurityStatus) SetPropertyMaintenanceScheduleAllowWakeup(value bool) (err error) {
	return instance.SetProperty("MaintenanceScheduleAllowWakeup", (value))
}

// GetMaintenanceScheduleAllowWakeup gets the value of MaintenanceScheduleAllowWakeup for the instance
func (instance *MDM_SecurityStatus) GetPropertyMaintenanceScheduleAllowWakeup() (value bool, err error) {
	retValue, err := instance.GetProperty("MaintenanceScheduleAllowWakeup")
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

// SetMaintenanceScheduleDelayPattern sets the value of MaintenanceScheduleDelayPattern for the instance
func (instance *MDM_SecurityStatus) SetPropertyMaintenanceScheduleDelayPattern(value string) (err error) {
	return instance.SetProperty("MaintenanceScheduleDelayPattern", (value))
}

// GetMaintenanceScheduleDelayPattern gets the value of MaintenanceScheduleDelayPattern for the instance
func (instance *MDM_SecurityStatus) GetPropertyMaintenanceScheduleDelayPattern() (value string, err error) {
	retValue, err := instance.GetProperty("MaintenanceScheduleDelayPattern")
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

// SetMaintenanceScheduleStartHour sets the value of MaintenanceScheduleStartHour for the instance
func (instance *MDM_SecurityStatus) SetPropertyMaintenanceScheduleStartHour(value uint32) (err error) {
	return instance.SetProperty("MaintenanceScheduleStartHour", (value))
}

// GetMaintenanceScheduleStartHour gets the value of MaintenanceScheduleStartHour for the instance
func (instance *MDM_SecurityStatus) GetPropertyMaintenanceScheduleStartHour() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaintenanceScheduleStartHour")
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

// SetRequireEncryption sets the value of RequireEncryption for the instance
func (instance *MDM_SecurityStatus) SetPropertyRequireEncryption(value bool) (err error) {
	return instance.SetProperty("RequireEncryption", (value))
}

// GetRequireEncryption gets the value of RequireEncryption for the instance
func (instance *MDM_SecurityStatus) GetPropertyRequireEncryption() (value bool, err error) {
	retValue, err := instance.GetProperty("RequireEncryption")
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
