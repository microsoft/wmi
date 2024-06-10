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

// MSFT_NetIPsecDoSPSetting struct
type MSFT_NetIPsecDoSPSetting struct {
	*MSFT_NetSettingData

	//
	DefBlockExemptDscp uint16

	//
	DefBlockExemptRateLimitBytesPerSec uint32

	//
	EffectiveAddressFamily uint16

	//
	EnabledKeyingModules uint32

	//
	FilteringFlags uint32

	//
	IcmpV6Dscp uint16

	//
	IcmpV6RateLimitBytesPerSec uint32

	//
	IpV6FilterExemptDscp uint32

	//
	IpV6FilterExemptRateLimitBytesPerSec uint32

	//
	IpV6IPsecAuthDscp uint16

	//
	IpV6IPsecAuthRateLimitBytesPerSec uint32

	//
	IpV6IPsecUnauthDscp uint32

	//
	IpV6IPsecUnauthPerIPRateLimitBytesPerSec uint32

	//
	IpV6IPsecUnauthRateLimitBytesPerSec uint32

	//
	MaxPerIPRateLimitQueues uint32

	//
	MaxStateEntries uint32

	//
	PerIPRateLimitQueueIdleTimeoutSeconds uint32

	//
	PrivateInterfaceAliases []string

	//
	PrivateV6Address string

	//
	PublicInterfaceAliases []string

	//
	PublicV6Address string

	//
	StateIdleTimeoutSeconds uint32
}

func NewMSFT_NetIPsecDoSPSettingEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetIPsecDoSPSetting, err error) {
	tmp, err := NewMSFT_NetSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetIPsecDoSPSetting{
		MSFT_NetSettingData: tmp,
	}
	return
}

func NewMSFT_NetIPsecDoSPSettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetIPsecDoSPSetting, err error) {
	tmp, err := NewMSFT_NetSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetIPsecDoSPSetting{
		MSFT_NetSettingData: tmp,
	}
	return
}

// SetDefBlockExemptDscp sets the value of DefBlockExemptDscp for the instance
func (instance *MSFT_NetIPsecDoSPSetting) SetPropertyDefBlockExemptDscp(value uint16) (err error) {
	return instance.SetProperty("DefBlockExemptDscp", (value))
}

