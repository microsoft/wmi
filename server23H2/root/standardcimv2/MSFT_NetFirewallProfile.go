// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NetFirewallProfile struct
type MSFT_NetFirewallProfile struct {
	*CIM_ManagedElement

	//
	AllowInboundRules uint16

	//
	AllowLocalFirewallRules uint16

	//
	AllowLocalIPsecRules uint16

	//
	AllowUnicastResponseToMulticast uint16

	//
	AllowUserApps uint16

	//
	AllowUserPorts uint16

	//
	DefaultInboundAction uint16

	//
	DefaultOutboundAction uint16

	//
	DisabledInterfaceAliases []string

	//
	Enabled uint16

	//
	EnableStealthModeForIPsec uint16

	//
	LogAllowed uint16

	//
	LogBlocked uint16

	//
	LogFileName string

	//
	LogIgnored uint16

	//
	LogMaxSizeKilobytes uint64

	//
	Name string

	//
	NotifyOnListen uint16
}

func NewMSFT_NetFirewallProfileEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetFirewallProfile, err error) {
	tmp, err := NewCIM_ManagedElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetFirewallProfile{
		CIM_ManagedElement: tmp,
	}
	return
}

func NewMSFT_NetFirewallProfileEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetFirewallProfile, err error) {
	tmp, err := NewCIM_ManagedElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetFirewallProfile{
		CIM_ManagedElement: tmp,
	}
	return
}

// SetAllowInboundRules sets the value of AllowInboundRules for the instance
func (instance *MSFT_NetFirewallProfile) SetPropertyAllowInboundRules(value uint16) (err error) {
	return instance.SetProperty("AllowInboundRules", (value))
}

