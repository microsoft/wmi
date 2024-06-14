// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NetFirewallHyperVProfile struct
type MSFT_NetFirewallHyperVProfile struct {
	*CIM_ManagedElement

	//
	AllowLocalFirewallRules uint16

	//
	DefaultInboundAction uint16

	//
	DefaultOutboundAction uint16

	//
	Enabled uint16

	//
	Name string

	//
	Profile uint16
}

func NewMSFT_NetFirewallHyperVProfileEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetFirewallHyperVProfile, err error) {
	tmp, err := NewCIM_ManagedElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetFirewallHyperVProfile{
		CIM_ManagedElement: tmp,
	}
	return
}

func NewMSFT_NetFirewallHyperVProfileEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetFirewallHyperVProfile, err error) {
	tmp, err := NewCIM_ManagedElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetFirewallHyperVProfile{
		CIM_ManagedElement: tmp,
	}
	return
}

// SetAllowLocalFirewallRules sets the value of AllowLocalFirewallRules for the instance
func (instance *MSFT_NetFirewallHyperVProfile) SetPropertyAllowLocalFirewallRules(value uint16) (err error) {
	return instance.SetProperty("AllowLocalFirewallRules", (value))
}

// GetAllowLocalFirewallRules gets the value of AllowLocalFirewallRules for the instance
func (instance *MSFT_NetFirewallHyperVProfile) GetPropertyAllowLocalFirewallRules() (value uint16, err error) {
	retValue, err := instance.GetProperty("AllowLocalFirewallRules")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetDefaultInboundAction sets the value of DefaultInboundAction for the instance
func (instance *MSFT_NetFirewallHyperVProfile) SetPropertyDefaultInboundAction(value uint16) (err error) {
	return instance.SetProperty("DefaultInboundAction", (value))
}

// GetDefaultInboundAction gets the value of DefaultInboundAction for the instance
func (instance *MSFT_NetFirewallHyperVProfile) GetPropertyDefaultInboundAction() (value uint16, err error) {
	retValue, err := instance.GetProperty("DefaultInboundAction")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetDefaultOutboundAction sets the value of DefaultOutboundAction for the instance
func (instance *MSFT_NetFirewallHyperVProfile) SetPropertyDefaultOutboundAction(value uint16) (err error) {
	return instance.SetProperty("DefaultOutboundAction", (value))
}

// GetDefaultOutboundAction gets the value of DefaultOutboundAction for the instance
func (instance *MSFT_NetFirewallHyperVProfile) GetPropertyDefaultOutboundAction() (value uint16, err error) {
	retValue, err := instance.GetProperty("DefaultOutboundAction")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetEnabled sets the value of Enabled for the instance
func (instance *MSFT_NetFirewallHyperVProfile) SetPropertyEnabled(value uint16) (err error) {
	return instance.SetProperty("Enabled", (value))
}

// GetEnabled gets the value of Enabled for the instance
func (instance *MSFT_NetFirewallHyperVProfile) GetPropertyEnabled() (value uint16, err error) {
	retValue, err := instance.GetProperty("Enabled")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetName sets the value of Name for the instance
func (instance *MSFT_NetFirewallHyperVProfile) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *MSFT_NetFirewallHyperVProfile) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
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

// SetProfile sets the value of Profile for the instance
func (instance *MSFT_NetFirewallHyperVProfile) SetPropertyProfile(value uint16) (err error) {
	return instance.SetProperty("Profile", (value))
}

// GetProfile gets the value of Profile for the instance
func (instance *MSFT_NetFirewallHyperVProfile) GetPropertyProfile() (value uint16, err error) {
	retValue, err := instance.GetProperty("Profile")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}