// GetDefBlockExemptDscp gets the value of DefBlockExemptDscp for the instance
func (instance *MSFT_NetIPsecDoSPSetting) GetPropertyDefBlockExemptDscp() (value uint16, err error) {
	retValue, err := instance.GetProperty("DefBlockExemptDscp")
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

// SetDefBlockExemptRateLimitBytesPerSec sets the value of DefBlockExemptRateLimitBytesPerSec for the instance
func (instance *MSFT_NetIPsecDoSPSetting) SetPropertyDefBlockExemptRateLimitBytesPerSec(value uint32) (err error) {
	return instance.SetProperty("DefBlockExemptRateLimitBytesPerSec", (value))
}

// GetDefBlockExemptRateLimitBytesPerSec gets the value of DefBlockExemptRateLimitBytesPerSec for the instance
func (instance *MSFT_NetIPsecDoSPSetting) GetPropertyDefBlockExemptRateLimitBytesPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("DefBlockExemptRateLimitBytesPerSec")
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

// SetEffectiveAddressFamily sets the value of EffectiveAddressFamily for the instance
func (instance *MSFT_NetIPsecDoSPSetting) SetPropertyEffectiveAddressFamily(value uint16) (err error) {
	return instance.SetProperty("EffectiveAddressFamily", (value))
}

// GetEffectiveAddressFamily gets the value of EffectiveAddressFamily for the instance
func (instance *MSFT_NetIPsecDoSPSetting) GetPropertyEffectiveAddressFamily() (value uint16, err error) {
	retValue, err := instance.GetProperty("EffectiveAddressFamily")
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

// SetEnabledKeyingModules sets the value of EnabledKeyingModules for the instance
func (instance *MSFT_NetIPsecDoSPSetting) SetPropertyEnabledKeyingModules(value uint32) (err error) {
	return instance.SetProperty("EnabledKeyingModules", (value))
}

// GetEnabledKeyingModules gets the value of EnabledKeyingModules for the instance
func (instance *MSFT_NetIPsecDoSPSetting) GetPropertyEnabledKeyingModules() (value uint32, err error) {
	retValue, err := instance.GetProperty("EnabledKeyingModules")
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

// SetFilteringFlags sets the value of FilteringFlags for the instance
func (instance *MSFT_NetIPsecDoSPSetting) SetPropertyFilteringFlags(value uint32) (err error) {
	return instance.SetProperty("FilteringFlags", (value))
}

// GetFilteringFlags gets the value of FilteringFlags for the instance
func (instance *MSFT_NetIPsecDoSPSetting) GetPropertyFilteringFlags() (value uint32, err error) {
	retValue, err := instance.GetProperty("FilteringFlags")
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

// SetIcmpV6Dscp sets the value of IcmpV6Dscp for the instance
func (instance *MSFT_NetIPsecDoSPSetting) SetPropertyIcmpV6Dscp(value uint16) (err error) {
	return instance.SetProperty("IcmpV6Dscp", (value))
}

// GetIcmpV6Dscp gets the value of IcmpV6Dscp for the instance
func (instance *MSFT_NetIPsecDoSPSetting) GetPropertyIcmpV6Dscp() (value uint16, err error) {
	retValue, err := instance.GetProperty("IcmpV6Dscp")
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

// SetIcmpV6RateLimitBytesPerSec sets the value of IcmpV6RateLimitBytesPerSec for the instance
func (instance *MSFT_NetIPsecDoSPSetting) SetPropertyIcmpV6RateLimitBytesPerSec(value uint32) (err error) {
	return instance.SetProperty("IcmpV6RateLimitBytesPerSec", (value))
}

// GetIcmpV6RateLimitBytesPerSec gets the value of IcmpV6RateLimitBytesPerSec for the instance
func (instance *MSFT_NetIPsecDoSPSetting) GetPropertyIcmpV6RateLimitBytesPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("IcmpV6RateLimitBytesPerSec")
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

// SetIpV6FilterExemptDscp sets the value of IpV6FilterExemptDscp for the instance
func (instance *MSFT_NetIPsecDoSPSetting) SetPropertyIpV6FilterExemptDscp(value uint32) (err error) {
	return instance.SetProperty("IpV6FilterExemptDscp", (value))
}

// GetIpV6FilterExemptDscp gets the value of IpV6FilterExemptDscp for the instance
func (instance *MSFT_NetIPsecDoSPSetting) GetPropertyIpV6FilterExemptDscp() (value uint32, err error) {
	retValue, err := instance.GetProperty("IpV6FilterExemptDscp")
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

// SetIpV6FilterExemptRateLimitBytesPerSec sets the value of IpV6FilterExemptRateLimitBytesPerSec for the instance
func (instance *MSFT_NetIPsecDoSPSetting) SetPropertyIpV6FilterExemptRateLimitBytesPerSec(value uint32) (err error) {
	return instance.SetProperty("IpV6FilterExemptRateLimitBytesPerSec", (value))
}

// GetIpV6FilterExemptRateLimitBytesPerSec gets the value of IpV6FilterExemptRateLimitBytesPerSec for the instance
func (instance *MSFT_NetIPsecDoSPSetting) GetPropertyIpV6FilterExemptRateLimitBytesPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("IpV6FilterExemptRateLimitBytesPerSec")
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

// SetIpV6IPsecAuthDscp sets the value of IpV6IPsecAuthDscp for the instance
func (instance *MSFT_NetIPsecDoSPSetting) SetPropertyIpV6IPsecAuthDscp(value uint16) (err error) {
	return instance.SetProperty("IpV6IPsecAuthDscp", (value))
}

// GetIpV6IPsecAuthDscp gets the value of IpV6IPsecAuthDscp for the instance
func (instance *MSFT_NetIPsecDoSPSetting) GetPropertyIpV6IPsecAuthDscp() (value uint16, err error) {
	retValue, err := instance.GetProperty("IpV6IPsecAuthDscp")
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

// SetIpV6IPsecAuthRateLimitBytesPerSec sets the value of IpV6IPsecAuthRateLimitBytesPerSec for the instance
func (instance *MSFT_NetIPsecDoSPSetting) SetPropertyIpV6IPsecAuthRateLimitBytesPerSec(value uint32) (err error) {
	return instance.SetProperty("IpV6IPsecAuthRateLimitBytesPerSec", (value))
}

// GetIpV6IPsecAuthRateLimitBytesPerSec gets the value of IpV6IPsecAuthRateLimitBytesPerSec for the instance
func (instance *MSFT_NetIPsecDoSPSetting) GetPropertyIpV6IPsecAuthRateLimitBytesPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("IpV6IPsecAuthRateLimitBytesPerSec")
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

// SetIpV6IPsecUnauthDscp sets the value of IpV6IPsecUnauthDscp for the instance
func (instance *MSFT_NetIPsecDoSPSetting) SetPropertyIpV6IPsecUnauthDscp(value uint32) (err error) {
	return instance.SetProperty("IpV6IPsecUnauthDscp", (value))
}

// GetIpV6IPsecUnauthDscp gets the value of IpV6IPsecUnauthDscp for the instance
func (instance *MSFT_NetIPsecDoSPSetting) GetPropertyIpV6IPsecUnauthDscp() (value uint32, err error) {
	retValue, err := instance.GetProperty("IpV6IPsecUnauthDscp")
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

// SetIpV6IPsecUnauthPerIPRateLimitBytesPerSec sets the value of IpV6IPsecUnauthPerIPRateLimitBytesPerSec for the instance
func (instance *MSFT_NetIPsecDoSPSetting) SetPropertyIpV6IPsecUnauthPerIPRateLimitBytesPerSec(value uint32) (err error) {
	return instance.SetProperty("IpV6IPsecUnauthPerIPRateLimitBytesPerSec", (value))
}

// GetIpV6IPsecUnauthPerIPRateLimitBytesPerSec gets the value of IpV6IPsecUnauthPerIPRateLimitBytesPerSec for the instance
func (instance *MSFT_NetIPsecDoSPSetting) GetPropertyIpV6IPsecUnauthPerIPRateLimitBytesPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("IpV6IPsecUnauthPerIPRateLimitBytesPerSec")
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

// SetIpV6IPsecUnauthRateLimitBytesPerSec sets the value of IpV6IPsecUnauthRateLimitBytesPerSec for the instance
func (instance *MSFT_NetIPsecDoSPSetting) SetPropertyIpV6IPsecUnauthRateLimitBytesPerSec(value uint32) (err error) {
	return instance.SetProperty("IpV6IPsecUnauthRateLimitBytesPerSec", (value))
}

// GetIpV6IPsecUnauthRateLimitBytesPerSec gets the value of IpV6IPsecUnauthRateLimitBytesPerSec for the instance
func (instance *MSFT_NetIPsecDoSPSetting) GetPropertyIpV6IPsecUnauthRateLimitBytesPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("IpV6IPsecUnauthRateLimitBytesPerSec")
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

// SetMaxPerIPRateLimitQueues sets the value of MaxPerIPRateLimitQueues for the instance
func (instance *MSFT_NetIPsecDoSPSetting) SetPropertyMaxPerIPRateLimitQueues(value uint32) (err error) {
	return instance.SetProperty("MaxPerIPRateLimitQueues", (value))
}

// GetMaxPerIPRateLimitQueues gets the value of MaxPerIPRateLimitQueues for the instance
func (instance *MSFT_NetIPsecDoSPSetting) GetPropertyMaxPerIPRateLimitQueues() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxPerIPRateLimitQueues")
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

// SetMaxStateEntries sets the value of MaxStateEntries for the instance
func (instance *MSFT_NetIPsecDoSPSetting) SetPropertyMaxStateEntries(value uint32) (err error) {
	return instance.SetProperty("MaxStateEntries", (value))
}

// GetMaxStateEntries gets the value of MaxStateEntries for the instance
func (instance *MSFT_NetIPsecDoSPSetting) GetPropertyMaxStateEntries() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxStateEntries")
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

// SetPerIPRateLimitQueueIdleTimeoutSeconds sets the value of PerIPRateLimitQueueIdleTimeoutSeconds for the instance
func (instance *MSFT_NetIPsecDoSPSetting) SetPropertyPerIPRateLimitQueueIdleTimeoutSeconds(value uint32) (err error) {
	return instance.SetProperty("PerIPRateLimitQueueIdleTimeoutSeconds", (value))
}

// GetPerIPRateLimitQueueIdleTimeoutSeconds gets the value of PerIPRateLimitQueueIdleTimeoutSeconds for the instance
func (instance *MSFT_NetIPsecDoSPSetting) GetPropertyPerIPRateLimitQueueIdleTimeoutSeconds() (value uint32, err error) {
	retValue, err := instance.GetProperty("PerIPRateLimitQueueIdleTimeoutSeconds")
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

// SetPrivateInterfaceAliases sets the value of PrivateInterfaceAliases for the instance
func (instance *MSFT_NetIPsecDoSPSetting) SetPropertyPrivateInterfaceAliases(value []string) (err error) {
	return instance.SetProperty("PrivateInterfaceAliases", (value))
}

// GetPrivateInterfaceAliases gets the value of PrivateInterfaceAliases for the instance
func (instance *MSFT_NetIPsecDoSPSetting) GetPropertyPrivateInterfaceAliases() (value []string, err error) {
	retValue, err := instance.GetProperty("PrivateInterfaceAliases")
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

// SetPrivateV6Address sets the value of PrivateV6Address for the instance
func (instance *MSFT_NetIPsecDoSPSetting) SetPropertyPrivateV6Address(value string) (err error) {
	return instance.SetProperty("PrivateV6Address", (value))
}

// GetPrivateV6Address gets the value of PrivateV6Address for the instance
func (instance *MSFT_NetIPsecDoSPSetting) GetPropertyPrivateV6Address() (value string, err error) {
	retValue, err := instance.GetProperty("PrivateV6Address")
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

// SetPublicInterfaceAliases sets the value of PublicInterfaceAliases for the instance
func (instance *MSFT_NetIPsecDoSPSetting) SetPropertyPublicInterfaceAliases(value []string) (err error) {
	return instance.SetProperty("PublicInterfaceAliases", (value))
}

// GetPublicInterfaceAliases gets the value of PublicInterfaceAliases for the instance
func (instance *MSFT_NetIPsecDoSPSetting) GetPropertyPublicInterfaceAliases() (value []string, err error) {
	retValue, err := instance.GetProperty("PublicInterfaceAliases")
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

// SetPublicV6Address sets the value of PublicV6Address for the instance
func (instance *MSFT_NetIPsecDoSPSetting) SetPropertyPublicV6Address(value string) (err error) {
	return instance.SetProperty("PublicV6Address", (value))
}

// GetPublicV6Address gets the value of PublicV6Address for the instance
func (instance *MSFT_NetIPsecDoSPSetting) GetPropertyPublicV6Address() (value string, err error) {
	retValue, err := instance.GetProperty("PublicV6Address")
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

// SetStateIdleTimeoutSeconds sets the value of StateIdleTimeoutSeconds for the instance
func (instance *MSFT_NetIPsecDoSPSetting) SetPropertyStateIdleTimeoutSeconds(value uint32) (err error) {
	return instance.SetProperty("StateIdleTimeoutSeconds", (value))
}

// GetStateIdleTimeoutSeconds gets the value of StateIdleTimeoutSeconds for the instance
func (instance *MSFT_NetIPsecDoSPSetting) GetPropertyStateIdleTimeoutSeconds() (value uint32, err error) {
	retValue, err := instance.GetProperty("StateIdleTimeoutSeconds")
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