// GetAllowInboundRules gets the value of AllowInboundRules for the instance
func (instance *MSFT_NetFirewallProfile) GetPropertyAllowInboundRules() (value uint16, err error) {
	retValue, err := instance.GetProperty("AllowInboundRules")
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

// SetAllowLocalFirewallRules sets the value of AllowLocalFirewallRules for the instance
func (instance *MSFT_NetFirewallProfile) SetPropertyAllowLocalFirewallRules(value uint16) (err error) {
	return instance.SetProperty("AllowLocalFirewallRules", (value))
}

// GetAllowLocalFirewallRules gets the value of AllowLocalFirewallRules for the instance
func (instance *MSFT_NetFirewallProfile) GetPropertyAllowLocalFirewallRules() (value uint16, err error) {
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

// SetAllowLocalIPsecRules sets the value of AllowLocalIPsecRules for the instance
func (instance *MSFT_NetFirewallProfile) SetPropertyAllowLocalIPsecRules(value uint16) (err error) {
	return instance.SetProperty("AllowLocalIPsecRules", (value))
}

// GetAllowLocalIPsecRules gets the value of AllowLocalIPsecRules for the instance
func (instance *MSFT_NetFirewallProfile) GetPropertyAllowLocalIPsecRules() (value uint16, err error) {
	retValue, err := instance.GetProperty("AllowLocalIPsecRules")
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

// SetAllowUnicastResponseToMulticast sets the value of AllowUnicastResponseToMulticast for the instance
func (instance *MSFT_NetFirewallProfile) SetPropertyAllowUnicastResponseToMulticast(value uint16) (err error) {
	return instance.SetProperty("AllowUnicastResponseToMulticast", (value))
}

// GetAllowUnicastResponseToMulticast gets the value of AllowUnicastResponseToMulticast for the instance
func (instance *MSFT_NetFirewallProfile) GetPropertyAllowUnicastResponseToMulticast() (value uint16, err error) {
	retValue, err := instance.GetProperty("AllowUnicastResponseToMulticast")
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

// SetAllowUserApps sets the value of AllowUserApps for the instance
func (instance *MSFT_NetFirewallProfile) SetPropertyAllowUserApps(value uint16) (err error) {
	return instance.SetProperty("AllowUserApps", (value))
}

// GetAllowUserApps gets the value of AllowUserApps for the instance
func (instance *MSFT_NetFirewallProfile) GetPropertyAllowUserApps() (value uint16, err error) {
	retValue, err := instance.GetProperty("AllowUserApps")
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

// SetAllowUserPorts sets the value of AllowUserPorts for the instance
func (instance *MSFT_NetFirewallProfile) SetPropertyAllowUserPorts(value uint16) (err error) {
	return instance.SetProperty("AllowUserPorts", (value))
}

// GetAllowUserPorts gets the value of AllowUserPorts for the instance
func (instance *MSFT_NetFirewallProfile) GetPropertyAllowUserPorts() (value uint16, err error) {
	retValue, err := instance.GetProperty("AllowUserPorts")
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
func (instance *MSFT_NetFirewallProfile) SetPropertyDefaultInboundAction(value uint16) (err error) {
	return instance.SetProperty("DefaultInboundAction", (value))
}

// GetDefaultInboundAction gets the value of DefaultInboundAction for the instance
func (instance *MSFT_NetFirewallProfile) GetPropertyDefaultInboundAction() (value uint16, err error) {
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
func (instance *MSFT_NetFirewallProfile) SetPropertyDefaultOutboundAction(value uint16) (err error) {
	return instance.SetProperty("DefaultOutboundAction", (value))
}

// GetDefaultOutboundAction gets the value of DefaultOutboundAction for the instance
func (instance *MSFT_NetFirewallProfile) GetPropertyDefaultOutboundAction() (value uint16, err error) {
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

// SetDisabledInterfaceAliases sets the value of DisabledInterfaceAliases for the instance
func (instance *MSFT_NetFirewallProfile) SetPropertyDisabledInterfaceAliases(value []string) (err error) {
	return instance.SetProperty("DisabledInterfaceAliases", (value))
}

// GetDisabledInterfaceAliases gets the value of DisabledInterfaceAliases for the instance
func (instance *MSFT_NetFirewallProfile) GetPropertyDisabledInterfaceAliases() (value []string, err error) {
	retValue, err := instance.GetProperty("DisabledInterfaceAliases")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetEnabled sets the value of Enabled for the instance
func (instance *MSFT_NetFirewallProfile) SetPropertyEnabled(value uint16) (err error) {
	return instance.SetProperty("Enabled", (value))
}

// GetEnabled gets the value of Enabled for the instance
func (instance *MSFT_NetFirewallProfile) GetPropertyEnabled() (value uint16, err error) {
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

// SetEnableStealthModeForIPsec sets the value of EnableStealthModeForIPsec for the instance
func (instance *MSFT_NetFirewallProfile) SetPropertyEnableStealthModeForIPsec(value uint16) (err error) {
	return instance.SetProperty("EnableStealthModeForIPsec", (value))
}

// GetEnableStealthModeForIPsec gets the value of EnableStealthModeForIPsec for the instance
func (instance *MSFT_NetFirewallProfile) GetPropertyEnableStealthModeForIPsec() (value uint16, err error) {
	retValue, err := instance.GetProperty("EnableStealthModeForIPsec")
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

// SetLogAllowed sets the value of LogAllowed for the instance
func (instance *MSFT_NetFirewallProfile) SetPropertyLogAllowed(value uint16) (err error) {
	return instance.SetProperty("LogAllowed", (value))
}

// GetLogAllowed gets the value of LogAllowed for the instance
func (instance *MSFT_NetFirewallProfile) GetPropertyLogAllowed() (value uint16, err error) {
	retValue, err := instance.GetProperty("LogAllowed")
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

// SetLogBlocked sets the value of LogBlocked for the instance
func (instance *MSFT_NetFirewallProfile) SetPropertyLogBlocked(value uint16) (err error) {
	return instance.SetProperty("LogBlocked", (value))
}

// GetLogBlocked gets the value of LogBlocked for the instance
func (instance *MSFT_NetFirewallProfile) GetPropertyLogBlocked() (value uint16, err error) {
	retValue, err := instance.GetProperty("LogBlocked")
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

// SetLogFileName sets the value of LogFileName for the instance
func (instance *MSFT_NetFirewallProfile) SetPropertyLogFileName(value string) (err error) {
	return instance.SetProperty("LogFileName", (value))
}

// GetLogFileName gets the value of LogFileName for the instance
func (instance *MSFT_NetFirewallProfile) GetPropertyLogFileName() (value string, err error) {
	retValue, err := instance.GetProperty("LogFileName")
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

// SetLogIgnored sets the value of LogIgnored for the instance
func (instance *MSFT_NetFirewallProfile) SetPropertyLogIgnored(value uint16) (err error) {
	return instance.SetProperty("LogIgnored", (value))
}

// GetLogIgnored gets the value of LogIgnored for the instance
func (instance *MSFT_NetFirewallProfile) GetPropertyLogIgnored() (value uint16, err error) {
	retValue, err := instance.GetProperty("LogIgnored")
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

// SetLogMaxSizeKilobytes sets the value of LogMaxSizeKilobytes for the instance
func (instance *MSFT_NetFirewallProfile) SetPropertyLogMaxSizeKilobytes(value uint64) (err error) {
	return instance.SetProperty("LogMaxSizeKilobytes", (value))
}

// GetLogMaxSizeKilobytes gets the value of LogMaxSizeKilobytes for the instance
func (instance *MSFT_NetFirewallProfile) GetPropertyLogMaxSizeKilobytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("LogMaxSizeKilobytes")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetName sets the value of Name for the instance
func (instance *MSFT_NetFirewallProfile) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *MSFT_NetFirewallProfile) GetPropertyName() (value string, err error) {
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

// SetNotifyOnListen sets the value of NotifyOnListen for the instance
func (instance *MSFT_NetFirewallProfile) SetPropertyNotifyOnListen(value uint16) (err error) {
	return instance.SetProperty("NotifyOnListen", (value))
}

// GetNotifyOnListen gets the value of NotifyOnListen for the instance
func (instance *MSFT_NetFirewallProfile) GetPropertyNotifyOnListen() (value uint16, err error) {
	retValue, err := instance.GetProperty("NotifyOnListen")
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
