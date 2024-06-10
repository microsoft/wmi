// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WebAdministration
//
// ////////////////////////////////////////////
package webadministration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// SmtpMailSettingsSection struct
type SmtpMailSettingsSection struct {
	*ConfigurationSection

	//
	DeliveryMethod int32

	//
	From string

	//
	Network SmtpNetworkSettings

	//
	SpecifiedPickupDirectory SmtpSpecifiedPickupDirectorySettings
}

func NewSmtpMailSettingsSectionEx1(instance *cim.WmiInstance) (newInstance *SmtpMailSettingsSection, err error) {
	tmp, err := NewConfigurationSectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SmtpMailSettingsSection{
		ConfigurationSection: tmp,
	}
	return
}

func NewSmtpMailSettingsSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SmtpMailSettingsSection, err error) {
	tmp, err := NewConfigurationSectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SmtpMailSettingsSection{
		ConfigurationSection: tmp,
	}
	return
}

// SetDeliveryMethod sets the value of DeliveryMethod for the instance
func (instance *SmtpMailSettingsSection) SetPropertyDeliveryMethod(value int32) (err error) {
	return instance.SetProperty("DeliveryMethod", (value))
}

// GetDeliveryMethod gets the value of DeliveryMethod for the instance
func (instance *SmtpMailSettingsSection) GetPropertyDeliveryMethod() (value int32, err error) {
	retValue, err := instance.GetProperty("DeliveryMethod")
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

// SetFrom sets the value of From for the instance
func (instance *SmtpMailSettingsSection) SetPropertyFrom(value string) (err error) {
	return instance.SetProperty("From", (value))
}

// GetFrom gets the value of From for the instance
func (instance *SmtpMailSettingsSection) GetPropertyFrom() (value string, err error) {
	retValue, err := instance.GetProperty("From")
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

// SetNetwork sets the value of Network for the instance
func (instance *SmtpMailSettingsSection) SetPropertyNetwork(value SmtpNetworkSettings) (err error) {
	return instance.SetProperty("Network", (value))
}

// GetNetwork gets the value of Network for the instance
func (instance *SmtpMailSettingsSection) GetPropertyNetwork() (value SmtpNetworkSettings, err error) {
	retValue, err := instance.GetProperty("Network")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(SmtpNetworkSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " SmtpNetworkSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = SmtpNetworkSettings(valuetmp)

	return
}

// SetSpecifiedPickupDirectory sets the value of SpecifiedPickupDirectory for the instance
func (instance *SmtpMailSettingsSection) SetPropertySpecifiedPickupDirectory(value SmtpSpecifiedPickupDirectorySettings) (err error) {
	return instance.SetProperty("SpecifiedPickupDirectory", (value))
}

// GetSpecifiedPickupDirectory gets the value of SpecifiedPickupDirectory for the instance
func (instance *SmtpMailSettingsSection) GetPropertySpecifiedPickupDirectory() (value SmtpSpecifiedPickupDirectorySettings, err error) {
	retValue, err := instance.GetProperty("SpecifiedPickupDirectory")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(SmtpSpecifiedPickupDirectorySettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " SmtpSpecifiedPickupDirectorySettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = SmtpSpecifiedPickupDirectorySettings(valuetmp)

	return
}
