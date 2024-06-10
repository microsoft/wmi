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

// WebLimitsSection struct
type WebLimitsSection struct {
	*ConfigurationSection

	//
	ConnectionTimeout string

	//
	DemandStartThreshold uint32

	//
	DynamicIdleThreshold uint32

	//
	DynamicRegistrationThreshold uint32

	//
	HeaderWaitTimeout string

	//
	MaxGlobalBandwidth uint32

	//
	MinBytesPerSecond uint32
}

func NewWebLimitsSectionEx1(instance *cim.WmiInstance) (newInstance *WebLimitsSection, err error) {
	tmp, err := NewConfigurationSectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WebLimitsSection{
		ConfigurationSection: tmp,
	}
	return
}

func NewWebLimitsSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WebLimitsSection, err error) {
	tmp, err := NewConfigurationSectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WebLimitsSection{
		ConfigurationSection: tmp,
	}
	return
}

// SetConnectionTimeout sets the value of ConnectionTimeout for the instance
func (instance *WebLimitsSection) SetPropertyConnectionTimeout(value string) (err error) {
	return instance.SetProperty("ConnectionTimeout", (value))
}

// GetConnectionTimeout gets the value of ConnectionTimeout for the instance
func (instance *WebLimitsSection) GetPropertyConnectionTimeout() (value string, err error) {
	retValue, err := instance.GetProperty("ConnectionTimeout")
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

// SetDemandStartThreshold sets the value of DemandStartThreshold for the instance
func (instance *WebLimitsSection) SetPropertyDemandStartThreshold(value uint32) (err error) {
	return instance.SetProperty("DemandStartThreshold", (value))
}

// GetDemandStartThreshold gets the value of DemandStartThreshold for the instance
func (instance *WebLimitsSection) GetPropertyDemandStartThreshold() (value uint32, err error) {
	retValue, err := instance.GetProperty("DemandStartThreshold")
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

// SetDynamicIdleThreshold sets the value of DynamicIdleThreshold for the instance
func (instance *WebLimitsSection) SetPropertyDynamicIdleThreshold(value uint32) (err error) {
	return instance.SetProperty("DynamicIdleThreshold", (value))
}

// GetDynamicIdleThreshold gets the value of DynamicIdleThreshold for the instance
func (instance *WebLimitsSection) GetPropertyDynamicIdleThreshold() (value uint32, err error) {
	retValue, err := instance.GetProperty("DynamicIdleThreshold")
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

// SetDynamicRegistrationThreshold sets the value of DynamicRegistrationThreshold for the instance
func (instance *WebLimitsSection) SetPropertyDynamicRegistrationThreshold(value uint32) (err error) {
	return instance.SetProperty("DynamicRegistrationThreshold", (value))
}

// GetDynamicRegistrationThreshold gets the value of DynamicRegistrationThreshold for the instance
func (instance *WebLimitsSection) GetPropertyDynamicRegistrationThreshold() (value uint32, err error) {
	retValue, err := instance.GetProperty("DynamicRegistrationThreshold")
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

// SetHeaderWaitTimeout sets the value of HeaderWaitTimeout for the instance
func (instance *WebLimitsSection) SetPropertyHeaderWaitTimeout(value string) (err error) {
	return instance.SetProperty("HeaderWaitTimeout", (value))
}

// GetHeaderWaitTimeout gets the value of HeaderWaitTimeout for the instance
func (instance *WebLimitsSection) GetPropertyHeaderWaitTimeout() (value string, err error) {
	retValue, err := instance.GetProperty("HeaderWaitTimeout")
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

// SetMaxGlobalBandwidth sets the value of MaxGlobalBandwidth for the instance
func (instance *WebLimitsSection) SetPropertyMaxGlobalBandwidth(value uint32) (err error) {
	return instance.SetProperty("MaxGlobalBandwidth", (value))
}

// GetMaxGlobalBandwidth gets the value of MaxGlobalBandwidth for the instance
func (instance *WebLimitsSection) GetPropertyMaxGlobalBandwidth() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxGlobalBandwidth")
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

// SetMinBytesPerSecond sets the value of MinBytesPerSecond for the instance
func (instance *WebLimitsSection) SetPropertyMinBytesPerSecond(value uint32) (err error) {
	return instance.SetProperty("MinBytesPerSecond", (value))
}

// GetMinBytesPerSecond gets the value of MinBytesPerSecond for the instance
func (instance *WebLimitsSection) GetPropertyMinBytesPerSecond() (value uint32, err error) {
	retValue, err := instance.GetProperty("MinBytesPerSecond")
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
