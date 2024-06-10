// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.StandardCimv2
//
// ////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NetFirewallHyperVVMSetting struct
type MSFT_NetFirewallHyperVVMSetting struct {
	*CIM_ManagedElement

	//
	AllowHostPolicyMerge uint16

	//
	DefaultInboundAction uint16

	//
	DefaultOutboundAction uint16

	//
	Enabled uint16

	//
	LoopbackEnabled uint16

	//
	Name string
}

func NewMSFT_NetFirewallHyperVVMSettingEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetFirewallHyperVVMSetting, err error) {
	tmp, err := NewCIM_ManagedElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetFirewallHyperVVMSetting{
		CIM_ManagedElement: tmp,
	}
	return
}

func NewMSFT_NetFirewallHyperVVMSettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetFirewallHyperVVMSetting, err error) {
	tmp, err := NewCIM_ManagedElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetFirewallHyperVVMSetting{
		CIM_ManagedElement: tmp,
	}
	return
}

// SetAllowHostPolicyMerge sets the value of AllowHostPolicyMerge for the instance
func (instance *MSFT_NetFirewallHyperVVMSetting) SetPropertyAllowHostPolicyMerge(value uint16) (err error) {
	return instance.SetProperty("AllowHostPolicyMerge", (value))
}

// GetAllowHostPolicyMerge gets the value of AllowHostPolicyMerge for the instance
func (instance *MSFT_NetFirewallHyperVVMSetting) GetPropertyAllowHostPolicyMerge() (value uint16, err error) {
	retValue, err := instance.GetProperty("AllowHostPolicyMerge")
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
func (instance *MSFT_NetFirewallHyperVVMSetting) SetPropertyDefaultInboundAction(value uint16) (err error) {
	return instance.SetProperty("DefaultInboundAction", (value))
}

// GetDefaultInboundAction gets the value of DefaultInboundAction for the instance
func (instance *MSFT_NetFirewallHyperVVMSetting) GetPropertyDefaultInboundAction() (value uint16, err error) {
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
func (instance *MSFT_NetFirewallHyperVVMSetting) SetPropertyDefaultOutboundAction(value uint16) (err error) {
	return instance.SetProperty("DefaultOutboundAction", (value))
}

// GetDefaultOutboundAction gets the value of DefaultOutboundAction for the instance
func (instance *MSFT_NetFirewallHyperVVMSetting) GetPropertyDefaultOutboundAction() (value uint16, err error) {
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
func (instance *MSFT_NetFirewallHyperVVMSetting) SetPropertyEnabled(value uint16) (err error) {
	return instance.SetProperty("Enabled", (value))
}

// GetEnabled gets the value of Enabled for the instance
func (instance *MSFT_NetFirewallHyperVVMSetting) GetPropertyEnabled() (value uint16, err error) {
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

// SetLoopbackEnabled sets the value of LoopbackEnabled for the instance
func (instance *MSFT_NetFirewallHyperVVMSetting) SetPropertyLoopbackEnabled(value uint16) (err error) {
	return instance.SetProperty("LoopbackEnabled", (value))
}

// GetLoopbackEnabled gets the value of LoopbackEnabled for the instance
func (instance *MSFT_NetFirewallHyperVVMSetting) GetPropertyLoopbackEnabled() (value uint16, err error) {
	retValue, err := instance.GetProperty("LoopbackEnabled")
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
func (instance *MSFT_NetFirewallHyperVVMSetting) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *MSFT_NetFirewallHyperVVMSetting) GetPropertyName() (value string, err error) {
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
